package utils

// delete course with id, cid and term
func DeleteCourse(id, cid, term string) error {
	stmt, err := Db.Prepare("delete from CourseCalendar where id = ? and cid = ? and  term = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, cid, term)
	return err
}

func DeleteUser(id string) error {
	stmt, err := Db.Prepare("delete from User where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
func DeleteTeacher(tid string) error {
	stmt, err := Db.Prepare("delete from Teacher where tid = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tid)
	if err != nil {
		return err
	}
	return DeleteUser(tid)
}
func DeleteStudent(id string) error {
	stmt, err := Db.Prepare("delete from Student where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return DeleteUser(id)
}

// ------------------------------------------------
