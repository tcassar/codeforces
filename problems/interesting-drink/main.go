package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func solveTestcase(shops []int, money int) int {
	// offset zero-indexed array with +1
	return sort.SearchInts(shops, money+1)
}

// parseAndSolve solves a single test case
func parseAndSolve(reader *bufio.Reader) error {
	shops, err := parseShops(reader)
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
	return parseInt(reader)
}

func parseMoney(reader *bufio.Reader) (int, error) {
	return parseInt(reader)
}

func parseInt(reader *bufio.Reader) (int, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("failed to read line: %w", err)
	}

	trimmedStr := strings.TrimSpace(str)

	n, err := strconv.Atoi(trimmedStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert '%s' to int: %w", trimmedStr, err)
	}

	return n, nil
}

func parseShops(reader *bufio.Reader) ([]int, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read list of shop prices: %w", err)
	}

	priceStrings := strings.Fields(strings.TrimSpace(line))
	shops := make([]int, 0, len(priceStrings))

	for _, strPrice := range priceStrings {
		sInt, err := strconv.Atoi(strPrice)
		if err != nil {
			return nil, fmt.Errorf("failed to parse shop price '%s': %w", strPrice, err)
		}
		shops = append(shops, sInt)
	}

	return shops, nil
}

func main() {
	start := time.Now()
	reader := bufio.NewReader(os.Stdin)

	_, _ = parseInt(reader)

	if err := parseAndSolve(reader); err != nil {
		log.Fatalf("error while solving: %v", err)
	}

	fmt.Printf("%v\n", time.Since(start))
}
