package router

import (
	"fmt"
	"net/http"
	"strings"
	"taas/func/admin"
	"taas/func/login"
	"taas/func/student"
	"taas/func/teacher"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func Build() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/student/profile", student.Profile)
	http.HandleFunc("/student/courseCalendar", student.CourseCalendar)
	http.HandleFunc("/student/courseQuery", student.CourseQuery)
	http.HandleFunc("/student/scoreTable", student.ScoreTable)
	http.HandleFunc("/student/scoreSummary", student.ScoreSummary)
	http.HandleFunc("/student/scoreTrend", student.ScoreTrend)
	http.HandleFunc("/student/scoreCollege", student.ScoreCollege)

	http.HandleFunc("/teacher/profile", teacher.Profile)
	http.HandleFunc("/teacher/scoreManagement", teacher.ScoreManagement)
	http.HandleFunc("/teacher/scoreAnalysis", teacher.ScoreAnalysis)
	http.HandleFunc("/teacher/courseCalendar", teacher.CourseCalendar)
	http.HandleFunc("/teacher/classTable", teacher.ClassTable)

	http.HandleFunc("/admin/profile", admin.Profile)
	http.HandleFunc("/admin/teacherManagement", admin.TeacherManagement)
	http.HandleFunc("/admin/studenManagement", admin.StudentManagement)
}
