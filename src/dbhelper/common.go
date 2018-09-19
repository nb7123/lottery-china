package dbhelper

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"model"
	"errors"
)

var conn *sql.DB

const createTableSSC= `
	create table if not exists  ssc (
		id integer not null primary key,
		value01 integer,
		value02 integer,
		value03 integer,
		value04 integer,
		value05 integer);`

const createTableWelfare = `
	create table if not exists  welfare (
		id integer not null primary key,
		red01 integer,
		red02 integer,
		red03 integer,
		red04 integer,
		red05 integer,
		red06 integer,
		blue  integer);`

const selectLastInsertedWelfare = `select id from welfare order by id DESC LIMIT 1`

func InitDB() error {
	var err error

	conn, err = sql.Open("sqlite3", "../../Lottery.db")

	if nil == err {
		_, err = conn.Exec(createTableSSC)
	}

	if nil == err {
		_, err = conn.Exec(createTableWelfare)
	}

	return err
}

func CloseDB() {
	conn.Close()
}

func InsertWelfare(lotteries []model.Lottery) (error) {
	var err error = nil

	if len(lotteries) < 7 {
		err = errors.New("error: Invalid welfare number")
	}

	var tx *sql.Tx

	if nil == err {
		tx, err = conn.Begin()
	}

	if nil == err {
		for _, item := range lotteries {
			_, err = tx.Exec(`replace into welfare(id,
										red01,
										red02,
										red03,
										red04,
										red05,
										red06,
										blue) values (?, ?, ?, ?, ?, ?, ?, ?)`,
				item.Num,
				item.ResultNum[0],
				item.ResultNum[1],
				item.ResultNum[2],
				item.ResultNum[3],
				item.ResultNum[4],
				item.ResultNum[5],
				item.ResultNum[6])

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

func SearchLastInsertedWelfare() (int, error) {
	row := conn.QueryRow(selectLastInsertedWelfare)

	var lastId = 0

	err := row.Scan(&lastId)

	if err == sql.ErrNoRows {
		err = nil
	}

	return lastId, err
}

func SearchWelfareList(startId int, endId int) ([]model.Lottery, error) {
	var lottery []model.Lottery

	rows, err := conn.Query("select id, red01, red02, red03, red04, red05, red06, blue from welfare")

	if nil == err {
		var id, red01, red02, red03, red04, red05, red06, blue int
		for rows.Next() {
			err = rows.Scan(&id, &red01, &red02, &red03, &red04, &red05, &red06, &blue)
			if nil != err {
				break
			}

			var numbers = []int{red01, red02, red03, red04, red05, red06, blue}

			lottery = append(lottery, model.Lottery{Num: id, ResultNum: numbers})
		}
	}

	return lottery, err
}

func SearchNumberCount1(number int, startId int) ([]model.Lottery, error) {
	var lottery []model.Lottery

	rows, err := conn.Query("select count(*) from" +
		" (select * from welfare where id >= ? order by id desc)" +
		" where (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?)", startId,
		number, number, number, number, number, number)

	if nil == err {
		var id, red01, red02, red03, red04, red05, red06, blue int
		for rows.Next() {
			err = rows.Scan(&id, &red01, &red02, &red03, &red04, &red05, &red06, &blue)
			if nil != err {
				break
			}

			var numbers = []int{red01, red02, red03, red04, red05, red06, blue}

			lottery = append(lottery, model.Lottery{Num: id, ResultNum: numbers})
		}
	}

	return lottery, err
}

func SearchNumberCount2(number1 int, number2 int, startId int) ([]model.Lottery, error) {
	var lottery []model.Lottery

	rows, err := conn.Query("select id, red01, red02, red03, red04, red05, red06, blue from" +
		" (select * from welfare where id >= ? order by id desc)" +
		" where (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?) AND" +
		" (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?)", startId,
		number1, number1, number1, number1, number1, number1,
		number2, number2, number2, number2, number2, number2)

	if nil == err {
		var id, red01, red02, red03, red04, red05, red06, blue int
		for rows.Next() {
			err = rows.Scan(&id, &red01, &red02, &red03, &red04, &red05, &red06, &blue)
			if nil != err {
				break
			}

			var numbers = []int{red01, red02, red03, red04, red05, red06, blue}

			lottery = append(lottery, model.Lottery{Num: id, ResultNum: numbers})
		}
	}

	return lottery, err
}

func SearchNumberCount3(number1 int, number2 int, number3 int, startId int) ([]model.Lottery, error) {
	var lottery []model.Lottery

	rows, err := conn.Query("select id, red01, red02, red03, red04, red05, red06, blue from" +
		" (select * from welfare where id >= ? order by id desc)" +
		" where (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?) AND" +
		" (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?)  AND" +
		" (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?)", startId,
		number1, number1, number1, number1, number1, number1,
		number2, number2, number2, number2, number2, number2,
		number3, number3, number3, number3, number3, number3)

	if nil == err {
		var id, red01, red02, red03, red04, red05, red06, blue int
		for rows.Next() {
			err = rows.Scan(&id, &red01, &red02, &red03, &red04, &red05, &red06, &blue)
			if nil != err {
				break
			}

			var numbers = []int{red01, red02, red03, red04, red05, red06, blue}

			lottery = append(lottery, model.Lottery{Num: id, ResultNum: numbers})
		}
	}

	return lottery, err
}

func SearchNumberCount4(number1 int, number2 int, number3 int, number4 int, startId int) ([]model.Lottery, error) {
	var lottery []model.Lottery

	rows, err := conn.Query("select id, red01, red02, red03, red04, red05, red06, blue from" +
		" (select * from welfare where id >= ? order by id desc)" +
		" where (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?) AND" +
		" (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?) AND" +
		" (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?) AND" +
		" (red01 = ?" +
		" or red02 = ?" +
		" or red03 = ?" +
		" or red04 = ?" +
		" or red05 = ?" +
		" or red06 = ?)", startId,
		number1, number1, number1, number1, number1, number1,
		number2, number2, number2, number2, number2, number2,
		number3, number3, number3, number3, number3, number3,
		number4, number4, number4, number4, number4, number4)

	if nil == err {
		var id, red01, red02, red03, red04, red05, red06, blue int
		for rows.Next() {
			err = rows.Scan(&id, &red01, &red02, &red03, &red04, &red05, &red06, &blue)
			if nil != err {
				break
			}

			var numbers = []int{red01, red02, red03, red04, red05, red06, blue}

			lottery = append(lottery, model.Lottery{Num: id, ResultNum: numbers})
		}
	}

	return lottery, err
}
