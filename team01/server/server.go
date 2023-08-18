package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	db               = make(map[string]string)
	current, replica int
	job              []int
	duplicate        bool
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func get(w http.ResponseWriter, r *http.Request) {
	uuid4 := r.URL.Query().Get("uuid4")
	if v, ok := db[uuid4]; ok {
		fmt.Fprintf(w, "%s", v)
	} else {
		fmt.Fprintf(w, "Not found")
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	uuid4 := r.URL.Query().Get("uuid4")
	if !isValidUUID((uuid4)) {
		fmt.Fprintf(w, "Error: Key is not a proper UUID4")
	} else {
		json := r.URL.Query().Get("json")
		db[uuid4] = json
		if replica != 0 {
			ur := url.QueryEscape(json)
			request := fmt.Sprintf("set?uuid4=%s&json=%s", uuid4, ur)
			r, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s", replica, request))
			if err == nil {
				r.Body.Close()
			}
			fmt.Fprintf(w, "Created (2 replicas)")
		}

	}
}

func del(w http.ResponseWriter, r *http.Request) {
	uuid4 := r.URL.Query().Get("uuid4")
	delete(db, uuid4)
	if replica != 0 {
		request := fmt.Sprintf("delete?uuid4=%s", uuid4)
		r, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s", replica, request))
		if err == nil {
			r.Body.Close()
		}
		fmt.Fprintf(w, "Deleted (2 replicas)")
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	for key, value := range db {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(w); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}

func num(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", len(job))
}

func asc(w http.ResponseWriter, r *http.Request) {
	if duplicate {
		w.Write([]byte("true"))
	}
}

func repl(w http.ResponseWriter, r *http.Request) {
	duplicate = false
	if len(job) > 1 {
		setReplica()
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	db = make(map[string]string)
	repl(w, r)
}

func setReplica() {
	for _, val := range job {
		if val != current {
			replica = val
			break
		}
	}
	if replica != 0 {
		r, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/setRepl", replica))
		if err == nil {
			r.Body.Close()
		}
		copyDB()
	}
}

func copyDB() {
	for key, val := range db {
		e := url.QueryEscape(val)
		request := fmt.Sprintf("set?uuid4=%s&json=%s", key, e)
		r, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s", replica, request))
		if err == nil {
			r.Body.Close()
		}

	}
}

func copy(w http.ResponseWriter, r *http.Request) {
	setReplica()
}

func setRepl(w http.ResponseWriter, r *http.Request) {
	duplicate = true
}

func rechange(w http.ResponseWriter, r *http.Request) {
	resp, err1 := http.Get(fmt.Sprintf("http://localhost:%d", replica))
	if err1 != nil {
		setReplica()
	} else {
		resp.Body.Close()
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	var list string
	for i := range job {
		list += "127.0.0.1:" + strconv.Itoa(job[i]) + "\n"
	}
	fmt.Fprintf(w, "%v", list)
}

func checkPorts(port int) []int {
	var job = []int{}
	client := http.Client{}
	switch port {
	case 8765:
		job = append(job, 8765)
		r, err1 := client.Get("http://localhost:9876")
		if err1 == nil {
			job = append(job, 9876)
			r.Body.Close()
		}
		rr, err2 := client.Get("http://localhost:8697")
		if err2 == nil {
			job = append(job, 8697)
			rr.Body.Close()
		}
	case 9876:
		r, err1 := client.Get("http://localhost:8765")
		if err1 == nil {
			job = append(job, 8765)
			r.Body.Close()
		}
		job = append(job, 9876)
		rr, err2 := client.Get("http://localhost:8697")
		if err2 == nil {
			job = append(job, 8697)
			rr.Body.Close()
		}
	case 8697:
		r, err1 := client.Get("http://localhost:8765")
		if err1 == nil {
			job = append(job, 8765)
			r.Body.Close()
		}
		rr, err2 := client.Get("http://localhost:9876")
		if err2 == nil {
			job = append(job, 9876)
			rr.Body.Close()
		}
		job = append(job, 8697)
	}
	return job
}

func main() {
	flag.IntVar(&current, "P", 8765, "-P (8765, 9876 or 8697)")
	flag.Parse()
	main_server := http.NewServeMux()
	main_server.HandleFunc("/num", num)
	main_server.HandleFunc("/list", list)
	main_server.HandleFunc("/get", get)
	main_server.HandleFunc("/set", set)
	main_server.HandleFunc("/delete", del)
	main_server.HandleFunc("/test", test)
	main_server.HandleFunc("/repl", repl)
	main_server.HandleFunc("/copy", copy)
	main_server.HandleFunc("/rechange", rechange)
	main_server.HandleFunc("/asc", asc)
	main_server.HandleFunc("/start", start)
	main_server.HandleFunc("/setRepl", setRepl)
	go func() {
		for {
			time.Sleep(time.Second)
			job = checkPorts(current)
		}
	}()
	log.Printf("Starting server on http://localhost:%d ...\n", current)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", current), main_server); err != nil {
		log.Fatal(err)
	}
}
