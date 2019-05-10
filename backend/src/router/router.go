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
	http.HandleFunc("/api/", sayhelloName) //设置访问的路由
	http.HandleFunc("/api/login", login.Login)
	http.HandleFunc("/api/student/profile", student.Profile)
	http.HandleFunc("/api/student/courseCalendar", student.CourseCalendar)
	http.HandleFunc("/api/student/courseQuery", student.CourseQuery)
	http.HandleFunc("/api/student/scoreTable", student.ScoreTable)
	http.HandleFunc("/api/student/scoreSummary", student.ScoreSummary)
	http.HandleFunc("/api/student/scoreTrend", student.ScoreTrend)
	http.HandleFunc("/api/student/scoreCollege", student.ScoreCollege)

	http.HandleFunc("/api/teacher/profile", teacher.Profile)
	http.HandleFunc("/api/teacher/scoreManagement", teacher.ScoreManagement)
	http.HandleFunc("/api/teacher/scoreAnalysis", teacher.ScoreAnalysis)
	http.HandleFunc("/api/teacher/courseCalendar", teacher.CourseCalendar)
	http.HandleFunc("/api/teacher/classTable", teacher.ClassTable)

	http.HandleFunc("/api/admin/profile", admin.Profile)
	http.HandleFunc("/api/admin/teacherManagement", admin.TeacherManagement)
	http.HandleFunc("/api/admin/studenManagement", admin.StudentManagement)
}
