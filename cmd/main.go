package main

import "github.com/araceae101/urlmatch"

func main() {
	_, err := urlmatch.NewRulesfromYAMLPath("rules.yaml")
	if err != nil {
		panic(err)
	}

}
