package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ScanValue will show given message and return user input
func ScanValue(msg string) (string, error) {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf(msg)
	p1, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	p1 = strings.Trim(p1, "\n")
	return p1, nil
}

// GetRuleSet gives set of winning rules to check on moves
func GetRuleSet() map[int][][]int {
	ruleSet := map[int][][]int{
		0: [][]int{
			[]int{0, 1, 2},
			[]int{0, 3, 6},
			[]int{0, 4, 8},
		},
		1: [][]int{
			[]int{0, 1, 2},
			[]int{1, 4, 7},
		},
		2: [][]int{
			[]int{0, 1, 2},
			[]int{2, 5, 8},
			[]int{2, 4, 6},
		},
		3: [][]int{
			[]int{0, 3, 6},
			[]int{3, 4, 5},
		},
		4: [][]int{
			[]int{1, 4, 7},
			[]int{3, 4, 5},
			[]int{0, 4, 8},
			[]int{2, 4, 6},
		},
		5: [][]int{
			[]int{2, 5, 8},
			[]int{3, 4, 5},
		},
		6: [][]int{
			[]int{6, 7, 8},
			[]int{0, 3, 6},
			[]int{2, 4, 6},
		},
		7: [][]int{
			[]int{6, 7, 8},
			[]int{1, 4, 7},
		},
		8: [][]int{
			[]int{6, 7, 8},
			[]int{0, 4, 8},
			[]int{2, 5, 8},
		},
	}
	return ruleSet
}
