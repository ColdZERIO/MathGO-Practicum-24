package storage

// Мапы для хранения данных
var Students = make(map[int]string)       // ID -> Имя
var Subjects = make(map[int]string)       // ID -> Название предмета
var StudentSubjects = make(map[int][]int) // StudentID -> []SubjectID

func AddStudent(id int, name string) error {
	// TODO: Добавить студента в мапу Students
	// Проверить что ID не занят, вернуть ошибку если занят
}

func AddSubject(id int, name string) error {
	// TODO: Добавить предмет в мапу Subjects
	// Проверить что ID не занят, вернуть ошибку если занят
}

func AssignSubjectToStudent(studentID int, subjectID int) error {
	// TODO: Добавить предмет студенту
	// Проверить что студент и предмет существуют
	// Добавить subjectID в слайс StudentSubjects[studentID]
}

func GetStudentSubjects(studentID int) ([]int, error) {
	// TODO: Вернуть слайс ID предметов студента
	// Вернуть ошибку если студент не найден
}

func GetAllStudents() map[int]string {
	// TODO: Вернуть мапу всех студентов
}
