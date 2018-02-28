package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
	"path"
)

var (
	Prefix         = "/mhc"
	Api1      = path.Join(Prefix, "1")
	Api2   = path.Join(Prefix, "2")
)


func Handler() http.Handler {
	handler1 := &MyAPI1{message:"haha"}
	handler2 := &MyAPI2{message:"lala"}
	mux := http.NewServeMux()
	mux.Handle(Api1, handler1)
	mux.Handle(Api2, handler2)
	return mux
}

type MyAPI1 struct {
	message string
}

func (h *MyAPI1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		if h.message != "" {
			w.Write([]byte(h.message))
		} else {
			w.Write([]byte("empty"))
		}
	default:
		w.Header().Add("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

type MyAPI2 struct {
	message string
}

func (h *MyAPI2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		if h.message != "" {
			w.Write([]byte(h.message))
		} else {
			w.Write([]byte("empty"))
		}
	default:
		w.Header().Add("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// serveHttpKVAPI starts a key-value server with a GET/PUT API and listens.
func serveHttpKVAPI(port int) {
	srv := http.Server{
		Addr: ":" + strconv.Itoa(port),
		Handler: Handler(),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("run...")
	serveHttpKVAPI(3322)
}


