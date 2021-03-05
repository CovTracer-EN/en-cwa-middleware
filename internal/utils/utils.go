package utils

import (
	"database/sql"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"math"
	"os"
	"time"
)

func HttpPost(url string, body interface{}, headers map[string]string) (res *resty.Response, err error) {
	client := resty.New()

	req := client.R()
	if len(headers) == 0 {
		req.SetHeader("Content-Type", "application/json")
	} else {
		for key, value := range headers {
			req.SetHeader(key, value)
		}
	}

	req.SetBody(body)
	return req.Post(url)
}

func HttpGet(url string, headers map[string]string) (res *resty.Response, err error) {
	client := resty.New()

	req := client.R()
	if len(headers) == 0 {
		req.SetHeader("Content-Type", "application/json")
	} else {
		for key, value := range headers {
			req.SetHeader(key, value)
		}
	}

	return req.Get(url)
}

func CalculateDSOSVector(otp string) (dsosVector [15]int32) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DatabaseHost"),
		os.Getenv("DatabasePort"),
		os.Getenv("DatabaseUser"),
		os.Getenv("DatabasePass"),
		os.Getenv("DatabaseName"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Print(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Print(err)
		return
	}

	var _symptoms int
	var _testDate sql.NullTime
	var _symptomDate sql.NullTime

	err = db.
		QueryRow("SELECT \"symptomDate\", \"symptoms\", \"testDate\" FROM access_codes WHERE code = $1", otp).
		Scan(&_symptomDate, &_symptoms, &_testDate)
	if err != nil {
		log.Println(err)
		return
	}

	var dsos = int32(4000)
	//Are there any symptoms?
	if _symptoms == 0 { //yes
		//Is the date of the onset of symptoms known?
		if _symptomDate.Valid { //yes, specific date is known
			dsos = int32(math.Ceil(time.Now().Sub(_symptomDate.Time).Hours() / 24))
		} else {
			if _testDate.Valid {
				dsos = int32(math.Ceil(time.Now().Sub(_testDate.Time).Hours() / 24) + 2000)
			}
		}
	} else if _symptoms == 1 { //unknown
		dsos = 4000
	} else if _symptoms == 2 { //no
		if _testDate.Valid {
			dsos = int32(math.Ceil(time.Now().Sub(_testDate.Time).Hours() / 24) + 3000)
		}
	}

	for i := 0; i < len(dsosVector); i++ {
		dsosVector[i] = dsos
		dsos -= 1
	}
	return dsosVector
}