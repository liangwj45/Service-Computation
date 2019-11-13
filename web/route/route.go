package route

import (
	"net/http"

	"github.com/liangwj45/Service-Computing/web/controller"
)

var handler = new(controller.Handler)

func init() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/SignUp", handler.SignUp)
	http.HandleFunc("/user", handler.GetInfo)
	http.HandleFunc("/unknown", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		_, _ = w.Write([]byte("Method Not Implemented!"))
	})
}
