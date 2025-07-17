package main

import (
	"database/sql"
	"fmt"
)

type Users struct {
	Id           int
	UserName     string
	UserLogin    string
	UserPassword string
}

func checkLoginInfo(login string, password string) bool {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	query, err := db.Query("SELECT * FROM public.\"Users\" WHERE \"UserLogin\" = $1", login)

	if err != nil {
		fmt.Println(err)
	}

	defer query.Close()

	var users Users

	for query.Next() {
		err = query.Scan(&users.Id, &users.UserName, &users.UserLogin, &users.UserPassword)

		if err != nil {
			fmt.Println(err)
		}
	}

	if users.Id == 0 {
		fmt.Println("No user with this login!")
		return false
	} else if users.UserPassword != password {
		fmt.Println("Password is incorrect!")
		return false
	} else {
		return true
	}

	//fmt.Println(users)
}
