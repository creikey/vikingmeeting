package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	LOGFILE = "log.txt"
)

func getStyles(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for styles.css")
	page, err := ioutil.ReadFile("css/styles.css")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(page)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "H("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources")))) ello there! I love %s!", r.URL.Path)
	log.Println("Request for index.html")
	page, err := ioutil.ReadFile("html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(page)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Redirecting reqest for " + r.URL.Path + " to index.html")
	http.Redirect(w, r, "/index.html", http.StatusFound)
}

func main() {
	log.Println("Starting server")
	http.HandleFunc("/", handler)
	http.HandleFunc("/index.html", mainHandler)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("js"))))
	http.ListenAndServe(":8080", nil)
}
