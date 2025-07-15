package main

import "fmt"

func main() { 
	remote := NewSimpleRemoteControl()
 	livingRoomLight := NewLight("Living Room")	
	kitchenLight := NewLight("Kitchen")
	ceilingFan := NewCeilingFan("Living Room")
	garageDoor := NewGarageDoor("Garage")
	stereo := NewStereo("Living Room")

	livingRoomLightOnCommand := NewLightOnCommand(livingRoomLight)
	livingRoomLightOffCommand := NewLightOffCommand(livingRoomLight)
	kitchenLightOnCommand := NewLightOnCommand(kitchenLight)
	kitchenLightOffCommand := NewLightOffCommand(kitchenLight)
	ceilingFanOnCommand := NewCeilingFanOnCommand(ceilingFan)
	ceilingFanOffCommand := NewCeilingFanOffCommand(ceilingFan)
	// //
	 garageDoorOpenCommand := NewGarageDoorOpenCommand(garageDoor)
	 garageDoorCloseCommand := NewGarageDoorCloseCommand(garageDoor)
	 // 
	 stereoOnWithCDCommand := NewStereoOnWithCDCommand(stereo)
	 stereoOffCommand := NewStereoOffCommand(stereo)
	 
	remote.SetCommand(0, livingRoomLightOnCommand, livingRoomLightOffCommand)
	remote.SetCommand(1, kitchenLightOnCommand, kitchenLightOffCommand)
	remote.SetCommand(2, ceilingFanOnCommand, ceilingFanOffCommand)
	remote.SetCommand(3, garageDoorOpenCommand, garageDoorCloseCommand)
	remote.SetCommand(4, stereoOnWithCDCommand, stereoOffCommand)

	fmt.Println(remote.toString())
	remote.onButtonWasPressed(0)
	remote.offButtonWasPressed(0)
	remote.onButtonWasPressed(1)
	remote.offButtonWasPressed(1)
	remote.onButtonWasPressed(2)
	remote.offButtonWasPressed(2)
	remote.onButtonWasPressed(3)
	remote.offButtonWasPressed(3)
	remote.onButtonWasPressed(4)
	remote.offButtonWasPressed(4)



	remote.undoButtonWasPressed()

	partyOn := []Command{
		livingRoomLightOnCommand,
		kitchenLightOnCommand,
		ceilingFanOnCommand,
		garageDoorOpenCommand,
		stereoOnWithCDCommand,
	}

	partyOff := []Command{
		livingRoomLightOffCommand,
		kitchenLightOffCommand,
		ceilingFanOffCommand,
		garageDoorCloseCommand,
		stereoOffCommand,
	}

	partyOnMacro := NewMacroCommand(partyOn)
	partyOffMacro := NewMacroCommand(partyOff)

	remote.SetCommand(5, partyOnMacro, partyOffMacro)

	fmt.Println(remote.toString())
	fmt.Println("---- Pushing Macro On ----")
	remote.onButtonWasPressed(5)
	fmt.Println("---- Pushing Macro Off ----")
	remote.offButtonWasPressed(5)
}  
