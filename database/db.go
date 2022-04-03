package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

const (
	dbFile = "./database.db"
)

type PillDay struct {
	Id         int
	Date       time.Time
	IsFirstDay bool
}

type pillDayDto struct {
	Id         int
	Date       string
	IsFirstDay bool
}

func (pd *pillDayDto) Model() *PillDay {
	dt, _ := time.Parse(time.RFC3339Nano, pd.Date)
	return &PillDay{
		Id:         pd.Id,
		Date:       dt,
		IsFirstDay: pd.IsFirstDay,
	}
}

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", dbFile) // Open the created SQLite File
	if err != nil {
		log.Fatalf("Cannot create db with path %s", dbFile)
	}
	return db
}

func Dispose(db *sql.DB) {
	db.Close()
}

func insertPillDay(db *sql.DB, dt time.Time) error {
	query, _ := db.Prepare("INSERT INTO PillDay (Date, IsFirstDay) VALUES (?,?)")
	_date := dt.UTC().Format(time.RFC3339Nano)
	_isFirstDate := IsFirstDay(db, dt)
	_, err := query.Exec(_date, _isFirstDate)
	return err
}

func Take(db *sql.DB) error {
	dt := time.Now().Truncate(time.Second)
	return insertPillDay(db, dt)
}

func GetLastPill(db *sql.DB) *PillDay {
	var pillDay pillDayDto
	err := db.QueryRow("SELECT * FROM PillDay ORDER BY Date DESC").Scan(&pillDay.Id, &pillDay.Date, &pillDay.IsFirstDay)
	if err != nil {
		return nil
	}
	return pillDay.Model()
}

func GetFirstPillDay(db *sql.DB) time.Time {
	var dtString string
	err := db.QueryRow("SELECT Date FROM PillDay WHERE IsFirstDay = 1 ORDER BY Date DESC").Scan(&dtString)
	if err != nil {
		panic(fmt.Errorf("cannot find first pill day"))
	}
	dt, _ := time.Parse(time.RFC3339Nano, dtString)
	return dt
}

func GetLastBox(db *sql.DB) []PillDay {
	pillDays := make([]PillDay, 0, 28)

	firstPillDay := GetFirstPillDay(db).Format(time.RFC3339Nano)
	rows, err := db.Query("SELECT * FROM PillDay WHERE Date >= ? ORDER BY Date DESC LIMIT 28", firstPillDay)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var pillDay pillDayDto
	for rows.Next() {
		err := rows.Scan(&pillDay.Id, &pillDay.Date, &pillDay.IsFirstDay)
		if err != nil {
			continue
		}
		pillDays = append(pillDays, *pillDay.Model())
	}
	return pillDays
}
