package main

import (
	"fmt"
	"studentsystem/grades"
	"studentsystem/statistics"
	"studentsystem/storage"
)

func main() {
	// Инициализация тестовых данных
	initTestData()

	// Демонстрация работы системы
	demonstrateSystem()
}

func initTestData() {
	// Добавляем студентов
	storage.AddStudent(1, "Иван Петров")
	storage.AddStudent(2, "Мария Сидорова")
	storage.AddStudent(3, "Алексей Иванов")

	// Добавляем предметы
	storage.AddSubject(1, "Математика")
	storage.AddSubject(2, "Физика")
	storage.AddSubject(3, "История")

	// Назначаем предметы студентам
	storage.AssignSubjectToStudent(1, 1)
	storage.AssignSubjectToStudent(1, 2)
	storage.AssignSubjectToStudent(2, 1)
	storage.AssignSubjectToStudent(2, 3)
	storage.AssignSubjectToStudent(3, 2)
	storage.AssignSubjectToStudent(3, 3)

	// Добавляем оценки
	grades.AddGrade(1, 1, 8)  // Иван, Математика, 8
	grades.AddGrade(1, 1, 9)  // Иван, Математика, 9
	grades.AddGrade(1, 2, 7)  // Иван, Физика, 7
	grades.AddGrade(2, 1, 6)  // Мария, Математика, 6
	grades.AddGrade(2, 3, 10) // Мария, История, 10
	grades.AddGrade(3, 2, 8)  // Алексей, Физика, 8
	grades.AddGrade(3, 3, 9)  // Алексей, История, 9
}

func demonstrateSystem() {
	fmt.Println("=== СИСТЕМА УЧЕТА СТУДЕНТОВ ===\n")

	// Демонстрация storage
	fmt.Println("1. Все студенты:")
	for id, name := range storage.GetAllStudents() {
		fmt.Printf("   %d: %s\n", id, name)
	}

	// Демонстрация grades
	fmt.Println("\n2. Оценки студентов:")
	for studentID := range storage.GetAllStudents() {
		studentGrades, _ := grades.GetStudentGrades(studentID)
		fmt.Printf("   Студент %d: %v\n", studentID, studentGrades)
	}

	// Демонстрация statistics
	fmt.Println("\n3. Статистика:")
	for studentID, name := range storage.GetAllStudents() {
		avg, _ := statistics.CalculateStudentAverage(studentID)
		fmt.Printf("   %s: средний балл %.2f\n", name, avg)
	}

	// Находим лучшего студента
	topStudentID, topAvg, _ := statistics.FindTopStudent()
	topName := storage.GetAllStudents()[topStudentID]
	fmt.Printf("\n4. Лучший студент: %s со средним баллом %.2f\n", topName, topAvg)

	// Распределение оценок по математике
	fmt.Println("\n5. Распределение оценок по Математике:")
	distribution := statistics.GetGradeDistribution(1)
	for grade, count := range distribution {
		fmt.Printf("   Оценка %d: %d студентов\n", grade, count)
	}
}
