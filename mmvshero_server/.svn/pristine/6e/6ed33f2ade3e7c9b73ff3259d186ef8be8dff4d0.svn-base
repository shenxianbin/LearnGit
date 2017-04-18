package mysql

import (
	"database/sql"
	. "galaxy"
	"galaxy/utils"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type SqlCommand struct {
	query string
	args  []interface{}
	cb    func(query string, args ...interface{})
}

var db *sql.DB
var sqlExecChan chan *SqlCommand
var wg sync.WaitGroup

func Init(dataSourceName string, chan_size int32) error {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	sqlExecChan = make(chan *SqlCommand, chan_size)
	go asnyexec()

	return nil
}

func Wait() {
	wg.Wait()
}

func Quit() {
	sqlExecChan <- nil
}

func asnyexec() {
	defer utils.Stack()

	wg.Add(1)
	for cmd := range sqlExecChan {
		if cmd == nil {
			break
		}

		if cmd.cb != nil {
			cmd.cb(cmd.query, cmd.args...)
		}
	}
	wg.Done()
}

func AsynExec(cb func(query string, args ...interface{}), query string, args ...interface{}) {
	sqlCmd := &SqlCommand{
		query: query,
		args:  args,
		cb:    cb,
	}

	sqlExecChan <- sqlCmd
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	LogDebug(query, "args :", args)

	res, err := stmt.Exec(args...)
	return res, err
}
