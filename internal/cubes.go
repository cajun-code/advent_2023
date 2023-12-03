package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PlayCubeGame(file string) (int64, int64) {
	var total int64
	var totalPower int64
	data, err := os.ReadFile(file)
	Check(err)
	required := map[string]int64{"red": 12, "green": 13, "blue": 14}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		games := strings.Split(line, ":")
		if len(games) < 2 {
			continue
		}
		id := strings.Split(games[0], " ")[1]
		fmt.Printf("id: %s\n", id)
		if playableCube(games[1], required) {
			idVal, err := strconv.ParseInt(id, 10, 64)
			Check(err)
			total += idVal
		}
		power := calculatePower(games[1])
		totalPower += power
	}
	return total, totalPower
}

func calculatePower(game string) int64 {
	m := make(map[string]int64)
	var power int64
	bags := strings.Split(game, ";")
	for _, bag := range bags {
		fmt.Printf("bag: %s\n", bag)
		cubes := strings.Split(bag, ",")
		for _, cube := range cubes {
			fmt.Printf("cube: %s\n", cube)
			dice := strings.Split(strings.TrimSpace(cube), " ")
			num, err := strconv.ParseInt(dice[0], 10, 64)
			Check(err)
			color := dice[1]
			if m[color] < num {
				m[color] = num
			}
		}
	}
	fmt.Print("min cubes needed: ")
	fmt.Println(m)
	power = 1
	for _, value := range m {
		power *= value
	}
	fmt.Printf("power: %d\n", power)
	return power
}

func playableCube(game string, required map[string]int64) bool {
	m := make(map[string]int64)
	var playable bool
	bags := strings.Split(game, ";")
Loop:
	for _, bag := range bags {
		fmt.Printf("bag: %s\n", bag)
		cubes := strings.Split(bag, ",")
		for _, cube := range cubes {
			fmt.Printf("cube: %s\n", cube)
			dice := strings.Split(strings.TrimSpace(cube), " ")
			num, err := strconv.ParseInt(dice[0], 10, 64)
			Check(err)
			color := dice[1]
			m[color] = num
		}
		playable = m["blue"] <= required["blue"] && m["red"] <= required["red"] && m["green"] <= required["green"]
		fmt.Printf("playable: %t\n", playable)
		if !playable {
			break Loop
		}

	}
	return playable

}
