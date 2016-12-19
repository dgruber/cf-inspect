package main

import (
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

func renderIndex(w http.ResponseWriter) error {
	tplt, err := template.ParseFiles("./templates/index.template")
	if err != nil {
		return err
	}
	appEnv, err := cfenv.Current()
	if err != nil {
		return tplt.Execute(w, cfenv.App{})
	}
	return tplt.Execute(w, *appEnv)
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	err := renderIndex(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = ":8888"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)

	http.ListenAndServe(port, r)
}
