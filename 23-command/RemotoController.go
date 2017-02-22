package main

type remotoController struct {
	onCommand  iCommand
	offCommand iCommand
}

func newRemotoController() *remotoController {
	return &remotoController{}
}

func (r *remotoController) setOnCommand(command iCommand) {
	r.onCommand = command
}
func (r *remotoController) setOffCommand(command iCommand) {
	r.offCommand = command
}

func (r *remotoController) pressOnButton() {
	if r.onCommand == nil {
		return
	}
	r.onCommand.execute()
}
func (r *remotoController) pressOffButton() {
	if r.offCommand == nil {
		return
	}
	r.offCommand.execute()
}
