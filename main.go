package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"log"
	"net/http"
)

func main()  {
	config.ConnectDB()


	// 1. panggil home page
	http.HandleFunc("/", homecontroller.Welcome);

	// 2. panggil category page
	http.HandleFunc("/categories", categorycontroller.Index);
	// http.HandleFunc("/categories/create", categorycontroller.Create);
	http.HandleFunc("/categories/add", categorycontroller.Add);
	http.HandleFunc("/categories/edit", categorycontroller.Edit);
	// http.HandleFunc("/categories/update", categorycontroller.Update);
	http.HandleFunc("/categories/delete", categorycontroller.Delete);

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}