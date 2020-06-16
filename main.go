package main

import (
	"database/sql"

	"fmt"

	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type dictSQL struct {
	id       int
	word     string
	meaning1 string
	meaning2 string
	meaning3 string
	meaning4 string
	meaning5 string
}

var dictSQLX []dictSQL

// Read the command line
func getInput() string {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a word to search for")
	}
	var words string
	for _, v := range os.Args[1:] {
		words += v + " "
	}
	words = strings.TrimSuffix(words, " ")
	return words

}

func main() {

	// connect to mysql
	connectStr := "root:Pinepine01#@tcp(127.0.0.1:3306)/dictionary?charset=utf8mb4"
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		log.Fatalf("Could not connect to msql %v\n", err)
	}
	defer db.Close()
	fmt.Println("db is connected")
	// test connection
	err = db.Ping()
	if err != nil {
		log.Printf("Db not responding %v\n", err)
	}
	fmt.Println("db is available")

	// get work from user
	sWord := getInput()

	// query the database
	// Execute the query
	results, err := db.Query("SELECT * FROM dict WHERE word=?", sWord)
	if err != nil {
		log.Printf("Cannot find %s in db: %v\n", sWord, err)
	}

	for results.Next() {
		dictItem := dictSQL{}
		// for each row, scan the result into our tag composite object
		err = results.Scan(&dictItem.id, &dictItem.word, &dictItem.meaning1, &dictItem.meaning2, &dictItem.meaning3, &dictItem.meaning4, &dictItem.meaning5)
		if err != nil {
			log.Printf("Cannot retrieve from db: %v\n", err)
		}
		fmt.Printf("You searched for the meaning of (%s):\n", sWord)
		fmt.Println("Possible meanings: ")
		fmt.Printf(" %s\n %s\n %s\n %s\n %s\n", dictItem.meaning1, dictItem.meaning2, dictItem.meaning3, dictItem.meaning4, dictItem.meaning5)
	}
}
