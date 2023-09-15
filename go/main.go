package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type teamData struct {
	entryCount        int
	centimetersWalked float64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	res := make(map[string]teamData)
	var teamOccurrenceOrder []string

	for scanner.Scan() {
		f := strings.Fields(scanner.Text())

		var teamName string
		var stepSize float64
		var teamData teamData

		for i := 0; i < len(f); i++ {
			switch i {
			case 0:
				teamName = f[0]

				td, ok := res[teamName]
				if ok {
					teamData = td
				} else {
					teamOccurrenceOrder = append(teamOccurrenceOrder, teamName)
				}
			case 1:
				stepSize, _ = strconv.ParseFloat(f[1], 64)
			case 2:
				var entrySum float64

				for ; i < len(f); i++ {
					steps, _ := strconv.ParseFloat(f[i], 64)
					if steps == 0 {
						entrySum = 0
						break
					}
					entrySum += steps * stepSize
				}

				if entrySum != 0 {
					teamData.entryCount++
					teamData.centimetersWalked += entrySum
				}
				res[teamName] = teamData
			}

		}
	}

	for i := 0; i < len(teamOccurrenceOrder); i++ {
		key := teamOccurrenceOrder[i]
		teamData, ok := res[key]
		if !ok {
			continue
		}

		delete(res, key)

		distance := teamData.centimetersWalked

		if distance == 0 {
			continue
		}

		kilometers := math.Round(float64(distance)/1000) / 100.0
		fmt.Fprintf(os.Stdout, "%s %d %.2f\n", key, teamData.entryCount, kilometers)
	}
}
