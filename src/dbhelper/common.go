package dbhelper

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"model"
	"encoding/json"
)

var conn *sql.DB

const createStatement = `
	create table if not exists  lottery (
		id integer not null primary key,
		opening_time datetime,
		result json,
	 	d_type integer);`

const selectLastInsertedLottery = `select id from lottery where d_type = ? order by id DESC LIMIT 1`

func InitDB() error {
	var err error

	conn, err = sql.Open("sqlite3", "./Lottery.db")

	if nil == err {
		_, err = conn.Exec(createStatement)
	}

	return err
}

func InsertLottery(lotteries []model.Lottery) (error) {

	tx, err := conn.Begin()
	if nil == err {
		for _, item := range lotteries {
			var result []byte
			result, err = json.Marshal(item.ResultNum)
			
			_, err = tx.Exec(`replace into lottery(id,
 										opening_time,
   										result,
    									d_type) values (?, ?, ?, ?)`,
				item.Num, item.OpeningTime, result, item.Type)

			if nil != err {
				break
			}
		}
	}

	if nil == err {
		tx.Commit()
	} else if nil != tx {
		tx.Rollback()
	}
	if nil == err {
		if tx != nil {
			tx.Commit()
		}

	} else {
		tx.Rollback()
	}

	return err
}

func SearchLastInertedLottery(dType int) (int, error) {
	row := conn.QueryRow(selectLastInsertedLottery, dType)

	var lastId = 0

	err := row.Scan(&lastId)

	if err == sql.ErrNoRows {
		err = nil
	}

	return lastId, err
}
