package model

type student struct {
	Name string
	score float64
}

//通过工厂模式解决student首字母小写其他包调用问题

func NewStudent(n string, s float64) *student {
	return &student{
		Name : n,
		score: s,
	}
}

func (s *student) GetStudent() float64 {
	return s.score
}