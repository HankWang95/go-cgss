package ipc

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
