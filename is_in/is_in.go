package is_in

import (
	"fmt"
	"strings"
	
	"is_in/models"
)

type Params []string
type Errors []string

func (p Params) Mapify() map[string]int {
	result := make(map[string]int)
	for _, key := range p {
		result[key] = 1
	}
	return result
}

func (e Errors) Listify() string {
	result := "ОК"
	if len(e) > 0 {
		result = "Failed because:\n"
		for i, v := range e {
			result += fmt.Sprintf("%v. %v\n", i + 1, v)
		}
	}
	return result
}

func Run(string1, string2 string, o models.CLI) (string, Errors) {
	var result = "false"
	var errors Errors
	
	if o.IgnoreCase {
		string1 = strings.ToLower(string1)
		string2 = strings.ToLower(string2)
	}
	
	var what = splitParams(string1, o)
	var in_what = splitParams(string2, o)
	
	if len(what) == 0 {
		errors = append(errors, "First argument is empty!")
	}
	if len(in_what) == 0 {
		errors = append(errors, "Second argument is empty!")
	}
	
	for p, _ := range what {
		if _, found := in_what[p]; !found {
			errors = append(errors, "'" + p + "' is missing")
		}
	}
	
	if len(errors) == 0 {
		result = "true"
	}
	
	return result, errors
}

// splitParams converts a string into a map for future comparisons
func splitParams(str string, o models.CLI) map[string]int {
	var params Params
	
	for _, v := range strings.Split(str, o.Delimiter) {
		if o.Prepare {
			v = strings.Trim(v, `'" `)
		}
		if v != "" {
			params = append(params, v)
		}
	}
	
	return params.Mapify()
}
