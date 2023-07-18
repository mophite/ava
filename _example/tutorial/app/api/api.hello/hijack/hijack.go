package hijack

import (
	"github.com/go-ava/ava"
	"github.com/go-ava/ava/_example/tutorial/proto/phello"
	"net/http"
)

func HijackWriter(c *ava.Context, r *http.Request, w http.ResponseWriter, req, rsp *ava.Packet) bool {
	var sayRsp phello.HttpApiRsp
	ava.MustUnmarshal(rsp.Bytes(), &sayRsp)
	w.Write([]byte(sayRsp.Data))
	return true
}
