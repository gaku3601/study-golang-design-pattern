package main

type iCommand interface {
	execute()
}

type lightOnCommand struct {
	light *light
}

func newLightOnCommand(light *light) *lightOnCommand {
	return &lightOnCommand{light: light}
}
func (l *lightOnCommand) execute() {
	l.light.on()
}

type lightOffCommand struct {
	light *light
}

func newLightOffCommand(light *light) *lightOffCommand {
	return &lightOffCommand{light: light}
}
func (l *lightOffCommand) execute() {
	l.light.off()
}
