package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Draw struct {
	Cubes map[string]int
}

type Game struct {
	Id    int
	Draws []Draw
}

type Bag struct {
	Cubes map[string]int
}

func parseGame(arg string) Game {
	var game Game

	matcher, err := regexp.Compile("Game ([0-9]+): (.*)")
	check(err)

	result := matcher.FindAllStringSubmatch(arg, 1)
	id, err := strconv.Atoi(result[0][1])
	game.Id = id

	check(err)
	draws := strings.Split(result[0][2], ";")
	game.Draws = make([]Draw, len(draws))

	for _, d := range draws {
		var draw Draw
		draw.Cubes = make(map[string]int)

		for _, cube := range strings.Split(d, ",") {
			number, color, found := strings.Cut(strings.Trim(cube, " "), " ")

			if found {
				n, err := strconv.Atoi(number)

				check(err)
				draw.Cubes[color] = n
			}
		}

		game.Draws = append(game.Draws, draw)
	}

	return game
}

func assay(game Game, bag Bag) bool {
	for _, draw := range game.Draws {
		for cube, amount := range draw.Cubes {
			if bag.Cubes[cube] < amount {
				return false
			}
		}
	}

	return true
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func power(game Game) int {
	minimumCubes := make(map[string]int)

	for _, draw := range game.Draws {
		for cube, amount := range draw.Cubes {
			val, ok := minimumCubes[cube]
			if ok {
				minimumCubes[cube] = max(val, amount)
			} else {
				minimumCubes[cube] = amount
			}
		}
	}

	power := 1
	for _, minimum := range minimumCubes {
		power *= minimum
	}

	return power
}

func main() {
	var bag Bag
	bag.Cubes = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	idSum := 0
	totalPower := 0

	for scanner.Scan() {
		line := scanner.Text()

		game := parseGame(line)
		if assay(game, bag) {
			idSum += game.Id
		}

		totalPower += power(game)
	}

	fmt.Println("id sum: ", idSum)
	fmt.Println("total power: ", totalPower)
}
