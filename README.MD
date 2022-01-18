# Gover - a simple http server library for golang.

## Endpoints
### Creating an endpoint.
Creating an endpoint with gover is very simple. All you need to do is call the `Endpoint` method with the wanted path.
```go
ep := gover.Endpoint("/api")
```
### Handling requests on the endpoint
To handle requests all you need to do is call the method on the endpoint instance which is corresponding the http method you want to use.

Example:
```go
ep := gover.Endpoint("/api")

//handling a GET method on /api
ep.Get(func(w http.ResponseWriter, r *http.Request, md map[string]string) {
	w.Write([]byte("You used a GET method"))
})

//handling a POST method on /api
ep.Post(func(w http.ResponseWriter, r *http.Request, md map[string]string) {
	w.Write([]byte("You used a POST method"))
})
```

### Middleware
How add a middleware function to just chain the `Middleware` method to the previous method we called.

The middleware must return a Boolean. The Boolean represents wherever the middleware passed.

Example:
```go
ep := gover.Endpoint("/api")

//handling a get method on /api with a middlware
ep.
	Get(func(w http.ResponseWriter, r *http.Request, md map[string]string) {
		w.Write([]byte("You used a GET method"))
	}).
	Middleware(func(w http.ResponseWriter, r *http.Request, md *gover.MiddlewareData) bool {
		//this middlware passes
		w.Write([]byte("This requests was not stopped by the middleware \n"))

		return true
	})

//handling a get method on /api with a middlware
ep.
	Post(func(w http.ResponseWriter, r *http.Request, md map[string]string) {
		w.Write([]byte("You used a POST method"))
	}).
	Middleware(func(w http.ResponseWriter, r *http.Request, md *gover.MiddlewareData) bool {
		//this middleware fails
		w.Write([]byte("This requests was stopped by the middleware"))

		return false
	})
```
### Passing data from the middleware to the handler
On the middlware you get a pointer to a map. 
On that map you can store data that will be accessible from the handler.
 
Example:
```go
ep := gover.Endpoint("/api")

//handling a get method on /api
ep.
	Get(func(w http.ResponseWriter, r *http.Request, md map[string]string) {
		n := md["name"]
		ln := md["last-name"]

		fmt.Println(n, ln)
	}).
	Middleware(func(w http.ResponseWriter, r *http.Request, md *gover.MiddlewareData) bool {
		//passing data
		(*md)["name"] = "ilay"
		(*md)["last-name"] = "bokobza"

		return true
	})
```
### Activating the endpoint
Doing what we did so far will not actually work. All it does is store all that information without doing anything. Activating the endpoint is really easy. All we need to do is call the `Create` method on the endpoint instance after declaring the handlers. NOTE: Your endpoint won't work if you call the method before declaring the handlers.

Example:
```go
ep := gover.Endpoint("/api")

//handling a get method on /api
ep.Get(func(w http.ResponseWriter, r *http.Request, md map[string]string) {
	w.Write([]byte("You used a GET method"))
})

ep.Create()
``` 
## Hosting
### Normal storage bin hosting
To host a static storage bin website, all you need to do is call the `HostFolder` method with the path to the folder you want to host.

Example
```go
gover.HostFolder("public")
```

This method is just a wapper for the normal way of doing it in go.

### Signal Page Application Hosting.
To Host a SPA (Signal Page Application) app, you need to call the `HostSPA` method with two arguments. 

The first one is the `folder` argument. This argument is the path to the folder you want to host.

The Second argument is the `limit` argument. This argument represent how deep do you allow your website's URLs to go.

For example:

`www.example.com/static/js/temp/` is a length of 3, because is goes 3 folders deep.

`www.example.com/static/js/temp/old` is a length of 4, because is goes 4 folders deep.

`www.example.com/static/js/temp/old/index.js` is still a length of 4, because the file at the end doesn't count.

A smaller limit is better for performance.

Code example
```go
gover.HostSPA("public",3)
```
This code will host a spa application and will only allow up 3 folders long URLs.
## Utilities
Gover have some nice utilities to make your life easier. They are not necessary, but they are nice to have.

### Dynamic JSON based body Parser
This method will allow you to easily parse the json out of request body. This method returns a `map[string]interface{}` and an `error`.

Example:
```go
ep := gover.Endpoint("/api")

//handling a get method on /api
ep.Get(func(w http.ResponseWriter, r *http.Request, md map[string]string) {
	bodyData, err := gover.DynamicJSONBodyParser(r.Body)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(bodyData)
})

ep.Create()
```
### Starting the server with gover
The `Listen` method will start the server and log the port which is running on. This method is just a wrapper for `http.ListenAndServe()`.

Example:
```go
port := 3000
gover.Listen(port)
```

### Getting files from request
Gover has an useful method that allows you to get the file from the request very easily.The `GetFile` method revives two arguments. The first one is the `fieldname`, and the second one is the request.
The `fieldname` is name property of the file input in your HTML form.

For Example: This is your HTML form
```html
<form>
  <input type="file" name="myFile" />
  <button>Submit</button>
</form>
```
you will need to pass to the `"myFile"` to the `fieldname` parameter.

Code Example:
```go
ep := gover.Endpoint("/upload")

ep.Post(func(w http.ResponseWriter, r *http.Request, md map[string]string) {
	data, handler, err := gover.GetFile("myFile", r)

	if err != nil {
		fmt.Println(err)
		return
	}

	ioutil.WriteFile(handler.Filename, data, 0666)
})

ep.Create()
```