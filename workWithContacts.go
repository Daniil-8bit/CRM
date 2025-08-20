package main

import (
	"database/sql"
	"fmt"
)

func showContacts() []Contact {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	contacts := []Contact{}

	query, err := db.Query("SELECT * FROM public.\"Contact\"")

	if err != nil {
		fmt.Println(err)
	}

	defer query.Close()

	for query.Next() {
		contact := Contact{}

		err := query.Scan(&contact.ContactId, &contact.ContactSurname, &contact.ContactName, &contact.ContactMiddlename, &contact.ContactPhone, &contact.ContactEmail, &contact.ContactJobTitle)

		if err != nil {
			fmt.Println(err)
			continue
		}

		contacts = append(contacts, contact)
	}

	return contacts
}

func updateContact(contact Contact) {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	updateStatement := "UPDATE public.\"Contact\" SET \"ContactSurname\" = $1, \"ContactName\" = $2, \"ContactMiddlename\" = $3, \"ContactPhone\" = $4, \"ContactEmail\" = $5, \"ContactJobTitle\" = $6 WHERE \"ContactId\" = $7"

	update, err := db.Exec(updateStatement, contact.ContactSurname, contact.ContactName, contact.ContactMiddlename, contact.ContactPhone, contact.ContactEmail, contact.ContactJobTitle, contact.ContactId)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(update.RowsAffected())
}

func getContact(contactId int) Contact {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	query, err := db.Query("SELECT * FROM public.\"Contact\" WHERE \"ContactId\" = $1", contactId)

	if err != nil {
		fmt.Println(err)
	}

	defer query.Close()

	var ContactData Contact

	for query.Next() {

		err := query.Scan(&ContactData.ContactId, &ContactData.ContactSurname, &ContactData.ContactName, &ContactData.ContactMiddlename, &ContactData.ContactPhone, &ContactData.ContactEmail, &ContactData.ContactJobTitle)

		if err != nil {
			fmt.Println(err)
			continue
		}

	}

	return ContactData
}

func addContact(contactId int, contactSurname string, contactName string, contactMiddlename string, contactPhone string, contactEmail string, contactJobTitle string) {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	insertStatement := "INSERT INTO public.\"Contact\" VALUES ($1, $2, $3, $4, $5, $6, $7)"

	insert, err := db.Exec(insertStatement, contactId, contactSurname, contactName, contactMiddlename, contactPhone, contactEmail, contactJobTitle)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(insert.RowsAffected())
}

func deleteContact(contactId int) {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	delete, err := db.Exec("DELETE FROM public.\"Contact\" WHERE \"ContactId\"=$1", contactId)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(delete.RowsAffected())
}
