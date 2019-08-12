package znet

import (
	"fmt"
	"strconv"
	"zinx/ziface"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter //存放每个MsgId 所对应的处理方法的map属性
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (m *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok {
		// 没有注册路由
		fmt.Println("api msgId = ", request.GetMsgID(), " is not FOUND!")
		return
	}
	// 执行对应处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}
func (m *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	//1 判断当前msg绑定的API处理方法是否已经存在
	if _, ok := m.Apis[msgId]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}
	//2 添加msg与api的绑定关系
	m.Apis[msgId] = router
	fmt.Println("Add api msgId = ", msgId)
}
