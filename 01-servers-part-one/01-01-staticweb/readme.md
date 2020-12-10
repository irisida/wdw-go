# simple Fileserver

here we see a simple FileServer implementation that lostens on the port `8080` and will serve up the files in the `public` folder. 

```go
package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
}

```


