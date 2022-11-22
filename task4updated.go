package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
	subject    []string
}

func NewStudent(rollno int, name string, address string, subject []string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.subject = subject
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, subject []string) *Student {
	st := NewStudent(rollno, name, address, subject)
	ls.list = append(ls.list, st)
	return st
}

func (ls *StudentList) Print() {
	for i := range ls.list {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student Rollno: %d\n", ls.list[i].rollnumber)
		fmt.Printf("Student Name: %s\n", ls.list[i].name)
		fmt.Printf("Student Address: %s\n\n", ls.list[i].address)
		fmt.Printf("Student Subject: %s\n\n", ls.list[i].subject)
		var str1 string
		str1 += strconv.Itoa(ls.list[i].rollnumber) + ls.list[i].name + ls.list[i].address
		sum := sha256.Sum256([]byte(str1))
		fmt.Printf("%x\n", sum) //hexadecimal

	}
}

func main() {
	student := new(StudentList)
	var arr = []string{"Blockchain"}
	var arr2 = []string{"English"}
	student.CreateStudent(2034, "Talha", "Pasha House", arr)
	student.CreateStudent(1111, "Ahmed", "Judicial Colony, Rawalpindi", arr2)
	student.Print()
}
