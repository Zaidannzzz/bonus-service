package helpers

import (
	"fmt"
	"strings"
	"time"
)

type DateConverter struct {
	UnixTimestamp int64
	Time          time.Time
	Location      time.Location
}

type IDateConverter interface {
	New(UnixTimestamp int)
	ConvertTimeToBeginningOfYear() error
	ConvertTimeToBeginningOfDay() error
	ConvertTimeToBeginningOfMonth()
	ConvertTimeToEndOfMonth()
}

func (d *DateConverter) New(UnixTimestamp int64) {
	d.UnixTimestamp = UnixTimestamp
	d.Time = ConvertUnixToDate(int(d.UnixTimestamp))
	loc, _ := time.LoadLocation("Asia/Jakarta")
	d.Location = *loc
}

func (d *DateConverter) ConvertTimeToBeginningOfYear() error {
	year, _, _ := d.Time.Date()
	parsingString := fmt.Sprintf("%v", year)
	result, err := time.ParseInLocation("2006", parsingString, &d.Location)
	if err != nil {
		return err
	} else {
		d.Time = result
		return nil
	}
}

func (d *DateConverter) ConvertTimeToBeginningOfDay() error {
	year, month, day := d.Time.Date()
	parsingString := fmt.Sprintf("%v-%v-%v", year, month, day)
	fmt.Printf("parsingString : %v \n", parsingString)
	result, err := time.ParseInLocation("2006-January-2", parsingString, &d.Location)
	if err != nil {
		return err
	} else {
		d.Time = result
		return nil
	}
}
func (d *DateConverter) ConvertTimeToBeginningOfMonth() {
	year, month, day := d.Time.Date()
	d.Time = time.Date(year, month, day+(-d.Time.Day()+1), 0, 0, 0, 0, &d.Location)
}

func (d *DateConverter) ConvertTimeToEndOfMonth() {
	year, month, day := d.Time.Date()
	d.Time = time.Date(year, month+time.Month(1), day+(-d.Time.Day()), 0, 0, 0, 0, &d.Location)
}

/*
   |--------------------------------------------------------------------------
   | Validate date time and convert string to time.Time
   |--------------------------------------------------------------------------
   |
   | This function convert string to time.Time,
   | Before convert to time.Time string with "/" must replace with "-".
   | This function return time.Time type end error for time validation.
*/

func ValidateDate(date string) (*time.Time, error) {

	dateFormat := strings.Replace(date, "/", "-", -1)

	result, errParseTime := time.Parse("2006-01-02", dateFormat)
	if errParseTime != nil {
		return nil, errParseTime
	}
	return &result, nil

}

/*
   |--------------------------------------------------------------------------
   | Convert Unix Timestamp to time.Time
   |--------------------------------------------------------------------------
   |
   | This function convert UnixTimestamp to time.Time
*/

func ConvertUnixToDate(date int) time.Time {
	//make timezone WIB
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Err: ", err)
	}
	result := time.Unix(int64(date), 0).In(loc)

	return result
}

/*
   |--------------------------------------------------------------------------
   | Validate unixtamestamp
   |--------------------------------------------------------------------------
   |
   | This function Make Custom Binding Validation For Enum data Type
*/
// func Unix(sec int64, nsec int64) (*time.Time, error)

func ConvertUnix(date string) (*time.Time, error) {

	resultDate, err := ValidateDate(date)
	if err != nil {
		return nil, err
	}
	// if your unix timestamp is not in int64 format

	// int64 to time.Time
	myTime := time.Unix(resultDate.Unix(), 0)

	return &myTime, nil
}

/*
   |--------------------------------------------------------------------------
   | Convert Unix time.Time to Timestamp
   |--------------------------------------------------------------------------
   |
   | This function convert UnixTimestamp to time.Time
*/

func ConvertDateToUnix(date time.Time) int64 {

	result := date.Unix()

	return result

}
