package statistics

import (
	"errors"
	"studentsystem/grades"
	"studentsystem/storage"
)

var errStudent = errors.New("Студент не найден")
var errSubject = errors.New("Предмет не найден")

func CalculateStudentAverage(studentID int) (float64, error) {
	// TODO: Посчитать средний балл студента по всем предметам
	// Использовать циклы для обхода всех оценок
	var sum int
	var count int

	res, err := grades.Grades[studentID]
	if !err {
		return 0, errStudent
	}

	for _, val := range res {
		for ch := range val {
			sum += ch
			count++
		}
	}
	return float64(sum) / float64(count), nil
}

func CalculateSubjectAverage(subjectID int) (float64, error) {
	// TODO: Посчитать средний балл по предмету

	_, err := storage.Subjects[subjectID]
	if !err {
		return 0, errSubject
	}

	var sum int
	var count int
	for _, val := range grades.Grades {
		for _, ch := range val[subjectID] {
			sum += ch
			count++
		}
	}
	return float64(sum) / float64(count), nil
}

func FindTopStudent() (int, float64, error) {
	// TODO: Найти студента с самым высоким средним баллом
	// Вернуть ID студента и его средний балл

	var topAverage float64
	var topId int

	for idStudent := range storage.Students {
		var count int
		var sum int
		for _, grades := range grades.Grades[idStudent] {
			count += len(grades)
			for _, val := range grades {
				sum += val
			}
		}
		stdAverage := float64(sum) / float64(count)
		if topAverage < stdAverage {
			topAverage = stdAverage
			topId = idStudent
		}
	}
	return topId, topAverage, nil
}

func GetGradeDistribution(subjectID int) map[int]int {
	// TODO: Вернуть распределение оценок по предмету
	// Мапа: оценка -> количество студентов с этой оценкой
	result := make(map[int]int)

	for _, StudentSubjects := range grades.Grades {
		for _, grade := range StudentSubjects[subjectID] {
			result[grade]++
		}

	}
	return result
}

func CountStudentsWithGrade(subjectID int, targetGrade int) int {
	// TODO: Посчитать сколько студентов имеют конкретную оценку по предмету
	var count int

	for _, StudentSubjects := range grades.Grades {

		for _, grade := range StudentSubjects {
			for _, val := range grade {
				if targetGrade == val {
					count++
					break
				}
			}
		}
	}
	return count
}
