package html

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type NewsAggPages struct {
	Title string
}

func Gethtml(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("hi")
	p := NewsAggPages{Title: "A news app"}
	fp := path.Join("templates", "basictemplating.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(err)
}
