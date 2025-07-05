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

func showOpportunities() {

	connStr := "user=postgres password=postgres1234 dbname=crm sslmode=disable"
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

		err := query.Scan(&op.oppId, &op.oppNumber, &op.oppName)

		if err != nil {
			fmt.Println(err)
			continue
		}

		opportunities = append(opportunities, op)
	}

	showOppInfo(opportunities)
}

func showOppInfo(op []Opportunity) {

	for _, v := range op {
		fmt.Printf("id: %d\nDeal: %s_%d\n\n", v.oppId, v.oppName, v.oppNumber)
	}
}

func updateOpportunity() {

	connStr := "user=postgres password=postgres1234 dbname=crm sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
}

func deleteOpportunity() {

	connStr := "user=postgres password=postgres1234 dbname=crm sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
}
