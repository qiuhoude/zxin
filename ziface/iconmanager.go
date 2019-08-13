package ziface

// 连接管理抽象
type IConnManager interface {
	Add(conn IConnection)                   // 添加连接
	Remove(conn IConnection)                // 删除连接
	Get(connID uint32) (IConnection, error) // 获取连接
	Len() int                               //连接数量
	ClearConn()                             //清除所有连接
}
