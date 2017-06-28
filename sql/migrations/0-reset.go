package main

import (
	"examples/db"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() (err error) {
	db.Setup()
	db.Exec(`
		DROP DATABASE IF EXISTS projectnametest;
		CREATE DATABASE projectnametest;
		SET database = projectnametest;
	`)

	runSchemas("sql/schema.sql", "sql/test-data.sql")

	var company db.Company
	db.SelectOne(&company, `SELECT * FROM Company WHERE company_marker='acme'`)
	if company.CompanyId == 0 {
		panic("ASD")
	}
	return
}

func runSchemas(files ...string) {
	for _, file := range files {
		log.Println("Run schema:", file)
		lines := readSchema(file)
		for _, line := range lines {
			// log.Println(line)
			db.Exec(line)
		}
	}
}

func readSchema(filePath string) []string {
	bytes, err := ioutil.ReadFile(os.Getenv("PROJECTNAME_ROOT") + "/" + filePath)
	if err != nil {
		panic(err)
	}
	statements := strings.Split(string(bytes), ";")
	for i := range statements {
		statements[i] = strings.TrimSpace(statements[i])
	}
	return statements
}
