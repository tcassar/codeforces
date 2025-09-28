package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/tcassar/codeforces/parsers/goparse"
)

func solveTestcase(shops []int, money int) int {
	// offset zero-indexed array with +1
	return sort.SearchInts(shops, money+1)
}

// parseAndSolve solves a single test case
func parseAndSolve(reader *bufio.Reader) error {
	shops, err := goparse.ParseSliceOfInt(reader)
	if err != nil {
		return fmt.Errorf("failed to parse shops: %w", err)
	}

	sort.Ints(shops)

	days, err := parseDays(reader)
	if err != nil {
		return fmt.Errorf("failed to parse number of days: %w", err)
	}

	for i := range days {
		money, err := parseMoney(reader)
		if err != nil {
			return fmt.Errorf("failed to parse money for day %d: %w", i+1, err)
		}

		fmt.Println(solveTestcase(shops, money))
	}

	return nil
}

func parseDays(reader *bufio.Reader) (int, error) {
	return goparse.ParseInt(reader)
}

func parseMoney(reader *bufio.Reader) (int, error) {
	return goparse.ParseInt(reader)
}

func main() {
	start := time.Now()
	reader := bufio.NewReader(os.Stdin)

	_, _ = goparse.ParseInt(reader)

	if err := parseAndSolve(reader); err != nil {
		log.Fatalf("error while solving: %v", err)
	}

	fmt.Printf("%v\n", time.Since(start))
}
