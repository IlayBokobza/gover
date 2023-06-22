package gover

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

type RequestHandler func(w http.ResponseWriter, r *http.Request, md map[string]string)

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
func DynamicJSONBodyParser(body io.ReadCloser) (map[string]interface{}, error) {
	var out map[string]interface{}
	data, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

// Starts the server on a port
func Listen(port int) {
	fmt.Printf("Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%v", port), nil))
}

// Hosts a normal file bin
func HostFolder(path string) {
	http.Handle("/", http.FileServer(http.Dir(path)))
}

/*
Gets the file from the request.

The fieldname argument is the name of the input field in your HTML.

For Example: Your HTML input is this:

<input type="file" name="myFile">

In this case you will need to pass to the function "myFile".
*/
func GetFile(fieldname string, r *http.Request) ([]byte, *multipart.FileHeader, error) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile(fieldname)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, nil, err
	}

	return data, handler, nil
}
