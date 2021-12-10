package view

import (
	"html/template"
	"log"
	"net/http"
)

func Render(data map[string]interface{}, w http.ResponseWriter, tpl string) {
	var tplPaths []string
	tplPaths = append(tplPaths, "tpl/"+tpl+".html")
	t, err := template.ParseFiles(tplPaths...)
	if err != nil {
		log.Println("posts template err:", err)
		return
	}
	t.Execute(w, data)
}
