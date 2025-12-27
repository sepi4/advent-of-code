package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Derek"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, v := range intArr {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return strArr
}

type Date struct {
	day   int
	month int
	year  int
}
type Kissa struct {
	nimi string
}

// setters
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("incorrect day value")
	}
	d.day = day
	return nil
}
func (d *Date) SetMonth(m int) error {
	if m < 1 || m > 12 {
		return errors.New("incorrect month value")
	}
	d.month = m
	return nil
}
func (d *Date) SetYear(y int) error {
	if y < 1875 || y > time.Now().Year() {
		return errors.New("incorrect year value")
	}
	d.year = y
	return nil
}

// getters
func (d *Date) Day() int {
	return d.day
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Year() int {
	return d.year
}
