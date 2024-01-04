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
	rows := strings.Split(data, "\n")
	for v, row := range rows {
		for h, tile := range strings.Split(row, "") {
			loc := Loc{h, v}
			switch tile {
			case "O":
				RRock = append(RRock, loc)
			case "#":
				CRock = append(CRock, loc)
			}
		}
	}

	mapp := Map{RRock, CRock, len(rows[0]), len(rows)}

	for i := 0; i < 200; i++ {
		tilt(&mapp, "up")
		tilt(&mapp, "left")
		tilt(&mapp, "down")
		tilt(&mapp, "right")
	}

	acc := 0

	for _, rock := range mapp.RRock {
		acc += mapp.VLength - rock.V
	}

	//fmt.Printf("Empty Row: %+v\n", emptyRows)
	fmt.Printf("Total: %+v\n", acc)
	return acc
}

func tilt(mapp *Map, direction string) *Map {
	var rrNew []Loc
	for _, rr := range mapp.RRock {
		newLoc := getNewLoc(mapp, rr, direction)
		rrNew = append(rrNew, newLoc)
	}
	mapp.RRock = rrNew
	return mapp
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
		return Loc{loc.H, blocker.V + len(rrInDirection) + 1}
	} else if direction == "down" {
		crInDirection := lo.Filter(mapp.CRock, func(rock Loc, _ int) bool {
			return rock.H == loc.H && rock.V > loc.V
		})
		crInDirection = append(crInDirection, Loc{loc.H, mapp.VLength})
		blocker := lo.MinBy(crInDirection, func(a Loc, b Loc) bool {
			return a.V < b.V
		})
		rrInDirection := lo.Filter(mapp.RRock, func(rock Loc, _ int) bool {
			return rock.H == loc.H && rock.V > loc.V && rock.V < blocker.V
		})
		return Loc{loc.H, blocker.V - len(rrInDirection) - 1}
	} else if direction == "right" {
		crInDirection := lo.Filter(mapp.CRock, func(rock Loc, _ int) bool {
			return rock.V == loc.V && rock.H > loc.H
		})
		crInDirection = append(crInDirection, Loc{mapp.HLength, loc.V})
		blocker := lo.MinBy(crInDirection, func(a Loc, b Loc) bool {
			return a.H < b.H
		})
		rrInDirection := lo.Filter(mapp.RRock, func(rock Loc, _ int) bool {
			return rock.V == loc.V && loc.H < rock.H && rock.H < blocker.H
		})
		return Loc{blocker.H - len(rrInDirection) - 1, loc.V}
	} else if direction == "left" {
		crInDirection := lo.Filter(mapp.CRock, func(rock Loc, _ int) bool {
			return rock.V == loc.V && rock.H < loc.H
		})
		crInDirection = append(crInDirection, Loc{-1, loc.V})
		blocker := lo.MaxBy(crInDirection, func(a Loc, b Loc) bool {
			return a.H > b.H
		})
		rrInDirection := lo.Filter(mapp.RRock, func(rock Loc, _ int) bool {
			return rock.V == loc.V && loc.H > rock.H && rock.H > blocker.H
		})
		return Loc{blocker.H + len(rrInDirection) + 1, loc.V}
	}
	return loc
}

type Map struct {
	RRock   []Loc
	CRock   []Loc
	HLength int
	VLength int
}

type Loc struct {
	H int
	V int
}
