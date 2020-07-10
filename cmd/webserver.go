package main

import (
	"fmt"
	"net/http"
	"os"
	"webserver/pkg/ctrl"
)

func main() {
	httpServer := &http.Server{Addr: ":6100", Handler: ctrl.HttpRouter()}

	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Printf("http ListenAndServer err:%s\n", err)
		os.Exit(1)
	}
	select{}
}
