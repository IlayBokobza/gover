package gover

import (
	"fmt"
	"net/http"
)

type RequestHandler func(w http.ResponseWriter, req *http.Request)

func Endpoint(path string) crud {
	return crud{path: path}
}

func Listen(port int) {
	fmt.Printf("Server listening on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
