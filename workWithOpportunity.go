package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func addOpportunity(id int, num int, name string) {

	connStr := "user=postgres password=postgres1234 dbname=crm sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	insert, err := db.Exec("INSERT INTO public.\"Opportunity\" VALUES ($1, $2, $3)", id, num, name)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(insert.RowsAffected())
}

func showOpportunities() []Opportunity {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	opportunities := []Opportunity{}

	query, err := db.Query("SELECT * FROM public.\"Opportunity\"")

	if err != nil {
		fmt.Println(err)
	}

	defer query.Close()

	for query.Next() {
		op := Opportunity{}

		err := query.Scan(&op.OppId, &op.OppNumber, &op.OppName)

		if err != nil {
			fmt.Println(err)
			continue
		}

		opportunities = append(opportunities, op)
	}

	showOppInfo(opportunities)

	fmt.Println(opportunities[0].OppId, opportunities[0].OppNumber, opportunities[0].OppName)

	return opportunities
}

func showOppInfo(op []Opportunity) {

	for _, v := range op {
		fmt.Printf("id: %d\nDeal: %s_%d\n\n", v.OppId, v.OppName, v.OppNumber)
	}
}

func updateOpportunity(oppotunity Opportunity) {

	connStr := "user=postgres password=postgres1234 dbname=crm sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	updateStatement := "UPDATE public.\"Opportunity\" SET \"OpportunityNumber\" = $1, \"OpportunityName\" = $2 WHERE \"index\" = $3"

	update, err := db.Exec(updateStatement, oppotunity.OppNumber, oppotunity.OppName, oppotunity.OppId)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(update.RowsAffected())
}

// need update
func deleteOpportunity(id int) {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	delete, err := db.Exec("DELETE FROM public.\"Opportunity\" WHERE \"index\"=$1", id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(delete.RowsAffected())
}

func getOpportunity(id int) Opportunity {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	query, err := db.Query("SELECT * FROM public.\"Opportunity\" WHERE index = $1", id)

	if err != nil {
		fmt.Println(err)
	}

	defer query.Close()

	var OppData Opportunity

	for query.Next() {

		err := query.Scan(&OppData.OppId, &OppData.OppNumber, &OppData.OppName)

		if err != nil {
			fmt.Println(err)
			continue
		}

	}

	//fmt.Println(OppData)

	return OppData
}
