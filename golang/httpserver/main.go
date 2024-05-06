package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	for i := 1; i <= 10; i++ {
		endpoint := fmt.Sprintf("/endpoint%d", i)
		http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Endpoint %d", i)
		})
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
