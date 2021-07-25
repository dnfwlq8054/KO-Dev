package main

import (
	"github.com/dnfwlq8054/KO-Dev/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
