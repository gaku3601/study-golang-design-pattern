package main

type AgeComparator struct{}

func (self *AgeComparator) compare(human1 *Human, human2 *Human) int {
	if human1.Age > human2.Age {
		return 1
	} else if human1.Age == human2.Age {
		return 0
	} else {
		return -1
	}
}
