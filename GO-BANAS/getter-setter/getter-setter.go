package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

type Student struct {
	fname     string
	lname     string
	bdate     time.Time
	studentId int
}

type StudentInfo struct {
	fname     string
	lname     string
	studentId int
}

func (s *Student) getInfo() string {
	return fmt.Sprintf(
		"first name: \t%s\nlast name: \t%s\nstudent id: \t%d\nbirth date: \t%d.%d.%d\n",
		s.fname,
		s.lname,
		s.studentId,
		s.bdate.Year(),
		s.bdate.Month(),
		s.bdate.Day(),
	)
}

func (s *Student) setBDate(d time.Time) {
	s.bdate = d
}

func main() {
	s := Student{
		fname:     "Ville",
		lname:     "Valtonen",
		bdate:     time.Date(2001, 1, 1, 0, 0, 0, 0, time.Local),
		studentId: 123,
	}
	pl(s.getInfo())
	s.setBDate(s.bdate.AddDate(-5, -5, -1))
	pl(s.getInfo())
}
