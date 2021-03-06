package utils

import (
	"errors"
	"taas/models"
)

func InsertCourseSchedule(info map[string]string) error {
	err := CheckTidInTeacher(info["Tid"])
	if err != nil {
		return errors.New("invalid tid")
	}
	stmt, err := Db.Prepare("insert into CourseSchedule(tid,cid,term,classTime) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(info["Tid"], info["Cid"], info["Term"], info["ClassTime"])
	return err
}
func InsertCourse(info map[string]string) error {
	did, err := QueryDepartmentId(info["Dname"])
	if err != nil {
		return errors.New("invalid dname")
	}
	err = CheckCidInCourse(info["Cid"])
	if err == nil {
		return errors.New("exists cid")
	}
	err = CheckCnameInCourse(info["Cname"])
	if err == nil {
		return errors.New("exists cname")
	}

	stmt, err := Db.Prepare("insert into Course(cid,cname,credit,did) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(info["Cid"], info["Cname"], info["Credit"], did)
	return err
}

func InsertTerm(term string) error {
	stmt, err := Db.Prepare("insert into Term(term,isCurrent,id) values(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(term, "no", models.TermId)
	models.TermId++
	return err
}
func InsertCourseCalendar(info map[string]string) error {
	// isCourseCollision()
	stmt, err := Db.Prepare("insert into CourseCalendar(id,cid,tid,term,classTime) values(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(info["Id"], info["Cid"], info["Tid"], info["Term"], info["ClassTime"])
	return err
}

func InsertStudent(info map[string]string) error {
	err := InsertUser(info["Id"], "student", info["Id"])
	if err != nil {
		return err
	}
	stmt, err := Db.Prepare("insert into Student(id,did,name,gender,birthday,birthplace,phone) values(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(info["Id"], info["Did"], info["Name"], info["Gender"], info["Birthday"], info["Birthplace"], info["Phone"])
	if err != nil {
		return err
	}
	return err
}

func InsertUser(id, identity, password string) error {
	stmt, err := Db.Prepare("insert into User(id,password,identity) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, password, identity)
	return err
}

func InsertTeacher(info map[string]string) error {
	err := InsertUser(info["Tid"], "teacher", info["Tid"])
	if err != nil {
		return err
	}
	stmt, err := Db.Prepare("insert into Teacher(tid,did,tname,gender,birthday,education,wage) values(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(info["Tid"], info["Did"], info["Tname"], info["Gender"], info["Birthday"], info["Education"], info["Wage"])
	if err != nil {
		return err
	}
	return err
}

// ----------------------------------------------------
