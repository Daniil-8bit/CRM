package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func startWebApp() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/Login", loginPageHandler)
	http.HandleFunc("/LoginCheck", loginCheckHandler)
	http.HandleFunc("/addNewOpportunity", newOpportunityHandler)
	http.HandleFunc("/allOpportunities", allOpportunitiesHandler)
	http.HandleFunc("/Opportunity", getOpportunityHandler)
	http.HandleFunc("/update", updateOpportunityHandler)
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
	newOpp.OppId = int(oppId64)
	oppNum64, err := strconv.ParseInt(r.PostForm.Get("oppNumber"), 10, 64)
	newOpp.OppNumber = int(oppNum64)
	newOpp.OppName = r.PostForm.Get("oppName")

	fmt.Println(newOpp.OppId, newOpp.OppNumber, newOpp.OppName)

	addOpportunity(newOpp.OppId, newOpp.OppNumber, newOpp.OppName)
}

type OpportunitiesView struct {
	Opportunities []Opportunity
}

type TestStruct struct {
	Name string
	Num  int
}

type TestView struct {
	TestStrings []TestStruct
}

func allOpportunitiesHandler(w http.ResponseWriter, r *http.Request) {

	var OppArray []Opportunity = showOpportunities()

	data := OpportunitiesView{

		Opportunities: OppArray,
	}

	/*data1 := OpportunitiesView{

		Opportunities: []Opportunity{
			Opportunity{oppId: 111222333, oppNumber: 333222111, oppName: "New try!"},
			Opportunity{oppId: 111222332, oppNumber: 333222112, oppName: "New try1!"},
			Opportunity{oppId: 111222331, oppNumber: 333222113, oppName: "New try2!"},
		},
	}*/

	/*data2 := TestView{
		TestStrings: []TestStruct{
			TestStruct{"Test1", 1},
			TestStruct{"Test2", 2},
			TestStruct{"Test3", 3},
		},
	}*/

	t, err := template.ParseFiles("allOpportunities.html")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Opportunities[0].OppId)

	t.Execute(w, data)

}

func getOpportunityHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	OpportunityId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	var OneOpportunity Opportunity = getOpportunity(int(OpportunityId))

	fmt.Println(OneOpportunity)

	//fmt.Fprint(w, "Opportunity: ", id)

	fmt.Println("id: ", id)

	t, err := template.ParseFiles("infoOpportunity.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, OneOpportunity)
}

func updateOpportunityHandler(w http.ResponseWriter, r *http.Request) {

	oppId, err := strconv.ParseInt(r.FormValue("oppId"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	oppNumber, err := strconv.ParseInt(r.FormValue("oppNumber"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	oppName := r.FormValue("oppName")

	//fmt.Fprintf(w, "This is update page!\n\nID: %d\nNumber: %d\nName: %s", oppId, oppNumber, oppName)
	//fmt.Fprint(w, oppNumber)

	opp := Opportunity{
		OppId:     int(oppId),
		OppNumber: int(oppNumber),
		OppName:   oppName,
	}

	updateOpportunity(opp)

	//fmt.Printf("Number: %d\nName: %s\n", oppNumber, oppName)

	http.Redirect(w, r, "/allOpportunities", http.StatusFound)
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles("loginPage.html")

	if err != nil {
		fmt.Println(err)
	}

	templ.Execute(w, nil)
}

func loginCheckHandler(w http.ResponseWriter, r *http.Request) {

	login := r.PostFormValue("login")
	password := r.PostFormValue("password")

	ok := checkLoginInfo(login, password)

	if ok {
		http.Redirect(w, r, "/allOpportunities", http.StatusFound)
	} else {
		fmt.Fprint(w, "Data is incorrect!")
	}
}
