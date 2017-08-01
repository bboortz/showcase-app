package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

func main() {

	// startup logs
	log.Printf("Starting showcase-app on the port %s", os.Getenv("PORT"))
	for _, pair := range os.Environ() {
		log.Printf("Env variable: %s\n", pair)
	}

	name, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Processing request")
		fmt.Fprintf(w, "Timestamp: %q\n", time.Now().Format(time.RFC850))
		fmt.Fprintf(w, "Hostname: %q\n\n", html.EscapeString(name))
		keys := make([]string, len(r.Header))
		i := 0
		for k, _ := range r.Header {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, k := range keys {
			for _, h := range r.Header[k] {
				fmt.Fprintf(w, "%v: %v\n", k, h)
			}

		}

	})

	log.Fatal(http.ListenAndServeTLS(os.Getenv("PORT"), os.Getenv("TLS_CERT"), os.Getenv("TLS_KEY"), nil))

}
