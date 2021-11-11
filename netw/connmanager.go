package netw

import (
	"errors"
	"sync"
	"time"

	"github.com/xiaomingping/game/iface"
)

// ConnManager 连接管理模块
type ConnManager struct {
	connections map[int64]iface.Connection
	connLock    sync.RWMutex
}

// NewConnManager 创建一个链接管理
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[int64]iface.Connection),
	}
}

func (connMgr *ConnManager) Add(conn iface.Connection) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	connMgr.connections[conn.GetConnID()] = conn
}

func (connMgr *ConnManager) Remove(conn iface.Connection) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	delete(connMgr.connections, conn.GetConnID())
}

func (connMgr *ConnManager) Get(connID int64) (iface.Connection, error) {
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()
	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	}
	return nil, errors.New("connection not found")
}

func (connMgr *ConnManager) Len() int {
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()
	length := len(connMgr.connections)
	return length
}

func (connMgr *ConnManager) ClearConn() {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	// 停止并删除全部的连接信息
	for connID, conn := range connMgr.connections {
		// 停止
		conn.Stop()
		// 删除
		delete(connMgr.connections, connID)
	}
}

func (connMgr *ConnManager) Search(s iface.Search) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	// 停止并删除全部的连接信息
	for _, conn := range connMgr.connections {
		s(conn)
	}
}

// ClearOneConn  利用ConnID获取一个链接 并且删除
func (connMgr *ConnManager) ClearOneConn(connID int64) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	connections := connMgr.connections
	if conn, ok := connections[connID]; ok {
		// 停止
		conn.Stop()
		// 删除
		delete(connections, connID)
		return
	}
	return
}

// 心跳检测
func (connMgr *ConnManager) PingAuth() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			// 停止并删除全部的连接信息
			for _, conn := range connMgr.connections {
				if conn.IsHeartbeatTimeout() {
					conn.Stop()
				}
			}
		}
	}

}
