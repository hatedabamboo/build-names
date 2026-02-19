package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func getAdjective() string {
	return adjectives[rand.Intn(len(adjectives))]
}

func getNoun() string {
	return nouns[rand.Intn(len(nouns))]
}

func padLeft(s string, n int) string {
	return strings.Repeat("0", n-len(s)) + s
}

func main() {
	var output []string
	var dateString string

	randIntPtr := flag.Bool("r", false, "Include a random number in the name")
	intLenPtr := flag.Int("rl", 3, "Length of the random number (max 6; default 3)")
	datePtr := flag.Bool("d", false, "Include today's date (yyyy-mm-dd) in the name")
	numberPtr := flag.Int("n", 1, "How many names to generate (default 1)")
	alliteratePtr := flag.Bool("a", false, "Generate alliterative names (ex: abject-animal)")

	flag.Parse()

	if *datePtr {
		dateString = time.Now().Format("2006-01-02")
	}

	if *intLenPtr > 6 {
		fmt.Println("Integer length cannot be greater than 6")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < *numberPtr; i++ {
		adj := getAdjective()
		noun := getNoun()

		if *alliteratePtr && adj[0] != noun[0] {
			for {
				noun = getNoun()
				if adj[0] == noun[0] {
					break
				}
			}
		}

		outputString := adj + "-" + noun

		if *randIntPtr {
			pwr := int(math.Pow(10, float64(*intLenPtr))) - 1
			numAsString := strconv.Itoa(rand.Intn(pwr))
			outputString += "-" + padLeft(numAsString, *intLenPtr)
		}

		if *datePtr {
			outputString += "-" + dateString
		}

		output = append(output, outputString)
	}

	if len(output) > 0 {
		fmt.Println(strings.Join(output, "\n"))
	} else {
		fmt.Println(strings.Join(output, ""))
	}
}
