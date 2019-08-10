package znet

import (
	"fmt"
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
		_, err := conn.Write([]byte("hello ZINX"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ")
			return
		}

		fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}

//ping test 自定义路由
type PingRouter struct {
	BaseRouter //一定要先基础BaseRouter
}

func (r *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}
func (r *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}
func (r *PingRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Server 模块的测试函数
func TestServer(t *testing.T) {

	/*
		服务端测试
	*/
	//1 创建一个server 句柄 s
	s := NewServer("[zinx V0.1]")
	s.AddRoute(&PingRouter{})

	/*
		客户端测试
	*/
	go ClientTest()

	//2 开启服务
	s.Serve()

}
