package gover

import (
	"net/http"
)

/*
This is how a middleware should look.

If returns false then middleware didn not pass
If returns true then it did pass
*/
type MiddlewareFunc func(w http.ResponseWriter, r *http.Request, md *MiddlewareData) bool

// The data the middlware can pass
type MiddlewareData map[string]interface{}

type routeData struct {
	handler        RequestHandler
	middleware     MiddlewareFunc
	middlewareData MiddlewareData
}

// Sets the middleware for the route
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

// Creates the endpoint
func (c endpoint) Create() {
	http.HandleFunc(c.path, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if c.get.handler != nil {
				if c.get.middleware == nil || c.get.middleware(w, r, &c.get.middlewareData) {
					c.get.handler(w, r, c.get.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support GET requests."))
			}
		case "POST":
			if c.post.handler != nil {
				if c.post.middleware == nil || c.post.middleware(w, r, &c.post.middlewareData) {
					c.post.handler(w, r, c.post.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support POST requests."))
			}
		case "PUT":
			if c.put.handler != nil {
				if c.put.middleware == nil || c.put.middleware(w, r, &c.put.middlewareData) {
					c.put.handler(w, r, c.put.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support PUT requests."))
			}
		case "PATCH":
			if c.patch.handler != nil {
				if c.patch.middleware == nil || c.patch.middleware(w, r, &c.patch.middlewareData) {
					c.patch.handler(w, r, c.patch.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support PATCH requests."))
			}
		case "DELETE":
			if c.delete.handler != nil {
				if c.delete.middleware == nil || c.delete.middleware(w, r, &c.delete.middlewareData) {
					c.delete.handler(w, r, c.delete.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support DELETE requests."))
			}
		case "COPY":
			if c.copy.handler != nil {
				if c.copy.middleware == nil || c.copy.middleware(w, r, &c.copy.middlewareData) {
					c.copy.handler(w, r, c.copy.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support COPY requests."))
			}
		case "HEAD":
			if c.head.handler != nil {
				if c.head.middleware == nil || c.head.middleware(w, r, &c.head.middlewareData) {
					c.head.handler(w, r, c.head.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support HEAD requests."))
			}
		case "OPTIONS":
			if c.options.handler != nil {
				if c.options.middleware == nil || c.options.middleware(w, r, &c.options.middlewareData) {
					c.options.handler(w, r, c.options.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support OPTIONS requests."))
			}
		case "LINK":
			if c.link.handler != nil {
				if c.link.middleware == nil || c.link.middleware(w, r, &c.link.middlewareData) {
					c.link.handler(w, r, c.link.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support LINK requests."))
			}
		case "UNLINK":
			if c.unlink.handler != nil {
				if c.unlink.middleware == nil || c.unlink.middleware(w, r, &c.unlink.middlewareData) {
					c.unlink.handler(w, r, c.unlink.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support UNLINK requests."))
			}
		case "PURGE":
			if c.purge.handler != nil {
				if c.purge.middleware == nil || c.purge.middleware(w, r, &c.purge.middlewareData) {
					c.purge.handler(w, r, c.purge.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support PURGE requests."))
			}
		case "LOCK":
			if c.lock.handler != nil {
				if c.lock.middleware == nil || c.lock.middleware(w, r, &c.lock.middlewareData) {
					c.lock.handler(w, r, c.lock.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support LOCK requests."))
			}
		case "UNLOCK":
			if c.unlock.handler != nil {
				if c.unlock.middleware == nil || c.unlock.middleware(w, r, &c.unlock.middlewareData) {
					c.unlock.handler(w, r, c.unlock.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support UNLOCK requests."))
			}
		case "PROPFIND":
			if c.propfind.handler != nil {
				if c.propfind.middleware == nil || c.propfind.middleware(w, r, &c.propfind.middlewareData) {
					c.propfind.handler(w, r, c.propfind.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support PROPFIND requests."))
			}
		case "VIEW":
			if c.view.handler != nil {
				if c.view.middleware == nil || c.view.middleware(w, r, &c.view.middlewareData) {
					c.view.handler(w, r, c.view.middlewareData)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("This endpoint doesn't support VIEW requests."))
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unknown http method used."))
		}
	})
}

// Sets the GET method for this endpoint
func (c *endpoint) Get(f RequestHandler) *routeData {
	c.get.handler = f
	c.get.middlewareData = make(map[string]interface{})
	return &c.get
}

// Sets the POST method for this endpoint
func (c *endpoint) Post(f RequestHandler) *routeData {
	c.post.handler = f
	c.post.middlewareData = make(map[string]interface{})
	return &c.post
}

// Sets the PUT method for this endpoint
func (c *endpoint) Put(f RequestHandler) *routeData {
	c.put.handler = f
	c.put.middlewareData = make(map[string]interface{})
	return &c.put
}

// Sets the PATCH method for this endpoint
func (c *endpoint) Patch(f RequestHandler) *routeData {
	c.patch.handler = f
	c.patch.middlewareData = make(map[string]interface{})
	return &c.patch
}

// Sets the DELETE method for this endpoint
func (c *endpoint) Delete(f RequestHandler) *routeData {
	c.delete.handler = f
	c.delete.middlewareData = make(map[string]interface{})
	return &c.delete
}

// Sets the COPY method for this endpoint
func (c *endpoint) Copy(f RequestHandler) *routeData {
	c.copy.handler = f
	c.copy.middlewareData = make(map[string]interface{})
	return &c.copy
}

// Sets the HEAD method for this endpoint
func (c *endpoint) Head(f RequestHandler) *routeData {
	c.head.handler = f
	c.head.middlewareData = make(map[string]interface{})
	return &c.head
}

// Sets the OPTIONS method for this endpoint
func (c *endpoint) Options(f RequestHandler) *routeData {
	c.options.handler = f
	c.options.middlewareData = make(map[string]interface{})
	return &c.options
}

// Sets the LINK method for this endpoint
func (c *endpoint) Link(f RequestHandler) *routeData {
	c.link.handler = f
	c.link.middlewareData = make(map[string]interface{})
	return &c.link
}

// Sets the UNLINK method for this endpoint
func (c *endpoint) Unlink(f RequestHandler) *routeData {
	c.unlink.handler = f
	c.unlink.middlewareData = make(map[string]interface{})
	return &c.unlink
}

// Sets the PURGE method for this endpoint
func (c *endpoint) Purge(f RequestHandler) *routeData {
	c.purge.handler = f
	c.purge.middlewareData = make(map[string]interface{})
	return &c.purge
}

// Sets the LOCK method for this endpoint
func (c *endpoint) Lock(f RequestHandler) *routeData {
	c.lock.handler = f
	c.lock.middlewareData = make(map[string]interface{})
	return &c.lock
}

// Sets the UNLOCK method for this endpoint
func (c *endpoint) Unlock(f RequestHandler) *routeData {
	c.unlock.handler = f
	c.unlock.middlewareData = make(map[string]interface{})
	return &c.unlock
}

// Sets the PROPFIND method for this endpoint
func (c *endpoint) Propfind(f RequestHandler) *routeData {
	c.propfind.handler = f
	c.propfind.middlewareData = make(map[string]interface{})
	return &c.propfind
}

// Sets the VIEW method for this endpoint
func (c *endpoint) View(f RequestHandler) *routeData {
	c.view.handler = f
	c.view.middlewareData = make(map[string]interface{})
	return &c.view
}
