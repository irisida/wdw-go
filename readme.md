# Web development with Go 

This repo is following along with the `Apress` web development with Go book by Shiju Varghese.

## In the beginning. Servers Part One

- [01-01 - Static web server / fileserver](01-servers-part-one/01-01-staticweb) - Here we create a simple, trivial static fileserver demo that allows the user to select files (as links) and see them served. For the purposes of demonstration we have examples for an html page and simple text file. The principle extends to other file types as well as images. We're covering the basics of defining a server and running it with the http package via `http.ListenAndServe()`. 
- [01-02 - Custom handler](/01-servers-part-one/01-02-customHandler) - Here we have implemented a simple `mux` that calls a messageHandler func. This func implements the `ServeHTTP interface` and performs a simple write operation to the writer. We have defined three simple endpoints here that show the different routes and their different rendered messages in action. 
- [01-03 - Using functions as handlers](/01-servers-part-one/01-03-funcHandler) - Here we see a demo of the `HandlerFunc` as a way of reducing the potential verbosity of the approach taken in the `02-custom handler` section above. Utilising the HandlerFunc allows us to use/reuse the handler for multiple endpoints and keep the code fairly lean.  
- [01-04 - HandleFunc with a closure](/01-servers-part-one/01-04-handleFuncWithClosure) - Here we're creating a closure with the messageHandler so that it returns a `http.HandlerFunc` function with the message that was passed into it. When that endpoint is hit it will present the message. To call it we use the `mux.Handle("/routeName", messageHandler("message"))` format. 
- [01-05 - The default ServeMux](/01-servers-part-one/01-05-defaultServeMux) - here we're coving the use of the default ServeMux that comes built-in. Until now we've declared out mux as a value of `http.NewServeMux()` which worls fine but while we can declare and control our own mux we can easily use the default built-in servemux. Using the deault allows us to continue using the `HandlerFunc(route, function)` pattern we have seen in the last few examples.
- [01-06 - Defining the server struct](/01-servers-part-one/01-06-serverStruct) - here we are defining the parameters of the server using the `http.Server{}` struct. We are able t define the address, the read, write timeouts, buffer sizes etc.. It allows us to create a custom server and have different settings than we would have in the default server situation. We can still call the listen and serve, but here we do not need to pass additonal params because we have set them in the server struct.

## Gorilla Mux. Servers Part Two

- [02-01 - Gorilla mux](/01-servers-part-two/02-01-gorillaMux) - The samples above in part 1 have used the built in `http.ServeMux` and will work well for many if not most common scenarios. The `mux` package from gorilla (Gorilla web toolkit) is a fairly powerful request router. It is useful for RESTful services and implements the `http.Handler` interface so is compatible with the http package interfaces. With the mux package requests can be matched by the: `URL host`, `path`, `path prefix`, `schemes`, `header` and `query values`. It has custom matchers and routes as subrouters.   

## Templates and HTML with Go

- [03-01 - Basic templates](/03-templates/03-01-templates) - Here we see a single template being created that takes in two items, so the data is interpolated at compile time. 
- [03-02 - Templates with a slice of data](/03-templates/03-02-slice-of-data) - Here we define a template as a strong constant using the block definitions to interpolate our data `{{.AttributeName}}` to complete the render, where we're using a slice for our date we can put the template attributes in a `{{range .}} {{end}}` block to iterate over all the items (rows) in the slice. 