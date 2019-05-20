package utils

func InsertTerm(term string) error {
	stmt, err := Db.Prepare("insert into Term(term,isCurrent) values(term = ?, isCurrent = ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(term, "no")
	return err
}
func InsertCourse(info map[string]string) error {
	// isCourseCollision()
	stmt, err := Db.Prepare("insert into CourseCalendar(id,cid,tid,term,classTime) values(id = ?, cid = ?,tid = ?, term = ?, classTime=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(info["Id"], info["Cid"], info["Tid"], info["Term"], info["ClassTime"])
	return err
}

func InsertStudent(info map[string]string) error {
	stmt, err := Db.Prepare("insert into Student(id,did,name,gender,birthday,birthplace,phone) values(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(info["Id"], info["Did"], info["Name"], info["Gender"], info["Birthday"], info["Birthplace"], info["Phone"])
	if err != nil {
		return err
	}
	return InsertUser(info["Id"], "student", info["Id"])
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
	stmt, err := Db.Prepare("insert into Teacher(tid,did,tname,gender,birthday,education,wage) values(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(info["Tid"], info["Did"], info["Tname"], info["Gender"], info["Birthday"], info["Education"], info["Wage"])
	if err != nil {
		return err
	}
	return InsertUser(info["Tid"], "teacher", info["Tid"])

}

// ----------------------------------------------------
