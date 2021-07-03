package urlmatch

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func NewFromRuleYAML(ruleYAML io.Reader) (*URLMatch, error) {
	var rules []*Rule
	err := yaml.NewDecoder(ruleYAML).Decode(&rules)
	if err != nil {
		return nil, fmt.Errorf("failed to decode yaml, err: %w", err)
	}
	fmt.Println(rules)
	return nil, nil
}

func NewRulesfromYAMLPath(YAMLpath string) (*URLMatch, error) {
	rules, err := os.Open(YAMLpath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file, file: %v, err: %w", YAMLpath, err)
	}
	defer rules.Close()

	return NewFromRuleYAML(rules)
}
