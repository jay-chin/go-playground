package main

import (
	"flag"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
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

type SessionHistory struct {
	CompressionEnabled bool      `db:"COMPRESSION_ENABLED"`
	SessionName        string    `db:"SESSION_NAME"`
	ClientHostName     string    `db:"CLIENT_HOST_NAME"`
	AppName            string    `db:"APP_NAME"`
	CreateTime         time.Time `db:"CREATE_TIME"`
	InsertSeq          uint32    `db:"INSERT_SEQ"`
}

func init() {
	flag.Parse() // parse the command line args
}

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *portPtr)
	if *debugPtr {
		log.Printf("Connection String:%s\n", connString)
	}
	db, err := sqlx.Connect("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	err = db.Ping()
	defer db.Close()

	sql := fmt.Sprintf("SELECT TOP %d SESSION_NAME, CLIENT_HOST_NAME, APP_NAME, CREATE_TIME, INSERT_SEQ, COMPRESSION_ENABLED from dbo.SESSION_HISTORY WHERE INSERT_SEQ > %d", *rowsPtr, *seqPtr)
	sh := []SessionHistory{}
	err = db.Select(&sh, sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, element := range sh {
		fmt.Printf("%v\n", element)
	}
	fmt.Printf("Number of records returned : %d", len(sh))

}
