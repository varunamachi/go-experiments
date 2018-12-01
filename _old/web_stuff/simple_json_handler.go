package web_stuff
import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}


func Run_JsonServe() {
	http.HandleFunc("/simple_json", func( writer http.ResponseWriter, request *http.Request ) {
		person := Person{"Rob Pike", 50}
		output, err := json.Marshal(person)
		if err != nil {
			writer.Write([]byte("Oh Lord! Error"))
		} else {
			fmt.Fprintf(writer, string(output))
		}
	})
	http.ListenAndServe(":8080", nil)
}