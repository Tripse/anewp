package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"regexp"
)

func isValidURL(s string) bool {
	parsedURL, err := url.Parse(s)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	// Check if there are query parameters.
	return len(parsedURL.RawQuery) > 0
}

func main() {
	// 正则：匹配可能的 URL 结构
	re := regexp.MustCompile(`https?://[^\s]+`)

	seen := make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			if isValidURL(match) {
				if _, exists := seen[match]; !exists {
					fmt.Println(match)
					seen[match] = struct{}{}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
	}
}
