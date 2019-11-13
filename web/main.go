package main

import (
	"net/http"

	_ "github.com/liangwj45/Service-Computing/web/route"
)

func main() {
	_ = http.ListenAndServe(":2333", nil)
}
