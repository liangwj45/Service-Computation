package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/liangwj45/Service-Computing/web/model"
)

type Handler struct{}

var tpl *template.Template
var user = new(model.User)

func init() {
	tpl = template.Must(template.ParseFiles("templates/info.html"))
}

func (uc *Handler) SignUp(w http.ResponseWriter, req *http.Request) {
	if err := json.NewDecoder(req.Body).Decode(user); err != nil {
		req.Body.Close()
		log.Fatalln(err)
	}
	if err := tpl.Execute(w, user); err != nil {
		log.Fatalln(err)
	}
}

func (uc *Handler) GetInfo(w http.ResponseWriter, req *http.Request) {
	if err := tpl.Execute(w, req.URL.Query()); err != nil {
		log.Fatalln(err)
	}
}
