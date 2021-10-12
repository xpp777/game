package netw

import (
	"game-ws/global"
	"game-ws/iface"
	"strconv"

	"go.uber.org/zap"
)

// MsgHandle -
type MsgHandle struct {
	Apis           map[uint32]iface.Router // 存放每个MsgID 所对应的处理方法的map属性
	WorkerPoolSize uint32                  // 业务工作Worker池的数量
	TaskQueue      []chan iface.Request    // Worker负责取任务的消息队列
}

// NewMsgHandle 创建MsgHandle
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]iface.Router),
		WorkerPoolSize: global.Config.WorkerPoolSize,
		// 一个worker对应一个queue
		TaskQueue: make([]chan iface.Request, global.Config.WorkerPoolSize),
	}
}

func (mh MsgHandle) DoMsgHandler(request iface.Request) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		zap.S().Error("api msgID = ", request.GetMsgID(), " is not FOUND!")
		return
	}
	// 执行对应处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh MsgHandle) AddRouter(msgID uint32, router iface.Router) {
	// 1 判断当前msg绑定的API处理方法是否已经存在
	if _, ok := mh.Apis[msgID]; ok {
		panic("repeated api , msgID = " + strconv.Itoa(int(msgID)))
	}
	// 2 添加msg与api的绑定关系
	mh.Apis[msgID] = router
	zap.S().Debug("Add api msgID = ", msgID)
}

func (mh MsgHandle) StartWorkerPool() {
	// 遍历需要启动worker的数量，依此启动
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 一个worker被启动
		// 给当前worker对应的任务队列开辟空间
		mh.TaskQueue[i] = make(chan iface.Request, global.Config.MaxWorkerTaskLen)
		// 启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}
func (mh MsgHandle) SendMsgToTaskQueue(request iface.Request) {
	// 根据ConnID来分配当前的连接应该由哪个worker负责处理
	// 轮询的平均分配法则
	// 得到需要处理此条连接的workerID
	workerID := uint32(request.GetConnection().GetConnID()) % mh.WorkerPoolSize
	zap.S().Debug("Add ConnID=", request.GetConnection().GetConnID(), " request msgID=", request.GetMsgID(), "  to workerID=", workerID)
	// 将请求消息发送给任务队列
	mh.TaskQueue[workerID] <- request
}

// StartOneWorker 启动一个Worker工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan iface.Request) {
	zap.S().Debug("Worker ID = ", workerID, " is started.")
	// 不断的等待队列中的消息
	for {
		select {
		// 有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}
