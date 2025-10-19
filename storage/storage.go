package storage

import (
	"errors"
	"studentsystem/validation"
)

// Мапы для хранения данных
var Students = make(map[int]string)       // ID -> Имя
var Subjects = make(map[int]string)       // ID -> Название предмета
var StudentSubjects = make(map[int][]int) // StudentID -> []SubjectID

func AddStudent(id int, name string) error {
	// TODO: Добавить студента в мапу Students
	// Проверить что ID не занят, вернуть ошибку если занят
	ValidStudErr := validation.ValidateStudentID(id)
	ValidStudNameErr := validation.ValidateStudentName(name)

	_, err := Students[id]

	if err {
		return errors.New("StudentID is already added")
	}

	if ValidStudErr != nil {
		return ValidStudErr
	}

	if ValidStudNameErr != nil {
		return ValidStudNameErr
	}

	Students[id] = name
	return nil
}

func AddSubject(id int, name string) error {
	// TODO: Добавить предмет в мапу Subjects
	// Проверить что ID не занят, вернуть ошибку если занят
	ValidSubErr := validation.ValidateSubjectID(id)
	ValidSubNameErr := validation.ValidateSubjectName(name)

	_, err := Subjects[id]

	if err {
		return errors.New("SubjectID is already added")
	}

	if ValidSubErr != nil {
		return ValidSubErr
	}

	if ValidSubNameErr != nil {
		return ValidSubNameErr
	}

	Subjects[id] = name
	return nil
}

func AssignSubjectToStudent(studentID int, subjectID int) error {
	// TODO: Добавить предмет студенту
	// Проверить что студент и предмет существуют
	// Добавить subjectID в слайс StudentSubjects[studentID]

	_, StudErr := Students[studentID]
	_, SubErr := Subjects[studentID]

	if !StudErr {
		return errors.New("Student not found")
	}

	if !SubErr {
		return errors.New("Subject not found")
	}

	StudentSubjects[studentID] = append(StudentSubjects[studentID], subjectID)
	return nil
}

func GetStudentSubjects(studentID int) ([]int, error) {
	// TODO: Вернуть слайс ID предметов студента
	// Вернуть ошибку если студент не найден

	_, err := StudentSubjects[studentID]

	if err {
		return []int{}, errors.New("Student not found")
	}

	return StudentSubjects[studentID], nil
}

func GetAllStudents() map[int]string {
	// TODO: Вернуть мапу всех студентов
	return Students
}
