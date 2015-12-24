package web_stuff

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
	"encoding/json"
	"fmt"
)

func Run_GorillaServe() {
	router := mux.NewRouter()
	router.HandleFunc( "/api/{user:[a-zA-Z-0-9]+}", func( writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		name := params["user"]
		person := Person{name,30}
		output, err := json.Marshal(person)
		var message string = "";
		if err != nil {
			message = "Encode Error!" + err.Error()
		} else {
			message = string(output)
		}
		fmt.Fprintf(writer, message)
	})
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}