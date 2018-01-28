package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

func genericHandler(w http.ResponseWriter, r *http.Request, filepath string) {
	log.Println("Request for " + filepath + " in genericHandler")
	page, err := ioutil.ReadFile(filepath)
	if err != nil {
		http.NotFound(w, r)
		log.Fatal(err)
	}
	w.Write(page)
}

func makeHandler(filepath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		genericHandler(w, r, filepath)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Redirecting reqest for " + r.URL.Path + " to index.html")
	http.Redirect(w, r, "/index.html", http.StatusFound)
}

func main() {
	os.Remove(LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	output := io.MultiWriter(os.Stderr, f)
	output.Write([]byte("-- Log for " + time.Now().String() + " --\n"))
	log.SetOutput(output)
	log.Println("Starting server")
	http.HandleFunc("/", handler)
	http.HandleFunc("/index.html", mainHandler)
	http.HandleFunc("/about.html", makeHandler("html/about.html"))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("js"))))
	http.ListenAndServe(":8080", nil)
}
