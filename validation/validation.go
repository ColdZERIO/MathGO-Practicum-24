package validation

func ValidateStudentID(id int) error {
	// TODO: Проверить что ID положительный и не слишком большой
}

func ValidateStudentName(name string) error {
	// TODO: Проверить что имя не пустое, длина от 2 до 50 символов
	// Проверить что имя состоит только из букв
}

func ValidateSubjectID(id int) error {
	// TODO: Проверить что ID положительный
}

func ValidateSubjectName(name string) error {
	// TODO: Проверить что название предмета валидно
}

func ValidateGrade(grade int) error {
	// TODO: Проверить что оценка от 1 до 10
}

func ValidateStudentExists(studentID int) error {
	// TODO: Проверить что студент существует в мапе storage.Students
}

func ValidateSubjectExists(subjectID int) error {
	// TODO: Проверить что предмет существует в мапе storage.Subjects
}
