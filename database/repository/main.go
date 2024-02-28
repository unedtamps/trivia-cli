package repository

import (
	"TugasRPL/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/slog"
)

var Query *Queries
var db *sql.DB

var (
	ADMIN     int64            = 1
	USER      int64            = 2
	OK        int64            = 1
	REVIEW    int64            = 2
	REJECTED  int64            = 3
	EASY      int64            = 1
	MEDIUM    int64            = 2
	HARD      int64            = 3
	CORRECT   int64            = 1
	INCORRECT int64            = 2
	MPDIFF    map[string]int64 = map[string]int64{"Easy": EASY, "Medium": MEDIUM, "Hard": HARD}
	MPDIFFREV map[int64]string = map[int64]string{EASY: "Easy", MEDIUM: "Medium", HARD: "Hard"}
	MPSTATUS  map[int64]string = map[int64]string{OK: "OK", REVIEW: "REVIEW", REJECTED: "REJECTED"}
)

func init() {
	var err error
	db, err = sql.Open("mysql", config.DB_URI)
	if err != nil {
		slog.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		slog.Fatal(err)
	}
	Query = New(db)
}
