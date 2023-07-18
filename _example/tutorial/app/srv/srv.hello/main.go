package main

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/go-ava/ava"
	"github.com/go-ava/ava/_example/tutorial/app/srv/srv.hello/hello"
	"github.com/go-ava/ava/_example/tutorial/proto/phello"
)

func main() {
	s := ava.New(
		ava.HttpApiAdd("0.0.0.0:10000"),
		ava.Namespace("srv.hello"),
		ava.TCPApiPort(20001),
		ava.EtcdConfig(&clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}}),
	)

	phello.RegisterSaySrvServer(s.Server(), &hello.Say{})

	s.Run()
}
