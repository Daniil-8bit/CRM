package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func startWebApp() {

	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/Login", loginPageHandler)
	http.HandleFunc("/LoginCheck", loginCheckHandler)
	http.HandleFunc("/createOpportunity", createOpportunityHandler)
	http.HandleFunc("/addNewOpportunity", newOpportunityHandler)
	http.HandleFunc("/allOpportunities", allOpportunitiesHandler)
	http.HandleFunc("/Opportunity", getOpportunityHandler)
	http.HandleFunc("/update", updateOpportunityHandler)
	http.HandleFunc("/delete", deleteOpportunityHandler)
	http.HandleFunc("/constructor", showAllObjects)
	http.HandleFunc("/editObject", showObjectInfo)
	http.HandleFunc("/changeField", changeFieldObjectInfo)
	http.HandleFunc("/allLeads", allLeadsHandler)
	http.HandleFunc("/updateLead", updateLeadHandler)
	http.HandleFunc("/Lead", getLeadHandler)
	http.HandleFunc("/createLead", createLeadHandler)
	http.HandleFunc("/addNewLead", newLeadHandler)
	http.HandleFunc("/LeadDelete", deleteLeadHandler)
	http.ListenAndServe(":4545", nil)
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("mainPage.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

func createOpportunityHandler(w http.ResponseWriter, r *http.Request) {

	//w.Write([]byte("Hello world!"))

	t, err := template.ParseFiles("addOpportunity.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

func newOpportunityHandler(w http.ResponseWriter, r *http.Request) {

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

	http.Redirect(w, r, "/allOpportunities", http.StatusFound)
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

func deleteOpportunityHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	deleteOpportunity(int(id))

	http.Redirect(w, r, "/allOpportunities", http.StatusMovedPermanently)
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

type ObjectsView struct {
	SystemObjects []DBTables
}

func showAllObjects(w http.ResponseWriter, r *http.Request) {

	var SystemObjects []DBTables = getDBObjects()

	data := ObjectsView{
		SystemObjects: SystemObjects,
	}

	fmt.Println("System objects: ", SystemObjects)

	tmpl, err := template.ParseFiles("constructor.html")

	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, data)
}

func showObjectInfo(w http.ResponseWriter, r *http.Request) {

	tableName := r.URL.Query().Get("tableName")

	var tableInfo DBTableInfo = getDBObjectsData(tableName)

	tmpl, err := template.ParseFiles("editObject.html")

	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, tableInfo)
}

type FieldObjectInfo struct {
	TableName string
	FieldName string
	FieldType string
}

func changeFieldObjectInfo(w http.ResponseWriter, r *http.Request) {

	tableName := r.URL.Query().Get("tableName")
	fieldName := r.URL.Query().Get("fieldName")
	fieldType := r.URL.Query().Get("fieldType")

	changeField := FieldObjectInfo{
		TableName: tableName,
		FieldName: fieldName,
		FieldType: fieldType,
	}

	tmpl, err := template.ParseFiles("editField.html")

	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, changeField)

	http.Redirect(w, r, "/constructor", http.StatusFound)
}

type LeadsView struct {
	Leads []Lead
}

func allLeadsHandler(w http.ResponseWriter, r *http.Request) {

	var LeadsArray []Lead = showLeads()

	data := LeadsView{
		Leads: LeadsArray,
	}

	t, err := template.ParseFiles("allLeads.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, data)
}

func updateLeadHandler(w http.ResponseWriter, r *http.Request) {

	leadId, err := strconv.ParseInt(r.FormValue("leadId"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	leadName := r.FormValue("leadName")
	leadSource := r.FormValue("leadSource")

	lead := Lead{
		LeadId:     int(leadId),
		LeadName:   leadName,
		LeadSource: leadSource,
	}

	updateLead(lead)

	http.Redirect(w, r, "/allLeads", http.StatusFound)
}

func getLeadHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	LeadId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	var OneLead Lead = getLead(int(LeadId))

	t, err := template.ParseFiles("infoLead.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, OneLead)
}

func createLeadHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("addLead.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

func newLeadHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		fmt.Println(err)
	}

	var newLead Lead

	oppId64, err := strconv.ParseInt(r.PostForm.Get("leadId"), 10, 64)
	newLead.LeadId = int(oppId64)
	newLead.LeadName = r.PostForm.Get("leadName")
	newLead.LeadSource = r.PostForm.Get("leadSource")

	addLead(newLead.LeadId, newLead.LeadName, newLead.LeadSource)

	http.Redirect(w, r, "/allLeads", http.StatusFound)
}

func deleteLeadHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	deleteLead(int(id))

	http.Redirect(w, r, "/allLeads", http.StatusFound)
}
