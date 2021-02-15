package cli

import (
	"regexp"
)

const (
	noNameFlagPrefix = "{noname#"
)

var expTag = regexp.MustCompile("<([A-Za-z][a-zA-Z0-9_]*)>")
var expSpace = regexp.MustCompile(`^.*\s.*$`)

//"--help" "show hello" --flag="hello world" are single string parameters, but not flags
func detectString(s string) (string, bool) {
	if len(s) == 0 {
		return s, false
	}
	isString := false
	ret := s
	//consider --flag="hello world"
	if l := len(ret); l >= 2 {
		if c, cl := ret[0], ret[l-1]; cl == c && (c == '\'' || c == '"') {
			ret = ret[1 : l-1]
			isString = true
		}
	}
	if len(ret) > 0 && !isString && !isFlagLeadByte(ret[0]) && expSpace.MatchString(ret) {
		isString = true
	}

	return ret, isString
}

func (a *App) fnReplaceTag(src string) string {
	switch src {
	// case "<thiscmd>":
	// 	return thisCmd
	case "<appname>":
		return a.Name
	// case "<versiontime>":
	// 	return f.GetVersionTime()
	// case "<versiontag>":
	// 	return f.GetVersionTag()
	case "<version>":
		return a.Version
		// case "<validity>":
		// 	return f.GetValidity()
	}
	return src
}
