package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

)

type dollars float32 

func (d dollars) String() string{
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request){

	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}
	p, err := strconv.ParseFloat(price, 32)

	if  err != nil{
		msg := fmt.Sprintf("invalid price : %q", price)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}

	db[item] = dollars(p)

	fmt.Fprintf(w, "added %s with price %s\n", item, db[item])
}


func (db database) update(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}
	p, err := strconv.ParseFloat(price, 32)

	if  err != nil{
		msg := fmt.Sprintf("invalid price : %q", price)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}

	db[item] = dollars(p)

	fmt.Fprintf(w, "new price %s for price %s\n", db[item], item)
}


func (db database) fetch(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")

		if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}
	fmt.Fprintf(w, "item %s has price %s\n", item, db[item])

}

func (db database) drop(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")

		if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "dropped %s\n", item)

}


func (db database) rename(w http.ResponseWriter, req *http.Request){
	oldName := req.URL.Query().Get("old")
	newName := req.URL.Query().Get("new")
	price := req.URL.Query().Get("price")


	if _, ok := db[oldName]; !ok {
		msg := fmt.Sprintf("no such item: %q", oldName)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	if _, exists := db[newName]; exists {
		http.Error(w, fmt.Sprintf("item %q already exists", newName), http.StatusBadRequest)
		return
	}

	p, err := strconv.ParseFloat(price, 32)

		if err != nil {
		http.Error(w, fmt.Sprintf("invalid price: %q", price), http.StatusBadRequest)
		return
	}
	db[newName] = dollars(p)
	delete(db, oldName)

	fmt.Fprintf(w, "renamed its %q to %q with price %s\n", oldName, newName, db[newName])
	
}

func main() {
	db := database {
		"shoes": 50,
		"socks":5,
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.drop)
	http.HandleFunc("/read", db.fetch)
	http.HandleFunc("/rename", db.rename)

	log.Fatal(http.ListenAndServe(":8080", nil))
}