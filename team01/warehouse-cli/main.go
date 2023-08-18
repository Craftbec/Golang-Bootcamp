package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Error(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func parser(str string) []string {
	return strings.Split((str), " ")
}

func checkPort(port string) ([]string, error) {
	var (
		workport []string
		job      = []string{"8765", "9876", "8697"}
	)
	for _, val := range job {
		if val != port {
			r, err1 := http.Get(fmt.Sprintf("http://localhost:%s", val))
			if err1 == nil {
				workport = append(workport, val)
				r.Body.Close()
			}
		}
	}
	if len(workport) == 0 {
		return workport, errors.New("server is not availible now")
	}
	return workport, nil
}

func check(port string) error {
	r, err := http.Get(fmt.Sprintf("http://localhost:%s", port))
	if err == nil {
		r.Body.Close()
	}
	return err
}

func printnodes(host, port string) int {
	fmt.Println("Known nodes:")
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/list", host, port))
	Error(err)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	resp, err = http.Get(fmt.Sprintf("http://%s:%s/num", host, port))
	Error(err)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	numserverstart, _ := strconv.Atoi(string(body))
	if numserverstart == 1 {
		fmt.Println("WARNING: cluster size (1) is smaller than a replication factor (2)!")
	}
	return numserverstart
}

func connectPort(port string) error {
	r, err1 := http.Get(fmt.Sprintf("http://localhost:%s", port))
	if err1 != nil {
		return errors.New("no connection")
	} else {
		r.Body.Close()
	}
	return nil
}

func changePort(port *string, numserverstart *int) {
	for {
		time.Sleep(time.Second)
		err := connectPort(*port)
		if err != nil {
			p, er := checkPort(*port)
			if er != nil {
				*port = ""
				log.Fatal("All servers is down")
			}
			for _, val := range p {
				resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/asc", val))
				Error(err)
				defer resp.Body.Close()
				res, _ := io.ReadAll(resp.Body)
				if string(res) == "true" {
					*port = val
					break
				}
			}
			r, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/repl", *port))
			Error(err)
			if err == nil {
				r.Body.Close()
			}
			fmt.Printf("\nReconnected to a database of Warehouse 13 at 127.0.0.1:%s", *port)
		}
		num := num(*port)
		if num != *numserverstart {
			if *numserverstart == 1 {
				r, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/copy", *port))
				Error(err)
				if err == nil {
					r.Body.Close()
				}
			}
			if *numserverstart == 3 {
				r, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/rechange", *port))
				Error(err)
				if err == nil {
					r.Body.Close()
				}
			}
			numserverstart = &num
			fmt.Println("\nKnown nodes:")
			resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/list", *port))
			Error(err)
			io.Copy(os.Stdout, resp.Body)
			if err == nil {
				resp.Body.Close()
			}
			if num == 1 {
				fmt.Println("WARNING: cluster size (1) is smaller than a replication factor (2)!")
			}
			fmt.Print("> ")

		}
	}
}

func num(port string) int {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/num", port))
	Error(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	num, _ := strconv.Atoi(string(body))
	return num
}

func handler(port string, numserverstart int) {
	go changePort(&port, &numserverstart)
	var request string
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		sc.Scan()
		txt := sc.Text()
		str := parser(txt)
		switch str[0] {
		case "UUID":
			fmt.Println(uuid.New())
			continue
		case "TEST":
			request = "test"
		case "GET":
			if len(str) == 2 {
				request = fmt.Sprintf("get?uuid4=%s", str[1])
			}
		case "SET":
			if len(str) >= 3 {
				e := strings.Join(str[2:], " ")
				ur := url.QueryEscape(e)
				request = fmt.Sprintf("set?uuid4=%s&json=%s", str[1], ur)
			}
		case "DELETE":
			if len(str) == 2 {
				request = fmt.Sprintf("delete?uuid4=%s", str[1])
			}
		default:
			continue
		}
		if port != "" {
			resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/%s", port, request))
			Error(err)
			io.Copy(os.Stdout, resp.Body)
			fmt.Println()
			resp.Body.Close()
		}
	}
}

func main() {
	var (
		host, port string
	)
	flag.StringVar(&host, "H", "127.0.0.1", "-H <IP>")
	flag.StringVar(&port, "P", "8765", "-P <PORT>")
	flag.Parse()

	if err := check(port); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Connected to a database of Warehouse 13 at 127.0.0.1:%s\n", port)
	r, err := http.Get(fmt.Sprintf("http://%s:%s/start", host, port))
	Error(err)
	if err == nil {
		r.Body.Close()
	}
	num := printnodes(host, port)
	handler(port, num)
}
