package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Title     string
	FirstName string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", abot)
	http.HandleFunc("/contact", cntct)
	http.HandleFunc("/apply", aply)
	http.Handle("/index", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	pd := pageData{
		Title: "Index Page",
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println("LOGG", err)
		http.Error(w, "Internal Serverrrrr Error", http.StatusInternalServerError)
		return

	}

}
func abot(w http.ResponseWriter, req *http.Request) {
	pd := pageData{
		Title: "About Page",
	}
	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func cntct(w http.ResponseWriter, req *http.Request) {
	pd := pageData{
		Title: "Contact Page",
	}
	err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func aply(w http.ResponseWriter, req *http.Request) {
	pd := pageData{
		Title: "Index Page",
	}
	var first string
	if req.Method == http.MethodPost {
		first = req.FormValue("fname")
		pd.FirstName = first

	}
	err := tpl.ExecuteTemplate(w, "apply.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
