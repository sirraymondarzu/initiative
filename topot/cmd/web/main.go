package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// a struct is used to store multiple information in one return variable

type Welcome struct {
	Time    string
	Message string
}

// handler  function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi we are topot."))
	fmt.Println("hello world")
}
func displayTime(w http.ResponseWriter, r *http.Request) {
	//step 1 get the time
	localTime := time.Now().Format("3:04:05 PM")
	// step 2: read in the template file

	ts, _ := template.ParseFiles("./ui/html/display.time.tmpl")

	data := Welcome{
		Time:    localTime,
		Message: "Greetings Raymond! I hope you have a lovely day.",
	}

	// step 3: do the substitution (template engine)
	ts.Execute(w, data)
}

func main() {
	//create a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// url for time
	mux.HandleFunc("/time", displayTime)

	// create a file server

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// create a url mapping for a static directory

	mux.Handle("/resource/", http.StripPrefix("/resource", fileServer))

	//create a web server
	log.Println("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
