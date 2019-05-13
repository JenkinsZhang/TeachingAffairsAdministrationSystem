package utils

import "log"

// delete course with id, cid and term
func DeleteCourse(id, cid, term string) error {
	stmt, err := Db.Prepare("delete from CourseCalendar where id = ? and cid = ? and  term = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, cid, term)
	if err != nil {
		return err
	}
	return nil
}

// ------------------------------------------------
func DeleteUser(id string) {
	stmt, err := Db.Prepare("delete from User where id = ?")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
	}
}
func DeleteTeacher(tid string) {
	stmt, err := Db.Prepare("delete from Teacher where tid = ?")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(tid)
	if err != nil {
		log.Println(err)
	}
	DeleteUser(tid)
}
func DeleteStudent(id string) {
	stmt, err := Db.Prepare("delete from Student where id = ?")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
	}
	DeleteUser(id)
}
}
