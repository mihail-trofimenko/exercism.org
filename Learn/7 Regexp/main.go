package main

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	// Регулярное выражение для проверки префикса строки
	validPrefixRegex := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	// Проверяем, соответствует ли строка регулярному выражению
	return validPrefixRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	fmt.Println("Incoming string: ", text)
	// Регулярное выражение для разделения строки
	separatorRegex := regexp.MustCompile(`<[~*=\-]*>`)
	// Делим строку найденными разделителями на срез строк
	fields := separatorRegex.Split(text, -1)
	return fields
}

func CountQuotedPasswords(lines []string) int {
	// Регулярное выражение для поиска 'password' внутри кавычек
	passwordRegex := regexp.MustCompile(`["'][^"']*(?i:password)[^"']*["']`)
	fmt.Println("Incoming lines: ", lines)

	// Счетчик найденных случаев
	count := 0

	// Перебираем строки и проверяем наличие 'password' внутри кавычек
	for _, line := range lines {
		if passwordRegex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	// Regular expression to match "end-of-line" followed by digits without an intervening space
	endOfLineRegex := regexp.MustCompile(`end-of-line\d+`)

	// Remove matches from the input string
	cleanedString := endOfLineRegex.ReplaceAllString(text, "")

	return cleanedString
}

func TagWithUserName(lines []string) []string {
	// Regular expression to match "User " followed by a user name
	userRegex := regexp.MustCompile(`User\s+(\S+)`)

	// Process each line
	for i, line := range lines {
		// Find the user name in the line
		matches := userRegex.FindStringSubmatch(line)

		// If a user name is found, tag the line
		if len(matches) > 1 {
			username := matches[1]
			lines[i] = fmt.Sprintf("[USR] %s %s", username, line)
		}
	}

	return lines
}
