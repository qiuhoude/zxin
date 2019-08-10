package ziface

// 将请求封装到message中 ,抽象结构
type IMessage interface {
	GetDataLen() uint32 // 获取消息数据段长度
	GetMsgId() uint32   // 获取消息ID
	GetData() []byte    // 消息内容

	SetMsgId(uint32)   // 设置消息ID
	SetData([]byte)    // 设置消息内容
	SetDataLen(uint32) //设置消息段长度
}
