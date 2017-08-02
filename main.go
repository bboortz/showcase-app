package main

import (
	"os"
//	"fmt"
	"log"
//	"sort"
//	"html"
//	"time"
        "net/http"
        "crypto/tls"
        "html/template"
        "path/filepath"
)

type Header struct {
    Key string
    Value string
}

type Welcome struct {
	Title   string
	Message string
}

var templates = make(map[string]*template.Template)

func main() {

	// startup logs
	log.Printf("Starting showcase-app on the port %s", os.Getenv("PORT"))
	for _, pair := range os.Environ() {
		log.Printf("Env variable: %s\n", pair)
	}
        
        initializeTemplates()

	//name, _ := os.Hostname()

	// static files handling
  	fs := http.FileServer(http.Dir("static"))
  	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Processing request: %v %v %v", r.Method, r.URL, r.Proto)
                
                message := Welcome{Title: "Bootstrap, Go, and GAE", Message: "Bootstrap added to Golang on App Engine.  Feel free to customize further"}
                templates["welcome.html"].ExecuteTemplate(w, "outerTheme", &message)
/*
                t := template.New("tmpl/index.html")
                log.Printf("111")
                t, _ = t.ParseFiles("tmpl/index.html")
                log.Printf("222")
                p := Header{Key: "Testkey", Value: "Testvalue"}
                t.Execute(w, p)
                log.Printf("333")
*/
                /*

		fmt.Fprintf(w, "Timestamp: %q\n", time.Now().Format(time.RFC850))
		fmt.Fprintf(w, "Hostname: %q\n\n", html.EscapeString(name))

                fmt.Fprintf(w, "Request headers:\n")
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
                */

	})

	// https://golang.org/pkg/crypto/tls/
	tlsCipherMap := map[string]uint16{
		"TLS_RSA_WITH_RC4_128_SHA":                tls.TLS_RSA_WITH_RC4_128_SHA,
		"TLS_RSA_WITH_3DES_EDE_CBC_SHA":           tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		"TLS_RSA_WITH_AES_128_CBC_SHA":            tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		"TLS_RSA_WITH_AES_256_CBC_SHA":            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		"TLS_RSA_WITH_AES_128_CBC_SHA256":         tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
		"TLS_RSA_WITH_AES_128_GCM_SHA256":         tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		"TLS_RSA_WITH_AES_256_GCM_SHA384":         tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		"TLS_ECDHE_ECDSA_WITH_RC4_128_SHA":        tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
		"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA":    tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA":    tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		"TLS_ECDHE_RSA_WITH_RC4_128_SHA":          tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
		"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA":     tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":      tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":      tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256": tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256":   tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256":   tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256": tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384":   tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384": tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305":    tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305":  tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		"TLS_FALLBACK_SCSV":                       tls.TLS_FALLBACK_SCSV,
	}
	tlsMinVersionMap := map[string]uint16{
		"VersionSSL30": tls.VersionSSL30,
		"VersionTLS10": tls.VersionTLS10,
		"VersionTLS11": tls.VersionTLS11,
		"VersionTLS12": tls.VersionTLS12,
	}

	tlsConfig := &tls.Config{
		CipherSuites: []uint16{tlsCipherMap[os.Getenv("TLS_CIPHER")]},
		PreferServerCipherSuites: true,
		MinVersion:               tlsMinVersionMap[os.Getenv("TLS_MIN_VERSION")],
	}
	tlsConfig.BuildNameToCertificate()

	httpServer := &http.Server{
		Addr:      os.Getenv("PORT"),
		TLSConfig: tlsConfig,
	}

	log.Fatal(httpServer.ListenAndServeTLS(os.Getenv("TLS_CERT"), os.Getenv("TLS_KEY")))

}

// Base template is 'theme.html'  Can add any variety of content fillers in /layouts directory
func initializeTemplates() {
	layouts, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, layout := range layouts {
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(layout, "templates/layouts/theme.html"))
	}
}
