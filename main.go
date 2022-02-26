package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chrislgardner/battleship/package/battleship"
)

func main() {
	battleship.NewGame(8, 8)

	playGame()
}

func playGame() {

	hits := 0
	for i := 1; i <= 20; i++ {
		fmt.Println("Board:")
		fmt.Println(getBoard())
		fmt.Printf("Shot %d, please enter x,y: ", i)
		var shot string
		fmt.Scanln(&shot)

		res, err := playerFire(shot)
		if err != nil {
			i--
			fmt.Println("Invalid shot entered. Please check and try again in format x,y")
		}

		fmt.Println("")
		fmt.Println(res)
		if res == "hit" {
			hits++
			if hits == 2 {
				fmt.Println("Congratulations, you won.")
				break
			}
		}

		fmt.Println("")
	}

	if hits != 2 {
		fmt.Println("Sorry, you lost. Better luck next time.")
	}
}

func getBoard() string {
	var sb strings.Builder

	for i, row := range battleship.Board {
		if i == 0 {
			sb.WriteString(" ")
			for k := 0; k < len(row); k++ {
				sb.WriteString(fmt.Sprintf("  %d  ", k))
			}
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprint(i))
		for _, col := range row {
			if col == 'M' || col == 'H' {
				sb.WriteString(fmt.Sprintf(" [%s] ", string(col)))
			} else {
				sb.WriteString(" [ ] ")
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	sb.WriteString("H marks a hit, M marks a miss")

	return sb.String()
}

func playerFire(shot string) (string, error) {

	split := strings.Split(shot, ",")

	if len(split) != 2 {
		return "", fmt.Errorf("invalid shot")
	}

	shotX, err := strconv.Atoi(split[0])
	if err != nil {
		return "", fmt.Errorf("invalid shot: %s", err.Error())
	}
	shotY, err := strconv.Atoi(split[1])
	if err != nil {
		return "", fmt.Errorf("invalid shot: %s", err.Error())
	}

	res := battleship.PlayerMove(battleship.Position{X: shotX, Y: shotY})

	return res, nil
}
