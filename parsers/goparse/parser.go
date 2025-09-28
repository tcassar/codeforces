// Package goparse provides common parsing methods for codeforces problems
package goparse

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ParseInt(reader *bufio.Reader) (int, error) {
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

func ParseSliceOfInt(reader *bufio.Reader) ([]int, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read list of shop prices: %w", err)
	}

	asStrings := strings.Fields(strings.TrimSpace(line))
	asInts := make([]int, 0, len(asStrings))

	for _, strPrice := range asStrings {
		sInt, err := strconv.Atoi(strPrice)
		if err != nil {
			return nil, fmt.Errorf("failed to parse shop price '%s': %w", strPrice, err)
		}
		asInts = append(asInts, sInt)
	}

	return asInts, nil
}
