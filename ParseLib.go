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

	replaceExp := regexp.MustCompile("[^{}]+") 

	c := 0

	rrules := map[string]string{
		"/" : "\\/",
		"." : "\\.",
		"-" : "\\-",
		"+" : "\\+",
		"(" : "\\(",
		")" : "\\)",
		"[" : "\\[",
		"]" : "\\]",
		"*" : "\\*",
		"^" : "\\^",
		"$" : "\\&",
		"|" : "\\|",
	}

	rule = string( replaceExp.ReplaceAllFunc([]byte(rule), func (r []byte) []byte{
		
		c ++

		if (c % 2) == 1 {
			for k, v := range rrules {
				r = []byte(strings.Replace(string(r), k, v, -1))
			}
			return r
		}else{
			return r
		}

	}) )

	parseExp := regexp.MustCompile("{[A-Za-z0-9\\@\\.\\+\\*\\?\\(\\)\\\\\\[\\]\\$\\^\\-\\|]+:?([A-Za-z0-9]+)}")
	
	var variables []string

	regExp := parseExp.ReplaceAllFunc([]byte(rule), func (r []byte) []byte {
		var rr string
		
		for i:=1; i<(len(r) - 1); i++ {
			rr += string(r[i])
		}

		rra := strings.Split(rr, ":")

		if strings.HasPrefix(rra[0], "@"){
			rra[0] = Shortcuts[rra[0]]
		}

		var ruleReal string
		var name string

		if len(rra) == 1 {
			ruleReal = "(.*)"
			name = rra[0]
		}else{
			ruleReal = "(" + rra[0] + ")"
			name = rra[1]
		}

		variables = append(variables, name)

		return []byte(ruleReal)
	})

	regStr := string(regExp)

	return &ParsedRule{
		Variables: variables,
		RegExp: regStr,
	}

}

func CheckMatch(parsed *ParsedRule, uri string) (matched bool, vars map[string]string) {

	uri = strings.Split(uri, "?")[0]

	reg := regexp.MustCompile(parsed.RegExp)

	var varRslt = map[string]string{}

	varSlice := reg.FindStringSubmatch(uri)

	if varSlice != nil && varSlice[0] == uri {
		for k, v := range parsed.Variables {
			varRslt[v] = varSlice[k + 1]
		}
	}

	return (varSlice != nil && varSlice[0] == uri), varRslt

}