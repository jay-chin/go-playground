package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

// Define flags on cmd line
var (
	debugPtr = flag.Bool("d", true, "enable debugging")
	password = flag.String("p", "password", "Database password")
	portPtr  = flag.Int("t", 1433, "Database port")
	server   = flag.String("s", "localhost", "Database server")
	user     = flag.String("u", "user", "Database user")
	rowsPtr  = flag.Int("r", 100, "Number of rows to retrieve")
	seqPtr   = flag.Int("q", 0, "Get rows greater than this sequence number")
)

func init() {
	flag.Parse() // parse the command line args
}

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *portPtr)
	if *debugPtr {
		log.Printf("Connection String:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	sql := fmt.Sprintf("SELECT TOP %d SESSION_NAME, CLIENT_HOST_NAME, APP_NAME, COMPRESSION_ENABLED, CREATE_TIME, INSERT_SEQ from dbo.SESSION_HISTORY WHERE INSERT_SEQ > %d", *rowsPtr, *seqPtr)
	stmt, err := conn.Prepare(sql)
	if err != nil {
		log.Fatal("Prepare statement failed :", err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		log.Fatal("Query Failed :", err.Error())
	}

	var sessionName string
	var clientHost string
	var appName string
	var compression bool
	var createTime time.Time
	var insertSeq uint32

	for rows.Next() {
		err := rows.Scan(&sessionName, &clientHost, &appName, &compression, &createTime, &insertSeq)
		if err != nil {
			log.Printf("Error getting row : ", err.Error())
		}
		log.Printf("session Name : %s, Create Time: %s", sessionName, createTime.String())
	}

}
