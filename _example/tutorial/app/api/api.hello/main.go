package main

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/go-ava/ava"
	"github.com/go-ava/ava/_example/tutorial/app/api/api.hello/hello"
	"github.com/go-ava/ava/_example/tutorial/app/api/api.hello/hijack"
	"github.com/go-ava/ava/_example/tutorial/app/api/api.hello/http"
	"github.com/go-ava/ava/_example/tutorial/app/api/api.hello/im"
	"github.com/go-ava/ava/_example/tutorial/internal/ipc"
	"github.com/go-ava/ava/_example/tutorial/proto/phello"
	"github.com/go-ava/ava/_example/tutorial/proto/pim"
)

// ```shell
// curl -h "Content-Type:application/json" -X POST -d '{"ping": "ping"}' http://127.0.0.1:10000/hello/say/hi
// ```
func main() {
	s := ava.New(
		ava.HttpGetRootPathRedirect("/hello/say/hihttp"),
		ava.Namespace("api.hello"),
		ava.HttpApiAdd("0.0.0.0:10000"),
		ava.TCPApiPort(10001),
		ava.WssApiAddr("0.0.0.0:10002", "/hello"),
		ava.Hijacker(hijack.HijackWriter),
		ava.EtcdConfig(&clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}}),
	)

	phello.RegisterSaySrvServer(s.Server(), &hello.Say{})

	// for ava/_example/javascript service
	pim.RegisterImServer(s.Server(), im.NewIm())

	phello.RegisterHttpServer(s.Server(), &http.Http{})

	ipc.InitIpc(s)

	s.Run()
}
