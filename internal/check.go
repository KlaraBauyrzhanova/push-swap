package internal

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

var instructions = []string{"pa", "pb", "sa", "ra", "rr", "ss", "rrr", "rra", "rb", "sb", "rrb"}

func ValidateInstructions() ([]string, error) {
	comms := bufio.NewScanner(os.Stdin)
	input := []string{}
	for comms.Scan() {
		line := comms.Text()
		if len(line) == 0 {
			break
		}
		ok, command := supported(line)
		if !ok {
			return input, errors.New("Error: " + command + " instruction not supported\n")
		}
		input = append(input, line)
	}
	if len(input) < 1 {
		return input, errors.New("")
	}
	return input, nil
}

func supported(input string) (bool, string) {
	for _, instruction := range instructions {
		if instruction == input {
			return true, ""
		}
	}
	return false, input
}

func ValidateInput() error {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		return errors.New("")
	}
	if len(os.Args) == 2 {
		arr := strings.Split(os.Args[1], " ")
		if haveDuplicates(arr) {
			return errors.New("Error\n")
		}
		for i := len(arr) - 1; i >= 0; i-- {
			num, err := strconv.Atoi(arr[i])
			if err != nil {
				return errors.New("Error\n")
			}
			a.push(num)
		}
		if a.isSort() {
			return errors.New("")
		}
	}
	return nil
}

func haveDuplicates(arr []string) bool {
	nums := make(map[string]int, len(arr))
	for idx := 0; idx < len(arr); idx++ {
		if _, ok := nums[arr[idx]]; ok {
			return true
		}
		nums[arr[idx]]++
	}
	return false
}
