package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Price       string `json:"Price"`
}

var Products []Product

func returnProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Products")
	json.NewEncoder(w).Encode(Products)
}

func returnProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Endpoint Hit: Product ID ", key)
	for _, product := range Products {
		if product.Id == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func createNewProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)
	Products = append(Products, product)
	fmt.Println("Endpoint Hit: Create New Product ID ", product.Id)
	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Delete Product")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, product := range Products {
		if product.Id == id {
			Products = append(Products[:index], Products[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnProducts).Methods("GET")
	myRouter.HandleFunc("/product/{id}", returnProduct).Methods("GET")
	myRouter.HandleFunc("/product", createNewProduct).Methods("POST")
	myRouter.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("App Started")
	Products = []Product{
		{Id: "1", Name: "Producto 1", Description: "Descripción 1", Price: "9999"},
		{Id: "2", Name: "Producto 2", Description: "Descripción 2", Price: "9998"},
	}
	handleRequests()
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	fmt.Println("Error 404")
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "error 404")
	}
}
