package main
import
(
	"log"
	"net/http"
	//"fmt"
)


func main(){
	//create a file server
	fileServer := http.FileServer(http.Dir("."))




	
    //create a web server
	log.Println("Starting server on port :4000")
	err := http.ListenAndServe(":4000",fileServer)
	log.Fatal(err)

}

