package utils

import "errors"

func setTerm(term string, isCurrent string) error {
	stmt, err := Db.Prepare("update Term set isCurrent = ? where term = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(isCurrent, term)
	return err
}
func SetCurrentTerm(term string) error {
	_, err := GetCurrentTerm()
	if err == nil {
		return errors.New("Already have current term.")
	}
	return setTerm(term, "yes")
}
func EndTerm(term string) error {
	ok, err := DoesAllTeaHaveMarked(term)
	if err != nil {
		return err
	}
	if ok == true {
		return errors.New("This term still has teachers who don't finish marking.")
	}
	return setTerm(term, "no")
}

func DeleteTerm(term string) error {
	ok, err := DoesTermHaveCourses(term)
	if err != nil {
		return err
	}
	if ok == true {
		return errors.New("This term still has courses.")
	}
	stmt, err := Db.Prepare("delete from Term where term = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(term)
	return err
}
func DoesTermHaveCourses(term string) (bool, error) {
	var cnt int
	err := Db.QueryRow("select count(cid) from CourseSchedule where term = ?", term).Scan(&cnt)
	return cnt != 0, err
}

func DoesAllTeaHaveMarked(term string) (bool, error) {
	score := -1.0
	var cnt int
	err := Db.QueryRow("select count(id) from CourseCalendar where term = ? and score = ?", term, score).Scan(&cnt)
	return cnt != 0, err
}
func GetAllTerms() (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select term, isCurrent from Term order by id")
	if err != nil {
		return nil, err
	}
	var term, isCurrent string
	for rows.Next() {
		err := rows.Scan(&term, &isCurrent)
		if err != nil {
			return nil, err
		}
		ret["term"] = append(ret["term"], term)
		ret["isCurrent"] = append(ret["isCurrent"], isCurrent)
	}
	return ret, nil
}
func GetCurrentTerm() (string, error) {
	var term string
	tmp := "yes"
	err := Db.QueryRow("select term from Term where isCurrent = ?", tmp).Scan(&term)
	return term, err
}
