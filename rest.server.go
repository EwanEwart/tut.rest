//$GOROOT/bin/go run $0; $@; exit
/*
https://golangdocs.com/golang-rest-api
Moving on to the basic structure of the HTTP server which we shall be used to query from. This is the same as the one taught in our
https://golangdocs.com/http-server-in-golang
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Article ...
type Article struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

var Articles []Article
var nCalls uint

func main() {
	Articles = []Article{
		{
			"ONE IS ONE",
			"Lucy Irvine",
			"https://www.amazon.de/One-Novel-Lucy-Irvine/dp/0340501537/ref=sr_1_13?__mk_de_DE=%C3%85M%C3%85%C5%BD%C3%95%C3%91&dchild=1&keywords=%22one+is+one%22&qid=1626378224&s=books&sr=1-13",
		},
		{
			"ONE IS ONE",
			"Lucy Irvine",
			"https://www.amazon.de/One-Coronet-Books-Lucy-Irvine/dp/0340515775/ref=tmm_pap_swatch_0?_encoding=UTF8&qid=&sr=",
		},
		{
			"Runaway",
			"Lucy Irvine",
			"https://www.amazon.de/gp/product/B01F9R1V4C/ref=dbs_a_def_rwt_bibl_vppi_i6",
		},
		{
			"Python Intermediate and Advanced 101",
			"Arkaprabha Majumdar",
			"https://www.amazon.com/dp/B089KVK23P",
		},
		{
			"R programming Advanced",
			"Arkaprabha Majumdar",
			"https://www.amazon.com/dp/B089WH12CR",
		},
		{
			"R programming Fundamentals",
			"Arkaprabha Majumdar",
			"https://www.amazon.com/dp/B089S58WWG",
		},
		{
			"Misery",
			"Stephen King",
			"https://www.amazon.de/s?k=misery+stephen+king&hvadid=80058222341012&hvbmt=be&hvdev=c&hvqmt=e&tag=hyddemsn-21&ref=pd_sl_4e3mzlsjf0_e",
		},
	}
	nCalls = 0
	fmt.Println(http.DefaultServeMux)
	fmt.Println(os.Args[0])

	handleRequests()
}

/* handleRequests matches the hit URL path with a given handler function */
func handleRequests() {
	address := "localhost"
	address += ":"
	address += "8000"
	fmt.Printf("Server at address %v is up\n", address)

	http.HandleFunc("/", homePage)
	http.HandleFunc("/author", author)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(address, nil)) // nill: use DefaultServerMux
}

/* homePage manages all requests for the root URL */
func homePage(w http.ResponseWriter, r *http.Request) {
	nCalls += 1
	fmt.Fprintf(w, "Welcome to the tut.rest REST website!")
	fmt.Printf("Endpoint(%3v): REST website\n", nCalls)
}

/* author manages all requests for the root URL */
func author(w http.ResponseWriter, r *http.Request) {
	nCalls += 1
	fmt.Fprintf(w, "Author: Ewan Ewart")
	fmt.Printf("Endpoint(%3v): REST website author\n", nCalls)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	nCalls += 1
	fmt.Printf("Endpoint(%3v): REST website articles\n", nCalls)
	json.NewEncoder(w).Encode(Articles)

}
