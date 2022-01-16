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

func Endpoint(path string) endpoint {
	return endpoint{path: path}
}

func DynamicJsonBodyParser(body io.ReadCloser) map[string]interface{} {
	var out map[string]interface{}
	data, _ := ioutil.ReadAll(body)
	json.Unmarshal(data, &out)
	return out
}

func Listen(port int) {
	fmt.Printf("Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
