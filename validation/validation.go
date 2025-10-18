package validation

// kek lol
import (
	"errors"
	"regexp"
)

func StringChecker(str string) bool {
	re := regexp.MustCompile(`^[А-Яа-я ?]+$`)
	return re.MatchString(str)
}

func ValidateStudentID(id int) error {
	if id < 0 && id > 10000 {
		return errors.New("invalid student ID: must be between 1 and 10000")
	}

	return nil
}

func ValidateStudentName(name string) error {
	if len(name) >= 2 && len(name) <= 50 && StringChecker(name) {
		return nil
	}
	return errors.New("invalid student name")
}

func ValidateSubjectID(id int) error {
	if id < 0 && id > 1000 {
		return errors.New("invalid subject ID")
	}
	return nil
}

func ValidateSubjectName(name string) error {
	if len(name) >= 2 && len(name) <= 30 && StringChecker(name) {
		return nil
	}
	return errors.New("invalid student name")
}

func ValidateGrade(grade int) error {
	if grade >= 1 && grade <= 10 {
		return nil
	}
	return errors.New("invalid grade: must be between 1 and 10")
}
