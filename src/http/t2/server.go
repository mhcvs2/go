package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"strings"
	"os"
)

type MyAPI struct {
	message string
}

func (h *MyAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "PUT":
		v, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read on PUT (%v)\n", err)
			http.Error(w, "Failed on PUT", http.StatusBadRequest)
			return
		}
		old := h.message
		h.message = string(v)
		var res string
		if old != h.message {
			res = fmt.Sprintf("change from '%s' to '%s'", old, h.message)
		} else {
			res = fmt.Sprintf("no change")
		}
		w.Write([]byte(res))
	case r.Method == "GET":
		if h.message != "" {
			w.Write([]byte(h.message))
		} else {
			w.Write([]byte("empty"))
		}
	case r.Method == "POST":
		v, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read on POST (%v)\n", err)
			http.Error(w, "Failed on POST", http.StatusBadRequest)
			return
		}
		cmd := string(v)
		if strings.ToLower(cmd) == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else {
			fmt.Printf("get command %s\n", cmd)
		}
		w.Write([]byte("success"))
	case r.Method == "DELETE":
		h.message = ""
		w.Write([]byte("success"))
	default:
		w.Header().Set("Allow", "PUT")
		w.Header().Add("Allow", "GET")
		w.Header().Add("Allow", "POST")
		w.Header().Add("Allow", "DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// serveHttpKVAPI starts a key-value server with a GET/PUT API and listens.
func serveHttpKVAPI(port int) {
	srv := http.Server{
		Addr: ":" + strconv.Itoa(port),
		Handler: &MyAPI{},
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("run...")
	serveHttpKVAPI(3322)
}
