package statistics

func CalculateStudentAverage(studentID int) (float64, error) {
	// TODO: Посчитать средний балл студента по всем предметам
	// Использовать циклы для обхода всех оценок
}

func CalculateSubjectAverage(subjectID int) (float64, error) {
	// TODO: Посчитать средний балл по предмету
}

func FindTopStudent() (int, float64, error) {
	// TODO: Найти студента с самым высоким средним баллом
	// Вернуть ID студента и его средний балл
}

func GetGradeDistribution(subjectID int) map[int]int {
	// TODO: Вернуть распределение оценок по предмету
	// Мапа: оценка -> количество студентов с этой оценкой
}

func CountStudentsWithGrade(subjectID int, targetGrade int) int {
	// TODO: Посчитать сколько студентов имеют конкретную оценку по предмету
}
