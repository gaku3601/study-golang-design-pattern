package student

type Student struct {
	Name           string
	Sex            string
	DeviationValue int //偏差値
}

type StudentList []*Student

func NewStudent(name string, sex string, deviationValue int) *Student {
	return &Student{Name: name, Sex: sex, DeviationValue: deviationValue}
}
