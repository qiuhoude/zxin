package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/ziface"
)

//连接管理模块
type ConnManager struct {
	connections map[uint32]ziface.IConnection //管理的连接信息
	connLock    sync.RWMutex                  //读写连接的读写锁
}

//创建一个链接管理
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

//添加连接
func (c *ConnManager) Add(conn ziface.IConnection) {
	//保护共享资源Map 加写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()
	//将conn连接添加到ConnMananger中
	c.connections[conn.GetConnID()] = conn
	fmt.Println("connection add to ConnManager successfully: conn num = ", c.Len())

}

//删除连接信息
func (c *ConnManager) Remove(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if _, ok := c.connections[conn.GetConnID()]; ok { // 存在才删除
		delete(c.connections, conn.GetConnID())
		fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", c.Len())
	}
}

func (c *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	//保护共享资源Map 加读锁
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	//保护共享资源Map 加写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()
	//停止并删除全部的连接信息
	for connID, conn := range c.connections {
		delete(c.connections, connID)
		conn.Stop()
	}
	fmt.Println("Clear All Connections successfully: conn num = ", c.Len())

}
