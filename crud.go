package gover

import (
	"fmt"
	"net/http"
)

// This is how a middleware should look.
// If returns false then middleware didn not pass
// If returns true then it did pass
type MiddlewareFunc func(w http.ResponseWriter, req *http.Request) bool

type routeData struct {
	handler    RequestHandler
	middleware MiddlewareFunc
}

func (r *routeData) Middleware(m MiddlewareFunc) {
	r.middleware = m
}

type endpoint struct {
	path     string
	get      routeData
	post     routeData
	put      routeData
	patch    routeData
	delete   routeData
	copy     routeData
	head     routeData
	options  routeData
	link     routeData
	unlink   routeData
	purge    routeData
	lock     routeData
	unlock   routeData
	propfind routeData
	view     routeData
}

//Actives the endpoint
func (c endpoint) Create() {
	http.HandleFunc(c.path, func(w http.ResponseWriter, req *http.Request) {
		if c.get.handler != nil && req.Method == "GET" {
			if c.get.middleware == nil || c.get.middleware(w, req) {
				c.get.handler(w, req)
			}
		} else if c.post.handler != nil && req.Method == "POST" {
			if c.post.middleware == nil || c.post.middleware(w, req) {
				c.post.handler(w, req)
			}
		} else if c.put.handler != nil && req.Method == "PUT" {
			if c.put.middleware == nil || c.put.middleware(w, req) {
				c.put.handler(w, req)
			}
		} else if c.patch.handler != nil && req.Method == "PATCH" {
			if c.patch.middleware == nil || c.patch.middleware(w, req) {
				c.patch.handler(w, req)
			}
		} else if c.delete.handler != nil && req.Method == "DELETE" {
			if c.delete.middleware == nil || c.delete.middleware(w, req) {
				c.delete.handler(w, req)
			}
		} else if c.copy.handler != nil && req.Method == "COPY" {
			if c.copy.middleware == nil || c.copy.middleware(w, req) {
				c.copy.handler(w, req)
			}
		} else if c.head.handler != nil && req.Method == "HEAD" {
			if c.head.middleware == nil || c.head.middleware(w, req) {
				c.head.handler(w, req)
			}
		} else if c.options.handler != nil && req.Method == "OPTIONS" {
			if c.options.middleware == nil || c.options.middleware(w, req) {
				c.options.handler(w, req)
			}
		} else if c.link.handler != nil && req.Method == "LINK" {
			if c.link.middleware == nil || c.link.middleware(w, req) {
				c.link.handler(w, req)
			}
		} else if c.unlink.handler != nil && req.Method == "UNLINK" {
			if c.unlink.middleware == nil || c.unlink.middleware(w, req) {
				c.unlink.handler(w, req)
			}
		} else if c.purge.handler != nil && req.Method == "PURGE" {
			if c.purge.middleware == nil || c.purge.middleware(w, req) {
				c.purge.handler(w, req)
			}
		} else if c.lock.handler != nil && req.Method == "LOCK" {
			if c.lock.middleware == nil || c.lock.middleware(w, req) {
				c.lock.handler(w, req)
			}
		} else if c.unlock.handler != nil && req.Method == "UNLOCK" {
			if c.unlock.middleware == nil || c.unlock.middleware(w, req) {
				c.unlock.handler(w, req)
			}
		} else if c.propfind.handler != nil && req.Method == "PROFIND" {
			if c.propfind.middleware == nil || c.propfind.middleware(w, req) {
				c.propfind.handler(w, req)
			}
		} else if c.view.handler != nil && req.Method == "VIEW" {
			if c.view.middleware == nil || c.view.middleware(w, req) {
				c.view.handler(w, req)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("404 - No handler for %v found on this endpoint", req.Method)))
		}
	})
}

func (c *endpoint) Get(f RequestHandler) *routeData {
	c.get.handler = f
	return &c.get
}

func (c *endpoint) Post(f RequestHandler) *routeData {
	c.post.handler = f
	return &c.post
}

func (c *endpoint) Put(f RequestHandler) *routeData {
	c.put.handler = f
	return &c.put
}

func (c *endpoint) Patch(f RequestHandler) *routeData {
	c.patch.handler = f
	return &c.patch
}

func (c *endpoint) Delete(f RequestHandler) *routeData {
	c.delete.handler = f
	return &c.delete
}

func (c *endpoint) Copy(f RequestHandler) *routeData {
	c.copy.handler = f
	return &c.copy
}

func (c *endpoint) Head(f RequestHandler) *routeData {
	c.head.handler = f
	return &c.head
}

func (c *endpoint) Options(f RequestHandler) *routeData {
	c.options.handler = f
	return &c.options
}

func (c *endpoint) Link(f RequestHandler) *routeData {
	c.link.handler = f
	return &c.link
}

func (c *endpoint) Unlink(f RequestHandler) *routeData {
	c.unlink.handler = f
	return &c.unlink
}

func (c *endpoint) Purge(f RequestHandler) *routeData {
	c.purge.handler = f
	return &c.purge
}

func (c *endpoint) Lock(f RequestHandler) *routeData {
	c.lock.handler = f
	return &c.lock
}

func (c *endpoint) Unlock(f RequestHandler) *routeData {
	c.unlock.handler = f
	return &c.unlock
}

func (c *endpoint) Propfind(f RequestHandler) *routeData {
	c.propfind.handler = f
	return &c.propfind
}

func (c *endpoint) View(f RequestHandler) *routeData {
	c.view.handler = f
	return &c.view
}
