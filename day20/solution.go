package day20

import (
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"os"
	"path/filepath"
	"strings"
)

func solution(fileName string) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)
	println("")

	// Create a scanner to read the file line by line
	raw, _ := os.ReadFile(filePath)
	data := string(raw)
	system := parseSystem(data)
	lowPulse := 0
	highPulse := 0
	for i := 0; i < 1000; i++ {
		system.button.Send(Low)
	}
	for _, command := range system.sendPulseCommands {
		if command.Pulse == Low {
			lowPulse++
		} else if command.Pulse == High {
			highPulse++
		}
		//fmt.Printf("%s -%s-> %s\n", command.Sender, command.Pulse, command.Receiver)
	}
	total := lowPulse * highPulse

	fmt.Printf("\nTotal: %+v\n", total)
	return total
}

type Module struct {
	Type        Type
	Name        string
	Destination []*Module
	InputModule map[string]Pulse
	Status      Status
	OutputPulse Pulse
	System      *System
}

func (module *Module) Send(pulse Pulse) {
	module.System.SendPulse(module, pulse)
}

func (module *Module) Receive(pulse Pulse, name string) {
	switch module.Type {
	case Broadcaster:
		module.Send(pulse)
	case Flipflop:
		if pulse == High {
			return
		}
		module.Status = !module.Status
		module.Send(FromStatus(module.Status))
	case Conjuntion:
		module.InputModule[name] = pulse
		values := lo.Map(maps.Values(module.InputModule), func(pulse Pulse, _ int) int {
			if pulse == Low {
				return 0
			} else {
				return 1
			}
		})
		sum := lo.Sum(values)
		if sum == 0 {
			module.OutputPulse = High
		} else if sum == len(values) {
			module.OutputPulse = Low
		}
		module.Send(module.OutputPulse)
	}
}

type Pulse string

type Status bool

type Type int32

const (
	Broadcaster = 'b'
	Flipflop    = '%'
	Conjuntion  = '&'
)

const (
	Low  = "low"
	High = "high"
)

const (
	Off = false
	On  = true
)

func FromStatus(status Status) Pulse {
	if status == Off {
		return Low
	} else {
		return High
	}
}

type System struct {
	moduleMap         map[string]*Module
	button            *Module
	sendPulseCommands []SendPulseCommand
}

func (system *System) SendPulse(module *Module, pulse Pulse) {
	for _, destination := range module.Destination {
		system.sendPulseCommands = append(system.sendPulseCommands, SendPulseCommand{module.Name, pulse, destination.Name})
	}
	for _, destination := range module.Destination {
		(*destination).Receive(pulse, module.Name)
	}
}

func parseSystem(data string) *System {
	system := System{moduleMap: map[string]*Module{}}
	modulesStr := strings.Split(data, "\n")
	var destinations []DestinationMapping
	for _, moduleStr := range modulesStr {
		parts := strings.Split(moduleStr, " -> ")
		nameStr := parts[0]
		name := ""
		destinationsStr := parts[1]
		destinationsName := strings.Split(destinationsStr, ", ")
		var module Module

		switch nameStr[:1] {
		case "b":
			name = nameStr
			module = Module{Type: Broadcaster}
		case "%":
			name = nameStr[1:]
			module = Module{Type: Flipflop, Status: Off}
		case "&":
			name = nameStr[1:]
			module = Module{Type: Conjuntion, OutputPulse: High, InputModule: map[string]Pulse{}}
		}
		module.System = &system
		module.Name = name
		system.moduleMap[module.Name] = &module
		destinations = append(destinations, DestinationMapping{&module, destinationsName})
	}
	for _, destinationMapping := range destinations {
		module := destinationMapping.Module
		if module.Type == Broadcaster {
			button := Module{Type: Broadcaster, Name: "button", System: &system, Destination: []*Module{module}}
			system.moduleMap[button.Name] = &button
			system.button = &button
		}
		for _, name := range destinationMapping.Destinations {
			destination, exists := system.moduleMap[name]
			if !exists {
				destination = &Module{Type: Broadcaster, Name: "output", System: &system}
				system.moduleMap[name] = destination
			}
			module.Destination = append(module.Destination, destination)
			if destination.Type == Conjuntion {
				destination.InputModule[module.Name] = Low
			}
		}
	}
	return &system
}

type DestinationMapping struct {
	Module       *Module
	Destinations []string
}

type SendPulseCommand struct {
	Sender   string
	Pulse    Pulse
	Receiver string
}
