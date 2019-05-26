package utils

import "errors"

func DeleteCourse(cid string) error {
	var cnt int
	err := Db.QueryRow("select count(id) from CourseCalendar where cid = ?", cid).Scan(&cnt)
	if err != nil {
		return err
	}
	if cnt != 0 {
		return errors.New("student have selected this course. so you can't delete it.")
	}
	stmt, err := Db.Prepare("delete from Course where cid = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cid)
	return err
}

func DeleteCourseSchedule(cid, tid, term string) error {
	stmt, err := Db.Prepare("delete from CourseSchedule where cid = ? and tid = ? and  term = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cid, tid, term)
	return err
}

// delete course with id, cid and term
func DeleteCourseCalendar(id, cid, term string) error {
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
	stmt, err := Db.Prepare("delete from CourseSchedule where tid = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tid)
	if err != nil {
		return err
	}
	stmt, err = Db.Prepare("delete from Teacher where tid = ?")
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
