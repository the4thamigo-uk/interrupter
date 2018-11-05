package main

import (
	"context"
	"fmt"
	"github.com/the4thamigo-uk/interrupter"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/", indexHandler)

	svr := &http.Server{Addr: ":8080"}

	irpt := interrupter.New(func() {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		svr.Shutdown(ctx)
	})
	defer irpt.Close()

	err := svr.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
