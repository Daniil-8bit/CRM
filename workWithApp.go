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

type DBTables struct {
	TableName string
}

func getDBObjects() []DBTables {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	query, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema NOT IN ('information_schema', 'pg_catalog')")

	if err != nil {
		fmt.Println(err)
	}

	var dbTable DBTables
	var dbTables []DBTables

	for query.Next() {

		err := query.Scan(&dbTable.TableName)

		if err != nil {
			fmt.Println(err)
		}

		dbTables = append(dbTables, dbTable)
	}

	return dbTables

	//fmt.Println(dbTables)
}

type DBTableInfo struct {
	TableName string
	Fields    map[string]string
}

func getDBObjectsData(tableName string) DBTableInfo {

	var cd ConfigData = readConfigFile()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cd.DbUser, cd.DbUserPassword, cd.DbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	query, err := db.Query("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = $1", tableName)

	if err != nil {
		fmt.Println(err)
	}

	var dbTableInfo DBTableInfo

	var fieldName string
	var fieldType string

	fields := make(map[string]string)

	for query.Next() {

		err := query.Scan(&fieldName, &fieldType)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fieldName)
		fmt.Println(fieldType)

		fields[fieldName] = fieldType
	}

	dbTableInfo.TableName = tableName
	dbTableInfo.Fields = fields

	fmt.Println(dbTableInfo)

	return dbTableInfo
}

func changeDBObjectsData(tableName string, fieldName string, fieldType string) {

}
