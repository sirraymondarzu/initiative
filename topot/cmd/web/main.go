package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv" // converts string to other
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

func getValues(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ts, _ := template.ParseFiles("./ui/html/input.area.page.tmpl")
		ts.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/area-calculator-2", http.StatusTemporaryRedirect)
		//if post metod is selected
	}

}
func calculateArea(w http.ResponseWriter, r *http.Request) {
	// sending multiple peices of data using a struct
	type UserData struct {
		Length float64
		Width  float64
		Area   float64
	}
	// get the length and the width for the form
	r.ParseForm()

	// save the values
	length := r.PostForm.Get("length")
	width := r.PostForm.Get("width")
	// calculate the area
	lengthORectangle, _ := strconv.ParseFloat(length, 64)
	widthORectangle, _ := strconv.ParseFloat(width, 64)
	areaORectangle := lengthORectangle * widthORectangle

	// create instande of the UserData
	data := UserData{
		Length: lengthORectangle,
		Width:  widthORectangle,
		Area:   areaORectangle,
	}

	// call the template engine
	ts, _ := template.ParseFiles("./ui/html/display.area.page.tmpl")
	ts.Execute(w, data)

}

func main() {
	//create a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// url for time
	mux.HandleFunc("/time", displayTime)
	mux.HandleFunc("/area-calculator", getValues)
	mux.HandleFunc("/area-calculator-2", calculateArea)

	// create a file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// create a url mapping for a static directory
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//create a web server
	log.Println("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
