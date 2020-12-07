# wdw-go 

- [01 - Static web server / fileserver](/01-staticweb) - A simple static fileserver demo that allows the user to select lined files and see them served. Examples are an html page and simple text file. 
- [02 - Custom handler](/02-customHandler) - here we have implemented a simple `mux` that calls a messageHandler func. This fun implements the ServeHTTP interface and performs a simple write to the writer. We have defined three simple endpoints here that show different routes and their rendered messages. 
- [03 - Using functions as handlers](/03-funcHandler) - here we see a demo of the `HandlerFunc` as a way of reducing the potential verbosity of the approach taken in the `02-custom handler` section above. Utilising the HandlerFunc allows us to use/reuse the handler for multiple endpoingts and keep the code fairly lean.  
- [04 - HandleFunc with a closure](/04-handleFuncWithClosure) - here we're creating a closure with the messageHandler so that it returns a `http.HandlerFunc` function with the message that was passed into it. When that endpoint is hit it will present the message. To call it we use the `mux.Handle("/routeName", messageHandler("messaage"))` format.  
