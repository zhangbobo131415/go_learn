package main

import (
	"context"
	"first-go/sayhello"
	"fmt"
	"os"
	"strconv"

	"github.com/apache/thrift/lib/go/thrift"
)

type rpcService struct {
}

func (p *rpcService) Sayhello(ctx context.Context, a string) (string, error) {
	fmt.Println("hello:" + a)
	return "hello:" + a, nil
}

func (p *rpcService) Sum(ctx context.Context, a int32, b int32) (r string, err error) {
	return "your ans is :" + strconv.Itoa(int(a+b)), nil
}

func main() {
	netAddres := "172.21.13.121:8000"
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolfactory := thrift.NewTBinaryProtocolFactoryDefault()
	severtranport, err := thrift.NewTServerSocket(netAddres)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	handler := &rpcService{}
	processer := sayhello.NewHelloAndSumProcessor(handler)
	server := thrift.NewTSimpleServer4(processer, severtranport, transportFactory, protocolfactory)

	fmt.Println("thrift server in :", netAddres)
	server.Serve()
}
