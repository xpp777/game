package iface

/*
	连接管理抽象层
*/
type Search func(Connection)
type ConnManager interface {
	Add(conn Connection)                  // 添加链接
	Remove(conn Connection)               // 删除连接
	Get(connID int64) (Connection, error) // 利用ConnID获取链接
	Len() int                             // 获取链接数量
	Search(Search)                        // 查找连接
	ClearConn()                           // 删除并停止所有链接
}
