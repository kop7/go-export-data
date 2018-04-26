package main

import (
	"database/sql"
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 "log"
	 "encoding/csv"
	 "os"
)

func getUserData(i int)(string, string){

db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testdb")
	 if err != nil {
		panic(err.Error())
		}

	defer db.Close()

		rows, err := db.Query("SELECT email,username FROM users WHERE id=?", i)
		if err != nil {
		        log.Fatal(err)
		}
		defer rows.Close()

			var (
				email string
				username string
				)

		for rows.Next() {

		        if err := rows.Scan(&email, &username); err != nil {
                log.Fatal(err)
       			 }
		}

 		return email, username
}


func main() {

		file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
		defer file.Close()

		if err != nil {
			os.Exit(1)
		}

		for i := 1; i <= 10; i++ {

         email, username := getUserData(i)
         fmt.Printf("%s -> %s is %s\n",i, email, username)

		data := []string{email, username}

		csvWriter := csv.NewWriter(file)
		strWrite := [][]string{data}
		csvWriter.WriteAll(strWrite)
		csvWriter.Flush()

    	}
	}

