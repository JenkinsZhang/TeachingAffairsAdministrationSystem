package utils

import (
	"log"
)

func InsertCourse(info map[string]string) ( error) {
	// isCourseCollision()
	stmt, err := Db.Prepare("insert into CourseCalendar(id,cid,tid,term,classTime) values(id = ?, cid = ?,tid = ?, term = ?, classTime=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(info["Id"], info["Cid"], info["Tid"], info["Term"], info["ClassTime"])
	if err != nil {
		return err
	}
	return nil
}

// ----------------------------------------------------
func InsertUser(id, identity, password string) {
	stmt, err := Db.Prepare("insert into User(id,password,identity) values(?,?,?)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(id, password, identity)
	if err != nil {
		log.Println(err)
	}
}
func InsertTeacher(info map[string]string) {
	stmt, err := Db.Prepare("insert into Teacher(tid,did,tname,gender,birthday,education,wage) values(?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(info["Tid"], info["Did"], info["Tname"], info["Gender"], info["Birthday"], info["Education"], info["Wage"])
	if err != nil {
		log.Println(err)
	}

	InsertUser(info["Tid"], "teacher", info["Tid"])

}

func InsertStudent(info map[string]string) {
	stmt, err := Db.Prepare("insert into Student(id,did,name,gender,birthday,birthplace,phone,grade) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(info["Id"], info["Did"], info["Name"], info["Gender"], info["Birthday"], info["Birthplace"], info["Phone"], info["Grade"])
	if err != nil {
		log.Println(err)
	}
	stmt, err = Db.Prepare("insert into User(id,password,identity) values(?,?,?)")
	InsertUser(info["Id"], "student", info["Id"])
}
