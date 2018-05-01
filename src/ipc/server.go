package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Responce struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Responce
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)
	go func(c chan string) {
		for {
			request := <-c

			if request == "CLOSE" {
				break
			}
			var req Request
			// marshal 把对象转换成json，Unmarshal 用来把json数据转换回对象
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}
			resp := server.Handle(req.Method, req.Params)

			b, err := json.Marshal(resp)

			c <- string(b)
		}

		fmt.Println("session closed")

	}(session)

	fmt.Println("A new session has been created successfully!")
	return session
}
