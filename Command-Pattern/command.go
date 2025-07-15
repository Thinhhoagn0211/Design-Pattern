package main

import "fmt"

type Light struct {
	room string
}

func NewLight(room string) *Light {
	return &Light{
		room: room,
	}
}

func (l *Light) On() {
	fmt.Println("Light is ON in", l.room)
}

func (l *Light) Off() { 
	fmt.Println("Light is OFF", l.room)
}

type LightOnCommand struct {
 	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}

func (c *LightOnCommand) ToString() string {
	return "LightOnCommand"
}

type LightOffCommand struct {
	light *Light 
	}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}

func (c *LightOffCommand) ToString() string {
	return "LightOffCommand"
}

type GarageDoor struct {
	room string
}

func NewGarageDoor(room string) *GarageDoor {
	return &GarageDoor{
		room: room,
	}
}

func (g *GarageDoor) Up() {
	fmt.Println("Garage Door is Open in", g.room)
}

func (g *GarageDoor) Down() {
	fmt.Println("Garage Door is Closed in", g.room)
}

func (g *GarageDoor) Stop() {
	fmt.Println("Garage Door is Stopped in", g.room)
}

func (g *GarageDoor) LightOn() {
	fmt.Println("Garage Door Light is ON in", g.room)
}

func (g *GarageDoor) LightOff() {
	fmt.Println("Garage Door Light is OFF in", g.room)
}

type GarageDoorOpenCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorOpenCommand(garageDoor *GarageDoor) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{garageDoor: garageDoor}
}

func (c *GarageDoorOpenCommand) Execute() {
	c.garageDoor.Up()
	c.garageDoor.LightOn()
}

func (c *GarageDoorOpenCommand) Undo() {
	c.garageDoor.Down()
	c.garageDoor.LightOff()
}

func (c *GarageDoorOpenCommand) ToString() string {
	return "GarageDoorOpenCommand"
}

type GarageDoorCloseCommand struct {
	garageDoor *GarageDoor
}

func NewGarageDoorCloseCommand(garageDoor *GarageDoor) *GarageDoorCloseCommand {
	return &GarageDoorCloseCommand{garageDoor: garageDoor}
}

func (c *GarageDoorCloseCommand) Execute() {
	c.garageDoor.Down()
	c.garageDoor.LightOff()
}

func (c *GarageDoorCloseCommand) Undo() {
	c.garageDoor.Up()
	c.garageDoor.LightOn()
}

func (c *GarageDoorCloseCommand) ToString() string {
	return "GarageDoorCloseCommand"
}

const  (
	High = iota
	Medium
	Low
	Off
)

type CeilingFan struct {
	room string
	speed int 
 }

func NewCeilingFan(room string) *CeilingFan {
	return &CeilingFan{
		room: room,
	}
}

func (f *CeilingFan) High() {
	f.speed = High
}

func (f *CeilingFan) Medium() {
	f.speed = Medium
}

func (f *CeilingFan) Low() {
	f.speed = Low
}

func (f *CeilingFan) On() {
	fmt.Println("Ceiling Fan is ON in", f.room)
}

func (f *CeilingFan) Off() {
	f.speed = Off
}

func (f *CeilingFan) GetSpeed() int {
	return f.speed
}

type CeilingFanHighCommand struct {
	ceilingFan *CeilingFan
	prevSpeed int
}

func NewCeilingFanHighCommand(ceilingFan *CeilingFan) *CeilingFanHighCommand {
	return &CeilingFanHighCommand{
	ceilingFan: ceilingFan}
}

func (c *CeilingFanHighCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.High()
	fmt.Println("Ceiling Fan is set to HIGH in", c.ceilingFan.room)
}

func (c *CeilingFanHighCommand) Undo() {
	switch c.prevSpeed {
	case High:
		c.ceilingFan.High()
	case Medium:
		c.ceilingFan.Medium()
	case Low:
		c.ceilingFan.Low()
	default:
		c.ceilingFan.Off()
	}
	fmt.Println("Ceiling Fan speed restored to", c.prevSpeed, "in", c.ceilingFan.room)
}

type CeilingFanOnCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOnCommand(ceilingFan *CeilingFan) *CeilingFanOnCommand {
	return &CeilingFanOnCommand{ceilingFan: ceilingFan}
}

func (c *CeilingFanOnCommand) Execute() {
	c.ceilingFan.On()
}

func (c *CeilingFanOnCommand) Undo() {
	c.ceilingFan.Off()
}

func (c *CeilingFanOnCommand) ToString() string {
	return "CeilingFanOnCommand"
}

type CeilingFanOffCommand struct {
	ceilingFan *CeilingFan
}

func NewCeilingFanOffCommand(ceilingFan *CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{ceilingFan: ceilingFan}
}

func (c *CeilingFanOffCommand) Execute() {
	c.ceilingFan.Off()
}

func (c *CeilingFanOffCommand) Undo() {
	c.ceilingFan.On()
}

func (c *CeilingFanOffCommand) ToString() string {
	return "CeilingFanOffCommand"
}

type Stereo struct {
	room string
}

func NewStereo(room string) *Stereo {
	return &Stereo{
		room: room,
	}
}

func (s *Stereo) On() {
	fmt.Println("Stereo is ON in", s.room)
}

func (s *Stereo) Off() {
	fmt.Println("Stereo is OFF in", s.room)
}

type StereoOnWithCDCommand struct {
	stereo *Stereo
}

func NewStereoOnWithCDCommand(stereo *Stereo) *StereoOnWithCDCommand {
	return &StereoOnWithCDCommand{stereo: stereo}
}

func (c *StereoOnWithCDCommand) Execute() {
	c.stereo.On()
	fmt.Println("Stereo is playing CD in", c.stereo.room)
}

func (c *StereoOnWithCDCommand) Undo() {
	c.stereo.Off()
	fmt.Println("Stereo is stopped in", c.stereo.room)
}

func (c *StereoOnWithCDCommand) ToString() string {
	return "StereoOnWithCDCommand"
}

type StereoOffCommand struct {
	stereo *Stereo
}

func NewStereoOffCommand(stereo *Stereo) *StereoOffCommand {
	return &StereoOffCommand{stereo: stereo}
}

func (c *StereoOffCommand) Execute() {
	c.stereo.Off()
	fmt.Println("Stereo is stopped in", c.stereo.room)
}

func (c *StereoOffCommand) Undo() {
	c.stereo.On()
	fmt.Println("Stereo is playing CD in", c.stereo.room)
}

func (c *StereoOffCommand) ToString() string { 
	return "StereoOffCommand"
}

type MacroCommand struct {
	commands []Command
}

func NewMacroCommand(commands []Command) *MacroCommand {
	return &MacroCommand{commands: commands}
}

func (m *MacroCommand) Execute() {
	for _, command := range m.commands {
		if command != nil {
			command.Execute()
		}
	}
}

func (m *MacroCommand) Undo() {
	for i := len(m.commands) - 1; i >= 0; i-- {
		if m.commands[i] != nil {
			m.commands[i].Undo()
		}
	}
}

func (m *MacroCommand) ToString() string {
	str := "MacroCommand: "
	for _, command := range m.commands {
		if command != nil {
			str += command.ToString() + ", "
		}
	}
	return str[:len(str)-2] // Remove trailing comma and space
}

