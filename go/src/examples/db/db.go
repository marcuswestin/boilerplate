package db

import (
	_ "github.com/lib/pq"
	sql "github.com/marcuswestin/go-x-sql"
)

var (
	driver           = "postgres"
	dataSourceString = "postgres://root@localhost/ostraatest?sslmode=disable&port=5432"
	Db               sql.Db
)

type Thing struct {
	ThingId int64
	Name    string
}

func Setup() {
	if Db != nil {
		return
	}
	Db = sql.MustConnect(driver, dataSourceString, sql.DbNameConvention_under_score)
}

func Select(dest interface{}, query string, args ...interface{}) {
	sql.Must(sql.CheckDest(dest))
	sql.Must(Db.Select(dest, query, args...))
}
func SelectOne(dest interface{}, query string, args ...interface{}) {
	sql.Must(sql.CheckDest(dest))
	sql.Must(Db.SelectOne(dest, query, args...))
}
func SelectOneMaybe(dest interface{}, query string, args ...interface{}) bool {
	sql.Must(sql.CheckDest(dest))
	return sql.MustBool(Db.SelectOneMaybe(dest, query, args...))
}
func InsertAndGetId(query string, args ...interface{}) (id int64) {
	return sql.MustInt(Db.InsertAndGetId(query, args...))
}
func InsertIgnoreId(query string, args ...interface{}) {
	sql.Must(Db.InsertIgnoreId(query, args...))
}
func InsertIgnoreDuplicate(query string, args ...interface{}) bool {
	return sql.MustBool(Db.InsertIgnoreDuplicate(query, args...))
}
func Update(query string, args ...interface{}) (rowsAffected int64) {
	return sql.MustInt(Db.Update(query, args...))
}
func UpdateOne(query string, args ...interface{}) {
	sql.Must(Db.UpdateOne(query, args...))
}
func UpdateNum(expectedRowsAffected int64, query string, args ...interface{}) {
	sql.Must(Db.UpdateNum(expectedRowsAffected, query, args...))
}
func Exec(query string, args ...interface{}) {
	sql.Must(Db.Exec(query, args...))
}

type TxFn func(tx Tx) error
type Tx sql.Tx

func Transact(txFn TxFn) {
	sql.Must(Db.Transact(func(tx sql.Tx) error {
		return txFn(tx)
	}))
}
