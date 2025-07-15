package main

import "strconv"

type SimpleRemoteControl struct {
	onCommands []Command
	offCommands []Command
	undoCommand Command
}

func NewSimpleRemoteControl() *SimpleRemoteControl {
	onCommands := make([]Command, 7)
	offCommands := make([]Command, 7)
	for i := 0; i < 7; i++ {
		onCommands[i] = nil
		offCommands[i] = nil
	}	
	
	return &SimpleRemoteControl{
		onCommands: onCommands,
		offCommands: offCommands,
		undoCommand: nil,
	}
}
func (src *SimpleRemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	src.onCommands[slot] = onCommand
	src.offCommands[slot] = offCommand
}

func (src *SimpleRemoteControl) onButtonWasPressed(slot int) {
	if src.onCommands[slot] != nil {
		src.onCommands[slot].Execute()
		src.undoCommand = src.onCommands[slot]
	} else {
		println("No command set")
	}
}

func (src *SimpleRemoteControl) offButtonWasPressed(slot int) {
	if src.offCommands[slot] != nil {
		src.offCommands[slot].Execute()
		src.undoCommand = src.offCommands[slot]
	} else {
		println("No command set")
	}
}

func (src *SimpleRemoteControl) undoButtonWasPressed() {
	if src.undoCommand != nil {
		src.undoCommand.Undo()
	} else {
		println("No command to undo")
	}
}

func (src *SimpleRemoteControl) toString() string {
	var str string
	str += "\n------ Remote Control ------\n"
	for i := 0; i < len(src.onCommands); i++ {
		str += "[slot " + strconv.Itoa(i) + "] " 
		if src.onCommands[i] != nil {
			str += src.onCommands[i].ToString() + "\n"
		} else {
			str += "No command set\n"
		}
	}
	return str
}
