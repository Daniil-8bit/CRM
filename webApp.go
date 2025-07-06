package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func startWebApp() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/addNewOpportunity", newOpportunityHandler)
	http.ListenAndServe(":4545", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	//w.Write([]byte("Hello world!"))

	t, err := template.ParseFiles("addOpportunity.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

func newOpportunityHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("New opportunity added successfuly"))

	err := r.ParseForm()

	if err != nil {
		fmt.Println(err)
	}

	var newOpp Opportunity

	oppId64, err := strconv.ParseInt(r.PostForm.Get("oppId"), 10, 64)
	newOpp.oppId = int(oppId64)
	oppNum64, err := strconv.ParseInt(r.PostForm.Get("oppNumber"), 10, 64)
	newOpp.oppNumber = int(oppNum64)
	newOpp.oppName = r.PostForm.Get("oppName")

	fmt.Println(newOpp.oppId, newOpp.oppNumber, newOpp.oppName)

	addOpportunity(newOpp.oppId, newOpp.oppNumber, newOpp.oppName)
}
