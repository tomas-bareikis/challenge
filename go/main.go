package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	res := make(map[string][]int)
	var firstKeyOccurrence []string

	for scanner.Scan() {
		f := strings.Fields(scanner.Text())

		var teamName string
		var stepSize int
		var teamData []int

		for i := 0; i < len(f); i++ {
			switch i {
			case 0:
				teamName = string(f[0])

				td, ok := res[teamName]
				if !ok {
					td = make([]int, 2)
					firstKeyOccurrence = append(firstKeyOccurrence, teamName)
				}

				teamData = td
			case 1:
				stepSize, _ = strconv.Atoi(string(f[1]))
			case 2:
				var entrySum int

				for ; i < len(f); i++ {
					steps, _ := strconv.Atoi(f[i])
					if steps == 0 {
						entrySum = 0
						break
					}
					entrySum += steps * stepSize
				}

				if entrySum != 0 {
					teamData[0]++
					teamData[1] += entrySum
				}
				res[teamName] = teamData
			}

		}
	}

	for i := 0; i < len(firstKeyOccurrence); i++ {
		key := firstKeyOccurrence[i]
		teamData, ok := res[key]
		if !ok {
			continue
		}

		delete(res, key)

		distance := teamData[1]

		if distance == 0 {
			continue
		}

		kilometers := math.Round(float64(distance)/1000) / 100.0
		fmt.Fprintf(os.Stdout, "%s %d %.2f\n", key, teamData[0], kilometers)
	}
}
