package day14

import (
	"fmt"
	"github.com/samber/lo"
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
	var RRock []Loc
	var CRock []Loc
	vLength := 0
	for v, row := range strings.Split(data, "\n") {
		for h, tile := range strings.Split(row, "") {
			loc := Loc{h, v}
			switch tile {
			case "O":
				RRock = append(RRock, loc)
			case "#":
				CRock = append(CRock, loc)
			}
		}
		vLength = v
	}

	mapp := Map{RRock, CRock}
	rrNew := tilt(&mapp, "up")

	acc := 0

	for _, rock := range rrNew {
		acc += vLength - rock.V + 1
	}

	//fmt.Printf("Empty Row: %+v\n", emptyRows)
	fmt.Printf("Total: %+v\n", acc)
	return acc
}

func tilt(mapp *Map, direction string) []Loc {
	rrNew := []Loc{}
	if direction == "up" {
		for _, rr := range mapp.RRock {
			newLoc := getNewLoc(mapp, rr, direction)
			rrNew = append(rrNew, newLoc)
		}
	}
	return rrNew
}

func getNewLoc(mapp *Map, loc Loc, direction string) Loc {
	if direction == "up" {
		crInDirection := lo.Filter(mapp.CRock, func(rock Loc, _ int) bool {
			return rock.H == loc.H && rock.V < loc.V
		})
		crInDirection = append(crInDirection, Loc{loc.H, -1})
		blocker := lo.MaxBy(crInDirection, func(a Loc, b Loc) bool {
			return a.V > b.V
		})
		rrInDirection := lo.Filter(mapp.RRock, func(rock Loc, _ int) bool {
			return rock.H == loc.H && rock.V < loc.V && rock.V > blocker.V
		})
		return Loc{blocker.H, blocker.V + len(rrInDirection) + 1}
	}
	return loc
}

type Map struct {
	RRock []Loc
	CRock []Loc
}

type Loc struct {
	H int
	V int
}
