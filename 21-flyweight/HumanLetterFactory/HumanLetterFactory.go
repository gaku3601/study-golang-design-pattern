package HumanLetterFactory

import "fmt"

type humanLetterFactory struct {
	humanLetterList map[string]*humanLetter
}

var single *humanLetterFactory = newHumanLetterFactory()

func newHumanLetterFactory() *humanLetterFactory {
	return &humanLetterFactory{humanLetterList: map[string]*humanLetter{}}
}

//singletonでfactoryを作成
func GetInstance() *humanLetterFactory {
	return single
}

type humanLetter struct {
	Letter string
}

func newHumanLetter(letter string) *humanLetter {
	return &humanLetter{Letter: letter}
}

func (h *humanLetterFactory) GetHumanLetter(letter string) *humanLetter {
	hl := h.humanLetterList[letter]
	if hl == nil {
		hl = newHumanLetter(letter)
		h.humanLetterList[letter] = hl
	}
	fmt.Println(len(h.humanLetterList))
	return hl
}
