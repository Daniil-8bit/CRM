package main

import (
	"database/sql"
	"fmt"
)

func showLeads() []Lead {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	leads := []Lead{}

	query, err := db.Query("SELECT * FROM public.\"Lead\"")

	if err != nil {
		fmt.Println(err)
	}

	defer query.Close()

	for query.Next() {
		lead := Lead{}

		err := query.Scan(&lead.LeadId, &lead.LeadName, &lead.LeadSource)

		if err != nil {
			fmt.Println(err)
			continue
		}

		leads = append(leads, lead)
	}

	return leads
}

func updateLead(lead Lead) {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	updateStatement := "UPDATE public.\"Lead\" SET \"LeadName\" = $1, \"LeadSource\" = $2 WHERE \"LeadId\" = $3"

	update, err := db.Exec(updateStatement, lead.LeadName, lead.LeadSource, lead.LeadId)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(update.RowsAffected())
}

func getLead(id int) Lead {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	query, err := db.Query("SELECT * FROM public.\"Lead\" WHERE \"LeadId\" = $1", id)

	if err != nil {
		fmt.Println(err)
	}

	defer query.Close()

	var LeadData Lead

	for query.Next() {

		err := query.Scan(&LeadData.LeadId, &LeadData.LeadName, &LeadData.LeadSource)

		if err != nil {
			fmt.Println(err)
			continue
		}

	}

	return LeadData
}

func addLead(leadId int, leadName string, leadSource string) {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	insert, err := db.Exec("INSERT INTO public.\"Lead\" VALUES ($1, $2, $3)", leadId, leadName, leadSource)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(insert.RowsAffected())
}

func deleteLead(leadId int) {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	delete, err := db.Exec("DELETE FROM public.\"Lead\" WHERE \"LeadId\"=$1", leadId)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(delete.RowsAffected())
}
