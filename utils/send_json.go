package utils

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, v any, err error) {
	if err != nil {
		SendError(w, err)
		return
	}
	var bs []byte
	bs, err = json.MarshalIndent(v, "", " ")
	if err != nil {
		SendError(w, err)
		return
	}
	w.Write(bs)
}
