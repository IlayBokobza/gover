package gover

import (
	"fmt"
	"net/http"
)

type crud struct {
	path     string
	get      RequestHandler
	post     RequestHandler
	put      RequestHandler
	patch    RequestHandler
	delete   RequestHandler
	copy     RequestHandler
	head     RequestHandler
	options  RequestHandler
	link     RequestHandler
	unlink   RequestHandler
	purge    RequestHandler
	lock     RequestHandler
	unlock   RequestHandler
	propfind RequestHandler
	view     RequestHandler
}

//Actives the endpoint
func (c crud) Create() {
	http.HandleFunc(c.path, func(w http.ResponseWriter, req *http.Request) {
		if c.get != nil && req.Method == "GET" {
			c.get(w, req)
		} else if c.post != nil && req.Method == "POST" {
			c.post(w, req)
		} else if c.put != nil && req.Method == "PUT" {
			c.put(w, req)
		} else if c.patch != nil && req.Method == "PATCH" {
			c.patch(w, req)
		} else if c.delete != nil && req.Method == "DELETE" {
			c.delete(w, req)
		} else if c.copy != nil && req.Method == "COPY" {
			c.copy(w, req)
		} else if c.head != nil && req.Method == "HEAD" {
			c.head(w, req)
		} else if c.options != nil && req.Method == "OPTIONS" {
			c.options(w, req)
		} else if c.link != nil && req.Method == "LINK" {
			c.link(w, req)
		} else if c.unlink != nil && req.Method == "UNLINK" {
			c.unlink(w, req)
		} else if c.purge != nil && req.Method == "PURGE" {
			c.purge(w, req)
		} else if c.lock != nil && req.Method == "LOCK" {
			c.lock(w, req)
		} else if c.unlock != nil && req.Method == "UNLOCK" {
			c.unlock(w, req)
		} else if c.propfind != nil && req.Method == "PROFIND" {
			c.propfind(w, req)
		} else if c.view != nil && req.Method == "VIEW" {
			c.view(w, req)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("404 - No handler for %v found on this endpoint", req.Method)))
		}
	})
}

func (c *crud) Get(f RequestHandler) {
	c.get = f
}

func (c *crud) Post(f RequestHandler) {
	c.post = f
}

func (c *crud) Put(f RequestHandler) {
	c.put = f
}

func (c *crud) Patch(f RequestHandler) {
	c.patch = f
}

func (c *crud) Delete(f RequestHandler) {
	c.delete = f
}

func (c *crud) Copy(f RequestHandler) {
	c.copy = f
}

func (c *crud) Head(f RequestHandler) {
	c.head = f
}

func (c *crud) Options(f RequestHandler) {
	c.options = f
}

func (c *crud) Link(f RequestHandler) {
	c.link = f
}

func (c *crud) Unlink(f RequestHandler) {
	c.unlink = f
}

func (c *crud) Purge(f RequestHandler) {
	c.purge = f
}

func (c *crud) Lock(f RequestHandler) {
	c.lock = f
}

func (c *crud) Unlock(f RequestHandler) {
	c.unlock = f
}

func (c *crud) Propfind(f RequestHandler) {
	c.propfind = f
}

func (c *crud) View(f RequestHandler) {
	c.view = f
}
