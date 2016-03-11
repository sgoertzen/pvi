package main

import (
	"flag"
	"log"
	//"github.com/sgoertzen/veye"
	"net/http"
	"io"
)

var globalpath string

// Program to read in poms and determine
func main() {
	var path = flag.String("path", ".", "The `directory` that contains subfolders with maven projects.  Example: '/user/code/projects/'")
	flag.Parse()

	//projects := pvi.GetProjects(*path)

	//veye.SetKey("something")
	runServer(*path)

	log.Println("Done")
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!" + globalpath)
}

func runServer (path string) {
	globalpath = path
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}