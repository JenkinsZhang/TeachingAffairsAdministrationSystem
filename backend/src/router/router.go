package router

import (
	"net/http"
	"taas/func/admin"
	"taas/func/login"
	"taas/func/student"
	"taas/func/teacher"
	"taas/func/tools"
)

func Build() {
	//http.HandleFunc("/api/", sayhelloName) //设置访问的路由
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
	http.HandleFunc("/api/admin/studentManagement", admin.StudentManagement)

	http.HandleFunc("/api/getDepartment", tools.GetDepartment)
	http.HandleFunc("/api/getCurrentTerm", tools.GetCurrentTerm)
}
