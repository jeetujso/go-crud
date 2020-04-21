package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type bookStruct struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

// type bookStruct struct {
// 	ID     int
// 	Title  string
// 	Author string
// 	Year   string
// }

var books []bookStruct //slice of bookStruct struct

func main() {
	// single struct data
	// b1 := bookStruct{
	// 	ID:     1,
	// 	Title:  "Data Structure",
	// 	Author: "Jams Bond",
	// 	Year:   "2020",
	// }
	//books := []bookStruct{}
	books = append(
		books,
		bookStruct{ID: 1, Title: "MERN Stack", Author: "Jams Bond", Year: "2009"},
		bookStruct{ID: 2, Title: "MEAN Stack", Author: "Jams", Year: "2020"},
		bookStruct{ID: 3, Title: "React", Author: "Bond", Year: "2001"},
		bookStruct{ID: 4, Title: "Angular", Author: "John", Year: "20200"},
		bookStruct{ID: 5, Title: "Node", Author: "APJ", Year: "2010"},
		bookStruct{ID: 6, Title: "Express", Author: "Expr", Year: "2000"},
	)

	r := mux.NewRouter() //invocking new routes
	//define routes
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/add-book", addBooks).Methods("POST")
	r.HandleFunc("/update", updateBook).Methods("PUT")
	r.HandleFunc("/remove/{id}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":8000", r) //listen the server at 8000
	fmt.Println("Starting server on the port 8000...")
}

//define functiontion

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home called")
}
func getBooks(w http.ResponseWriter, r *http.Request) { //all books
	//return an object of slice of book of type bookStruct struct type
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) { // single book by passing id
	params := mux.Vars(r)
	log.Println(params)                       // map[id:3]
	log.Println(reflect.TypeOf(params["id"])) // string
	// conver string to int
	i, _ := strconv.Atoi(params["id"])
	log.Println(reflect.TypeOf(i)) // int

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(book)
			//json.NewEncoder(w).Encode(&book)
		}
	}
}
func addBooks(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("addBooks called")
	//var newBook bookStruct
	newBook := bookStruct{} // assign newBook of stuct of types bookStruct
	//log.Println(newBook)    //{0   }
	json.NewDecoder(r.Body).Decode(&newBook) // address of newBook
	books = append(books, newBook)
	json.NewEncoder(w).Encode(books)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	bookList := bookStruct{}
	json.NewDecoder(r.Body).Decode(&bookList)
	for index, iteam := range books {
		if iteam.ID == bookList.ID {
			books[index] = bookList
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			//json.NewEncoder(w).Encode(&book)
		}
	}
}
