package main

type Student struct {
	Observer []Observer
}

func NewStudent() *Student {
	return &Student{}
}

func (s *Student) addObserver(obs Observer) {
	s.Observer = append(s.Observer, obs)
}
func (s *Student) notifyObservers() {
	for _, v := range s.Observer {
		v.update("通知します！")
	}
}
