package dbhelper

import (
	"testing"
	"model"
	"encoding/json"
)

func TestConvertCurrent(t *testing.T) {
	InitDB()
	defer CloseDB()

	var lottery []model.Lottery

	rows, err := conn.Query("select id, result from lottery")

	if nil == err {
		var id int
		var result []byte
		for rows.Next() {
			err = rows.Scan(&id, &result)
			if nil != err {
				break
			}

			var numbers []int
			err = json.Unmarshal(result, &numbers)

			if nil != err {
				break
			}

			lottery = append(lottery, model.Lottery{Num: id, ResultNum: numbers})
		}
	}

	if nil == err {
		err = InsertWelfare(lottery)
	}

	if nil != err {
		t.Fatal(err)
	}
}

func TestSearchNumberCount(t *testing.T) {
	InitDB()
	defer CloseDB()

	count, err := SearchNumberCount(2, 7)

	if nil != err {
		t.Fatal(err)
	}

	if 1 != count {
		t.Log(count)
		t.Fail()
	}
}
