package main

import (
	"aspiration/mapper"
	"fmt"
	"strings"
)

func main() {
	// 1
	tString := "Aspiration.com"
	fmt.Println(CapitalizeEveryThirdAlphanumericChar(tString))

	// 2
	s := mapper.NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	const alpha = "abcdefghijklmnopqrstuvwxyz0123456789"
	skipIndex := 0
	if len(s) == 0 {
		return s
	}

	res := strings.Split(s, "")

	// Consider space for non-numeric characters.
	for i := 0; i < len(res); i++ {
		if !strings.Contains(alpha, strings.ToLower(res[i])) {
			skipIndex = skipIndex + 1
			continue
		}

		if (i + 1 - skipIndex) % 3 == 0 {
			res[i] = strings.ToUpper(res[i])
		} else {
			res[i] = strings.ToLower(res[i])
		}
	}

	return strings.Join(res, "")
}
