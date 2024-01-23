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
	system.broadcaster.Send(Low)
	total := 0

	fmt.Printf("Total: %+v\n", total)
	return total
}

type HandlePulse interface {
	Receive(pulse Pulse, name string)
	Send(pulse Pulse)
}

type Module struct {
	Type        Type
	Name        string
	Destination []*Module
	InputModule map[string]Pulse
	Status      Status
	OutputPulse Pulse
	HandlePulse
}

func (module Module) Send(pulse Pulse) {
	Send(module.Destination, pulse, module.Name)
}

func Send(destinations []*Module, pulse Pulse, name string) {
	for _, destination := range destinations {
		fmt.Printf("%s -%s-> %s\n", name, pulse, destination.Name)
		(*destination).Receive(pulse, name)
	}
}

func (module Module) Receive(pulse Pulse, name string) {
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
	Flipflop    = 'f'
	Conjuntion  = 'c'
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
	moduleMap   map[string]*Module
	broadcaster *Module
}

func parseSystem(data string) System {
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
		if nameStr == "broadcaster" {
			name = nameStr
			module = Module{Type: Broadcaster}
			system.broadcaster = &module
		} else if nameStr[:1] == "%" {
			name = nameStr[1:]
			module = Module{Type: Flipflop, Status: Off}
		} else if nameStr[:1] == "&" {
			name = nameStr[1:]
			module = Module{Type: Conjuntion, OutputPulse: High, InputModule: map[string]Pulse{}}
		}
		module.Name = name
		system.moduleMap[module.Name] = &module
		destinations = append(destinations, DestinationMapping{&module, destinationsName})
	}
	for _, destinationMapping := range destinations {
		module := destinationMapping.Module
		for _, name := range destinationMapping.Destinations {
			destination := system.moduleMap[name]
			module.Destination = append(module.Destination, destination)
			if destination.Type == Conjuntion {
				destination.InputModule[module.Name] = Low
			}
		}
	}
	return system
}

type DestinationMapping struct {
	Module       *Module
	Destinations []string
}
