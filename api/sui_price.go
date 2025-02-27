package api

import (
	"net/http"

	"github.com/ipoluianov/cetuspools/system"
)

func init() {
}

func SuiPrice(w http.ResponseWriter, r *http.Request) {
	str := system.Get().PoolsAsTable()
	w.Write([]byte("SuiPrice:" + str))
}
