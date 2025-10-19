package grades

// Сделайте этот пакет уже и двигаемся дальше

// Мапа для хранения оценок: StudentID -> SubjectID -> []int (слайс оценок)
var Grades = make(map[int]map[int][]int)

func AddGrade(studentID int, subjectID int, grade int) error {
	// TODO: Добавить оценку студенту по предмету
	// Проверить что студент и предмет существуют
	// Проверить что оценка от 1 до 10
	// Инициализировать вложенные мапы если нужно
}

func GetStudentGrades(studentID int) (map[int][]int, error) {
	// TODO: Вернуть мапу предмет -> оценки для студента
	// Вернуть ошибку если студент не найден
}

func GetSubjectGrades(subjectID int) (map[int][]int, error) {
	// TODO: Вернуть мапу студент -> оценки по предмету
	// Вернуть ошибку если предмет не найден
}

func UpdateGrade(studentID int, subjectID int, oldGrade int, newGrade int) error {
	// TODO: Обновить конкретную оценку
	// Найти oldGrade в слайсе и заменить на newGrade
}

func GetStudentSubjectGrades(studentID int, subjectID int) ([]int, error) {
	// TODO: Вернуть все оценки студента по предмету
}
