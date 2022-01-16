package gover

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestHandler func(w http.ResponseWriter, req *http.Request, md map[string]string)

// Creates an endpoint instance
func Endpoint(path string) endpoint {
	return endpoint{path: path}
}

/*
Dynamicly parses the request body, if it is in json.

Tip: To convert an interface{} to a diffrent type, instead of using type(val) use val.(type)

For example:

var x interface{} = 5

y := x.(int)
*/
func DynamicJsonBodyParser(body io.ReadCloser) map[string]interface{} {
	var out map[string]interface{}
	data, _ := ioutil.ReadAll(body)
	json.Unmarshal(data, &out)
	return out
}

//Starts the server on a port
func Listen(port int) {
	fmt.Printf("Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

//Hosts a normal file bin
func HostFolder(path string) {
	http.Handle("/", http.FileServer(http.Dir(path)))
}
