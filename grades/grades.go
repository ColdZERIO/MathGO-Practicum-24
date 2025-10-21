package grades

import (
	"errors"
)

// Мапа для хранения оценок: StudentID -> SubjectID -> []int (слайс оценок)
var Grades = make(map[int]map[int][]int)

func AddGrade(studentID int, subjectID int, grade int) error {
	// TODO: Добавить оценку студенту по предмету
	// Проверить что студент и предмет существуют
	// Проверить что оценка от 1 до 10
	// Инициализировать вложенные мапы если нужно

	if _, ok := Grades[studentID]; !ok {
		return errors.New("Student not found")
	}

	Grades[studentID][subjectID] = append(Grades[studentID][subjectID], grade)

	return nil
}

func GetStudentGrades(studentID int) (map[int][]int, error) {
	// TODO: Вернуть мапу предмет -> оценки для студента
	// Вернуть ошибку если студент не найден

	if _, ok := Grades[studentID]; !ok {
		return nil, errors.New("Student not found")
	}

	return Grades[studentID], nil
}

func GetSubjectGrades(subjectID int) (map[int][]int, error) {
	// TODO: Вернуть мапу студент -> оценки по предмету
	// Вернуть ошибку если предмет не найден
	subjectStudentsGrades := make(map[int][]int)
	for studentID := range Grades {
		if _, ok := Grades[studentID]; !ok {
			return nil, errors.New("Student not found")
		}

		if _, ok := Grades[studentID][subjectID]; !ok {
			return nil, errors.New("Subject not found")
		}

		subjectStudentsGrades[studentID] = append(subjectStudentsGrades[studentID], Grades[studentID][subjectID]...)
	}
	return subjectStudentsGrades, nil
}

func UpdateGrade(studentID int, subjectID int, oldGrade int, newGrade int) error {
	// TODO: Обновить конкретную оценку
	// Найти oldGrade в слайсе и заменить на newGrade
	if _, ok := Grades[studentID]; !ok {
		return errors.New("Student not found")
	}

	if _, ok := Grades[studentID][subjectID]; !ok {
		return errors.New("Subject not found")
	}

	for i, grade := range Grades[studentID][subjectID] {
		if oldGrade == grade {
			Grades[studentID][subjectID][i] = newGrade
		}
	}

	return nil
}

func GetStudentSubjectGrades(studentID int, subjectID int) ([]int, error) {
	// TODO: Вернуть все оценки студента по предмету
	if _, ok := Grades[studentID]; !ok {
		return nil, errors.New("Student not found")
	}

	if _, ok := Grades[studentID][subjectID]; !ok {
		return nil, errors.New("Subject not found")
	}

	return Grades[studentID][subjectID], nil
}
