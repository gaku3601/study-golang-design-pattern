package main

import "fmt"

type light struct {
	state string
}

func newLight() *light {
	return &light{state: "off"}
}
func (l *light) on() {
	l.state = "on"
	fmt.Println("ライトが点きました。")
}
func (l *light) off() {
	l.state = "off"
	fmt.Println("ライトが消えました。")
}
func (l *light) displayState() {
	if l.state == "on" {
		fmt.Println("ライトが点いています。")
	} else {
		fmt.Println("ライトが消えています。")
	}
}
