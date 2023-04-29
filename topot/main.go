package main
import
(
	"log"
	"net/http"
	"fmt"
)
//handler  function
func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hi we are topot."))
	//fmt.Println("hello world")j
}

func main(){
	//create a new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
    //create a web server
	log.Println("Starting server on port :4000")
	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)

}
