# go-router

### Note

This is purely educational so you can make your own router.

### Usage


```GO
package main

import(
    "github.com/xxxvenom/go-router"
    "net/http"
    "log"
)

func boo(w http.ResponseWriter, r *http.Request){
    // do something
}

func foo(w http.ResponseWriter, r *http.Request){
    // do something
}


func main() {
	router := NewRouter()
	router.GET("/boo", boo)
	router.POST("/foo", foo)
	log.Fatal(http.ListenAndServe(":8080", router))
}
```