package urlmatch

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	iradix "github.com/hashicorp/go-immutable-radix"
)

func string2Regexp(pat string) (*regexp.Regexp, error) {
	if len(pat) == 0 {
		return nil, nil
	}
	if strings.HasPrefix(pat, "regexp:") {
		pat = pat[7:]
		if len(pat) == 0 {
			return nil, errors.New("empty regular expression")
		}
	} else {
		pat = `^` + regexp.QuoteMeta(pat) + `$`
	}
	res, err := regexp.Compile(pat)
	if err != nil {
		return nil, errors.New("invalid regular expression")
	}
	return res, nil
}

func (r *Rule) parseScheme() []string {
	if len(r.Scheme) == 0 {
		return nil
	}
	var res []string
	for _, schm := range strings.Split(r.Scheme, ",") {
		schm = strings.TrimSpace(schm)
		schm = strings.ToLower(schm)
		if len(schm) != 0 {
			res = append(res, schm)
		}
	}
	return res
}

func (r *Rule) parsePath() (*regexp.Regexp, error) {
	if len(r.Path) == 0 || r.Path == "/" {
		return nil, nil
	}

	res, err := string2Regexp(r.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Path, path: %v, err: %w", r.Path, err)
	}

	return res, nil
}

type Rule struct {
	Scheme   string `yaml:"scheme,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Path     string `yaml:"path,omitempty"`
	Query    string `yaml:"query,omitempty"`
	Fragment string `yaml:"fragment,omitempty"`
}

type URLMatch struct {
	fqdn *iradix.Tree
}
