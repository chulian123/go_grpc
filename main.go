package main

import (
	"fmt"
	"go_grpc/service"
	"google.golang.org/protobuf/proto"
)

func main() {
	user := &service.User{
		Username: "zhiqing",
		Age:      20,
	}
	//序列号过程
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(marshal)
	//反序列化
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser)

}
