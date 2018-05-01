package ipc

// ipc框架目的很简单，进行进程间的通信，封装通信编码包的细节，让使用者更专注于业务。
// 这里使用channel作为模块之间的通信方式，虽然channel可以传递任意数据类型，甚至可以包括一个channel，
// 但是我们为了我们的架构更容易分拆，在实现过程中严格限制了只能用于传递JSON格式的字符串类型数据，
// 这样如果之后想将这样的单进程实例改为多进程的分布式架构，也不需要全部重写，只需要更换通信层即可。

import "testing"

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Responce {
	resp := Responce{method, params}
	return &resp
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("1", "From Client1")
	resp2, _ := client2.Call("1", "From Client2")

	if resp1.Body != "From Client1" || resp2.Body != "From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2", resp2)
	}

	client1.Close()
	client2.Close()
}
