package gover

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/*
Hosts a signal page application

The folder parameter is the folder which the SPA app is located.

The limit variable is how deep do you allow the URLs to go.

For example:

"/static/js/temp/" is a length of 3, because is goes 3 folders deep.

"/static/js/temp/old" is a length of 4, because is goes 4 folders deep.

"/static/js/temp/old/index.js" is still a length of 4, because the file at the end doesn't count.

A smaller limit is better for performance.
*/
func HostSPA(folder string, limit int) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fileReg, _ := regexp.Compile(`([a-z]+\.[a-z]+)$`)
		requestingFile := fileReg.MatchString(req.URL.Path)
		deepUrl := strings.Count(req.URL.Path, "/") > 2

		//if requesting a file on the /
		if req.URL.Path == "/" || (requestingFile && !deepUrl) {
			http.FileServer(http.Dir(folder)).ServeHTTP(w, req)

			//if requesting a file deeper then / which is not index.html
		} else if deepUrl && requestingFile {
			url := fileReg.ReplaceAllString(req.URL.Path, "")
			path, found := doesPathExist(folder, url, limit, 0)

			if found {
				r, _ := regexp.Compile(`([a-z]+\.[a-z]+)$`)
				filename := r.FindString(req.URL.Path)
				filepath := folder + "/" + path + filename
				data, err := ioutil.ReadFile(filepath)

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Fatal(err)
					return
				}

				w.Write(data)
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - page not found"))
			}

			//if requesting a file deeper then / which is index.html
		} else {
			data, err := ioutil.ReadFile(fmt.Sprintf("%v/index.html", folder))

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatal(err)
				return
			}

			w.Write(data)
		}
	})
}

func doesPathExist(folder string, path string, limit int, index int) (string, bool) {
	if path == "" || path == "/" || index >= limit {
		return "Error", false
	}

	fullpath := folder + path

	if doesFileExist(fullpath) {
		return doesPathExist(folder, goOneBack(path), limit, index+1)
	}

	return path, true
}

func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func goOneBack(path string) string {
	r, _ := regexp.Compile(`^(/[^/]+)`)

	return r.ReplaceAllString(path, "")
}
