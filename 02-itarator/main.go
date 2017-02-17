package main

import (
	"fmt"

	"github.com/gaku3601/study-golang-design-pattern/02-itarator/student"
)

func main() {
	//生徒名簿
	var studentList student.StudentList

	//5人の生徒名簿を作成する
	studentList = append(studentList, student.NewStudent("gaku1", "男", 20))
	studentList = append(studentList, student.NewStudent("gaku2", "女", 20))
	studentList = append(studentList, student.NewStudent("gaku3", "男", 20))
	studentList = append(studentList, student.NewStudent("gaku4", "男", 20))
	studentList = append(studentList, student.NewStudent("gaku5", "男", 20))

	//名簿内容を出力する。
	for _, v := range studentList {
		fmt.Printf("私の名前は%sです。%sです。偏差値は%dです。\n", v.Name, v.Sex, v.DeviationValue)
	}
}
