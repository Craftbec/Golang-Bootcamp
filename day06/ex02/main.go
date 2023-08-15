package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/didip/tollbooth/v7"
	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type Basa struct {
	Id        uint16
	Content   template.HTML
	Full_text template.HTML
}

type Login struct {
	Login string
	Pas   string
}

type Str struct {
	Products []Basa
	Pg       []int
}

var products = []Basa{}
var showPost = Basa{}

func index(w http.ResponseWriter, r *http.Request) {
	var page float64
	var pg = []int{}
	num := 1
	tmp, _ := r.URL.Query()["page"]
	if len(tmp) > 0 {
		num, _ = strconv.Atoi(tmp[0])
	}
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	connStr := "user=craftbec password= dbname=basa sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT COUNT(*) FROM basa")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&page)
		if err != nil {
			fmt.Println(err)
		}
	}
	page = math.Ceil(page / 3)
	for i := 1; i <= int(page); i++ {
		pg = append(pg, i)
	}
	rows, err = db.Query(fmt.Sprintf("select * from basa ORDER BY id ASC LIMIT 3 OFFSET %d;", num*3-3))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	products = []Basa{}
	for rows.Next() {
		var p Basa
		err = rows.Scan(&p.Id, &p.Content, &p.Full_text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	t.ExecuteTemplate(w, "index", Str{products, pg})
}

func admin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/admin.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "admin", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "login", nil)
}

func chek_login(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("login") == "craftbec" && r.FormValue("pas") == "1234" {
		http.Redirect(w, r, "/addition", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func save_article(w http.ResponseWriter, r *http.Request) {
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")
	connStr := "user=craftbec password= dbname=basa sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	html := markdown.ToHTML([]byte(full_text), nil, nil)
	htmla := markdown.ToHTML([]byte(anons), nil, nil)
	insert, err := db.Query(fmt.Sprintf("INSERT INTO basa (cont, full_text)\n"+
		"VALUES ('%s', '%s')", htmla, html))
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func show_post(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	vars := mux.Vars(r)
	connStr := "user=craftbec password= dbname=basa sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(fmt.Sprintf("select * from basa where id = %s", vars["id"]))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	showPost = Basa{}
	for rows.Next() {
		var p Basa
		err := rows.Scan(&p.Id, &p.Content, &p.Full_text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		showPost = p
	}
	t.ExecuteTemplate(w, "show", showPost)
}

func handleRequest() {

	rtr := mux.NewRouter()
	lmt := tollbooth.NewLimiter(100, nil)
	lmt.SetMessage("429 Too Many Requests")
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	rtr.Handle("/", tollbooth.LimitFuncHandler(lmt, index)).Methods("GET")
	rtr.Handle("/admin", tollbooth.LimitFuncHandler(lmt, login)).Methods("POST", "GET")
	rtr.Handle("/addition", tollbooth.LimitFuncHandler(lmt, admin)).Methods("GET")
	rtr.Handle("/save_article", tollbooth.LimitFuncHandler(lmt, save_article)).Methods("POST")
	rtr.Handle("/chek_login", tollbooth.LimitFuncHandler(lmt, chek_login)).Methods("POST")
	rtr.Handle("/post/{id:[0-9]+}", tollbooth.LimitFuncHandler(lmt, show_post)).Methods("GET")
	http.Handle("/", rtr)
	http.ListenAndServe(":8888", nil)

}

func main() {
	exec.Command("unzip", "arc.zip").Run()
	handleRequest()
}
