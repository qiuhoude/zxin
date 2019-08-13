package ziface

import "net"

//定义连接接口
type IConnection interface {
	//启动连接，让当前连接开始工作
	Start()
	//停止连接，结束当前连接状态M
	Stop()
	//从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	//获取当前连接ID
	GetConnID() uint32
	//获取远程客户端地址信息
	RemoteAddr() net.Addr
	//直接将Message数据发送数据给远程的TCP客户端
	SendMsg(uint32, []byte) error
	//直接将Message数据发送给远程的TCP客户端(有缓冲)
	SendBuffMsg(uint32, []byte) error

	// 连接属性设置 ,类似于netty中 channel.attr
	SetProperty(key string, value interface{})
	//获取链接属性
	GetProperty(key string) (interface{}, error)
	//移除链接属性
	RemoveProperty(key string)
}

//定义一个统一处理链接业务的接口 int表示读到的大小
type HandFunc func(*net.TCPConn, []byte, int) error
