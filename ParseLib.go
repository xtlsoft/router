package router

import (
	"regexp"
	"strings"
)

type ParsedRule struct {

	Variables []string

	RegExp string

}

func ParseRule(rule string) *ParsedRule {

	parseExp := regexp.MustCompile("{[A-Za-z0-9\\.\\+\\*\\?\\(\\)\\\\\\[\\]\\$\\^\\-]+:?([A-Za-z0-9]+)}")
	
	var variables []string

	regExp := parseExp.ReplaceAllFunc([]byte(rule), func (r []byte) []byte {
		var rr string
		
		for i:=1; i<(len(r) - 1); i++ {
			rr += string(r[i])
		}

		rra := strings.Split(rr, ":")

		var ruleReal string
		var name string

		if len(rra) == 1 {
			ruleReal = "(.*)"
			name = rra[0]
		}else{
			ruleReal = rra[0]
			name = rra[1]
		}

		variables = append(variables, name)

		return []byte(ruleReal)
	})

	return &ParsedRule{
		Variables: variables,
		RegExp: string(regExp),
	}

}