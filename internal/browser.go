package internal

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func CatchRedirect(port string, ch chan bool) {
	handler := func(chan bool) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ch <- true
			w.Write([]byte("You can close this window now"))
		})
	}
	http.Handle("/", handler(ch))

	err := http.ListenAndServe(port, nil)

	if err != nil {
		panic(err)
	}
}

func CatchYNABRedirect(port string, ch chan string) {
	handler := func(chan string) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.URL.Fragment
			ch <- token
			w.Write([]byte("You can close this window now"))
		})
	}
	http.Handle("/ynab", handler(ch))

	err := http.ListenAndServe(port, nil)

	if err != nil {
		panic(err)
	}
}
