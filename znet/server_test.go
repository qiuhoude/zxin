package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
	"zinx/ziface"
)

/*
   模拟客户端
*/
func ClientTest() {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		dp := NewDataPack()
		msg, _ := dp.Pack(NewMsgPackage(0, []byte("Zinx V0.5 Client Test Message")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}
		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
			fmt.Println("read head error ")
			return
		}
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}
		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*Message)
			msg.Data = make([]byte, msg.GetDataLen())
			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))

		}
		time.Sleep(1 * time.Second)
	}
}

//ping test 自定义路由
type PingRouter struct {
	BaseRouter //一定要先基础BaseRouter
}

func (r *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("Call Router PreHandle")

}
func (r *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", req.GetMsgID(), ", data=", string(req.GetData()))

	req.GetConnection().SendMsg(1, []byte("ping...ping...ping"))

}
func (r *PingRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("Call Router PostHandle")

}

//Server 模块的测试函数
func TestServer(t *testing.T) {

	/*
		服务端测试
	*/
	//1 创建一个server 句柄 s
	s := NewServer("[zinx V0.5]")
	s.AddRoute(&PingRouter{})

	/*
		客户端测试
	*/
	go ClientTest()

	//2 开启服务
	s.Serve()

}
