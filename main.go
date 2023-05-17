package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	filePath := "log1.txt"
	logFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer logFile.Close()

	menuCount := make(map[string]int)

	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		line := scanner.Text()

		ids := strings.Split(line, ".")

		if len(ids) != 2 {
			log.Fatalf("Invalid log entry: %s", line)
		}

		//eaterID := ids[0]
		foodmenuID := ids[1]

		if menuCount[foodmenuID] > 0 {
			log.Fatalf("Duplicate entry for foodmenu_id: %s", foodmenuID)
		}
		menuCount[foodmenuID]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading log file: %s", err)
	}
	topMenuItems := findTopMenuItems(menuCount, 3)
	fmt.Println("Top 3 Menu Items Consumed:")
	for i, menuItem := range topMenuItems {
		fmt.Printf("%d. %s\n", i+1, menuItem)
	}
}

func findTopMenuItems(menuCount map[string]int, n int) []string {
	sortedMenuItems := make([]string, 0, len(menuCount))
	for menuItem := range menuCount {
		sortedMenuItems = append(sortedMenuItems, menuItem)
	}
	sortMenuItems(sortedMenuItems, menuCount)
	if len(sortedMenuItems) <= n {
		return sortedMenuItems
	}
	return sortedMenuItems[:n]
}

func sortMenuItems(menuItems []string, menuCount map[string]int) {
	sort.Slice(menuItems, func(i, j int) bool {
		return menuCount[menuItems[i]] > menuCount[menuItems[j]]
	})
}
