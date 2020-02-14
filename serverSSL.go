package main

import (
	"crypto/tls"
	"flag"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

// copied from https://gist.github.com/samthor/5ff8cfac1f80b03dfe5a9be62b29d7f2

func main() {

	path := flag.String("path", "www", "server path")
	domains := flag.String("domains", "", "eg : 'www.acme.com,foo.bar.com'")
	flag.Parse()

	_, err := os.Stat(*path)
	if !os.IsNotExist(err) {

		domainz := strings.Split(*domains, ",")

		// create the autocert.Manager with domains and path to the cache
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(domainz...),
		}

		// optionally use a cache dir
		dir := cacheDir()
		if dir != "" {
			certManager.Cache = autocert.DirCache(dir)
		}

		// create the server itself
		server := &http.Server{
			Addr:    ":https",
			Handler: http.FileServer(http.Dir(*path)),
			TLSConfig: &tls.Config{
				MinVersion:     tls.VersionTLS12,
				GetCertificate: certManager.GetCertificate,
			},
		}

		log.Printf("Serving http/https for domains: %+v", strings.Join(domainz, ","))
		go func() {
			// serve HTTP, which will redirect automatically to HTTPS
			h := certManager.HTTPHandler(nil)
			log.Fatal(http.ListenAndServe(":http", h))
		}()

		// serve HTTPS!
		log.Fatal(server.ListenAndServeTLS("", ""))

	} else {
		log.Printf("path: %s does not exists", *path)
	}

}

func cacheDir() (dir string) {
	if u, _ := user.Current(); u != nil {
		dir = filepath.Join(os.TempDir(), "cache-golang-autocert-"+u.Username)
		if err := os.MkdirAll(dir, 0700); err == nil {
			return dir
		}
	}
	return ""
}
