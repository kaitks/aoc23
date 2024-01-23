package day20p2

import (
	"fmt"
	"github.com/gammazero/deque"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"math"
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
	minPress := math.MaxInt
	lsModule := system.moduleMap["ls"]
	for i := 0; i < 10000; i++ {
		system.buttonPressed++
		system.button.Send(Low)
		for {
			if system.SendPulseCommandsQueue.Len() == 0 {
				break
			}
			command := system.SendPulseCommandsQueue.PopFront()
			if command.Receiver == lsModule && command.Pulse == High {
				if _, exists := system.lsLoopMap[command.Sender.Name]; !exists {
					system.lsLoopMap[command.Sender.Name] = system.buttonPressed
				}
			}
			command.Receiver.Receive(command.Pulse, command.Sender.Name)
		}
		if len(system.lsLoopMap) == len(lsModule.InputModule) {
			break
		}
	}

	minPress = lo.Reduce(maps.Values(system.lsLoopMap), func(agg int, item int, _ int) int {
		return agg * item
	}, 1)
	fmt.Printf("\nTotal: %+v\n", minPress)
	return minPress
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
	for _, destination := range module.Destination {
		module.System.SendPulseCommandsQueue.PushBack(SendPulseCommand{module, pulse, destination})
	}
}

func (module *Module) Receive(pulse Pulse, name string) {
	switch module.Type {
	case Broadcaster:
		module.Send(pulse)
	case Flipflop:
		if pulse == Low {
			module.Status = !module.Status
			module.Send(FromStatus(module.Status))
		}
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
		if sum == len(values) {
			module.OutputPulse = Low
		} else {
			module.OutputPulse = High
		}
		module.Send(module.OutputPulse)
	default:
		return
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
	moduleMap              map[string]*Module
	button                 *Module
	SendPulseCommandsQueue *deque.Deque[SendPulseCommand]
	buttonPressed          int
	lsLoopMap              map[string]int
}

func parseSystem(data string) *System {
	system := System{moduleMap: map[string]*Module{}, lsLoopMap: map[string]int{}, SendPulseCommandsQueue: deque.New[SendPulseCommand]()}
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
				destination = &Module{Type: Broadcaster, Name: name, System: &system}
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
	Sender   *Module
	Pulse    Pulse
	Receiver *Module
}
