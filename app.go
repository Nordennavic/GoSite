package main
import (
	"log"
	"net/http"
)

func main(){
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Listening...")
	err := http.ListenAndServe(":3004", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Current port: 3004")
	}

}
