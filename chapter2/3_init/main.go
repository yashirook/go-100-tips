package main

import (
	"fmt"
	_ "init/pkg/foo"
	"init/pkg/redis"
	"net/http"
)

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init first")
}

func init() {
	fmt.Println("init second")
}

func init() {
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/favicon.ico", static)
	http.Handle("/fonts.css", static)
	http.Handle("/fonts/", static)

	http.Handle("/lib/godoc", http.StripPrefix("/lib/godoc", http.HandlerFunc(staticHander)))
}

func staticHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func main() {
	fmt.Println("main")
	redis.Store("test", "testvalue")
}
