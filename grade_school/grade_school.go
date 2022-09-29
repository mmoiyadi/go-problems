package school

import "sort"

// Define the Grade and School types here.
type Grade struct{
	level int
	students []string
}

type School struct{
	grades map[int]Grade
	gradesKey []int
}

func New() *School {
	return &School{grades: make(map[int]Grade)}
}

func (s *School) Add(student string, grade int) {
	if g, ok := s.grades[grade]; ok {
		g.students = append(g.students, student)
		sort.Strings(g.students)
		s.grades[grade] = g
	}else{
		g := Grade{level: grade, students: []string{student}}
		s.grades[grade] = g
		s.gradesKey = append(s.gradesKey, grade)
	}
}

func (s *School) Grade(level int) []string {
	return s.grades[level].students
}

func (s *School) Enrollment() []Grade {
	var grades []Grade
	sort.Ints(s.gradesKey)
	for _,gradeKey := range s.gradesKey{
		grade := s.grades[gradeKey]
		grades = append(grades, grade)
	}
	return grades
}
