package main

import (
	"context"
	"first-go/sayhello"
	"fmt"
	"os"
	"strconv"

	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	netAddres := "172.21.15.216:8000"
	defaltContext := context.Background()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolfactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(netAddres)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	useTransport, _ := transportFactory.GetTransport(transport)
	client := sayhello.NewHelloAndSumClientFactory(useTransport, protocolfactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 172.21.13.121:8000", " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	var i int32
	for i = 0; i < 200; i++ {
		tmp, _ := client.Sum(defaltContext, i, i)
		s, _ := client.Sayhello(defaltContext, strconv.Itoa(int(i)))
		fmt.Println(tmp, s)
	}
}
