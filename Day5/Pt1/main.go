package main

import (
	"fmt"
	"strings"
)

func main() {
	testInput := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	input := strings.Split(testInput, "\n\n")
	rulesString, pagesString := input[0], input[1]

	rules := textToRules(rulesString)
	pages := textToPages(pagesString)
	fmt.Println(rules)
	fmt.Println(pages)

}

type Rule struct {
	First  string
	Second string
}

func textToPages(rawPages string) [][]string {
	var pages [][]string
	pageLines := strings.Split(rawPages, "\n")

	for _, page := range pageLines {
		pages = append(pages, strings.Split(page, ","))
	}
	return pages
}

func textToRules(rawRules string) []Rule {
	var rules []Rule
	for _, rawRule := range strings.Split(rawRules, "\n") {
		rule := strings.Split(rawRule, "|")
		newRule := Rule{
			First:  rule[0],
			Second: rule[1],
		}
		rules = append(rules, newRule)
	}
	return rules
}

func applicableRules(rules []Rule, pages [][]string) []Rule {
	var validRules []Rule
	for _, page := range pages {
		for _, rule := range rules {
			if doesRuleApply(rule, page) {
				validRules = append(validRules, rule)
			}
		}
	}
	return validRules
}

func doesRuleApply(rule Rule, page []string) bool {
	count := 0
	for _, word := range page {
		if rule.First == word || rule.Second == word {
			count += 1
		}
		if count >= 2 {
			return true
		}
	}
	return false
}
