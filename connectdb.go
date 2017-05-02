package main

import (
	"net/url"
	"os"
	"log"
	"strings"
	"errors"
	"fmt"
	"database/sql"
	"time"
	 _ "github.com/lib/pq"
)

type dbParams struct {
	username string
	password string
	host string
	port string
	db string
}

func fromEnvOrBust(varname string) string {
	val := os.Getenv(varname)
	if val == "" {
		log.Fatalf("%s not present in environment",varname)
	}
	return val
}

func extractDBParams() *dbParams {

	user := fromEnvOrBust("DB_USER")
	password := fromEnvOrBust("DB_PASSWORD")
	url := fromEnvOrBust("DB_URL")

	host,port,path, err := extractJDBCUrlParts(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &dbParams{
		username:user,
		password:password,
		host:host,
		port:port,
		db:path,
	}
}

func extractJDBCUrlParts(jdbcUrl string)(string,string,string, error) {
	parts := strings.SplitN(jdbcUrl, ":",2)
	if len(parts) != 2 {
		return "","","", errors.New("Only urls that start with jdbc: can be processed")
	}

	url,err := url.Parse(parts[1])
	if err != nil {
		return "","","",err
	}

	requestURI := url.RequestURI()
	if strings.HasPrefix(requestURI,"/") {
		parts := strings.SplitN(requestURI,"/",2)
		requestURI = parts[1]
	}

	return url.Hostname(), url.Port(), requestURI,nil
}

func connectDB(params *dbParams) bool {
	connectString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		params.username, params.password, params.db, params.host, params.port)
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Print(err.Error())
		return false
	}

	err = db.Ping()
	if err != nil {
		log.Print(err.Error())
		return false
	}

	var one int
	err = db.QueryRow("select 1").Scan(&one)
	if err != nil {
		log.Print(err.Error())
	}

	log.Printf("One is %d",one)

	return err == nil
}


func main() {
	connectParams := extractDBParams()
	for {
		log.Printf("check connection")
		if connectDB(connectParams) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}

}