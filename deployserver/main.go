package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch(){
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstpage(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "<h1>hello, this is my first deploy server!</h1>")
	reLaunch()
}

func main() {
	http.HandleFunc("/", firstpage)
	http.ListenAndServe(":8001", nil)
}