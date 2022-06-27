package leetcode811

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, v := range cpdomains {
		split_result := strings.Split(v, " ")
		counts, _ := strconv.Atoi(split_result[0])
		for _, domain := range helper(strings.Split(split_result[1], ".")) {
			m[domain] += counts
		}
	}
	res := []string{}
	for k, count := range m {
		res = append(res, fmt.Sprintf("%d %s", count, k))
	}
	return res
}

func helper(input []string) []string {
	res := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = strings.Join(input[i:], ".")
	}
	return res
}
