package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vharitonsky/iniflags"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Conditions map[string]string

var db *sql.DB
var err error

func query(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	var table string
	var out string
	conditions := Conditions{}
	for k, v := range r.Form {
		if k == "table" {
			table = v[0]
		} else {
			conditions[k] = v[0]
		}
	}

	if table == "" {
		out = "[]"
	} else {
		out = getRows(table, conditions)
	}

	fmt.Fprintln(w, out)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getRows(table string, conditions Conditions) string {
	var SQL string
	if conditions != nil && len(conditions) > 0 {
		var temp []string
		for k, v := range conditions {
			temp = append(temp, fmt.Sprintf("`%v`='%v'", k, v))
		}
		SQL = fmt.Sprintf("SELECT * FROM `%v` where %v", table, strings.Join(temp, " and "))
	} else {
		SQL = fmt.Sprintf("SELECT * FROM `%v`", table)
	}
	fmt.Println(SQL)
	rows, err := db.Query(SQL)
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	out := make([]map[string]string, 0)
	for rows.Next() {
		record := make(map[string]string)
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		out = append(out, record)
	}

	data, err := json.MarshalIndent(out, " ", "   ")
	checkErr(err)
	return string(data)
}

func signalHandler() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-sc
	db.Close()
	os.Exit(0)
}

func main() {
	listeningPort := flag.String("port", "8082", "listening port")

	mysqlhost := flag.String("mysqlHost", "", "mysql host")
	mysqlport := flag.Int("mysqlPort", 3306, "mysql port")
	username := flag.String("username", "", "mysql username")
	password := flag.String("password", "", "mysql password")
	database := flag.String("database", "", "mysql database")
	charset := flag.String("charset", "utf8", "mysql charset")

	iniflags.Parse()

	connectString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v", *username, *password, *mysqlhost, *mysqlport, *database, *charset)
	db, err = sql.Open("mysql", connectString)
	checkErr(err)
	err = db.Ping()
	checkErr(err)

	//启动web服务
	http.HandleFunc("/query", query)
	http.NotFoundHandler()
	fmt.Println("listening port: ", *listeningPort)

	go signalHandler()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *listeningPort), nil))
}
