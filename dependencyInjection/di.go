package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Nan")
}

func main() {
	//Greet(os.Stdout, "Nan")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler)))
}
