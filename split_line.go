// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally

package cli

import (
	"regexp"
)

var (
	//new line with any \t
	lineHeadExpr = regexp.MustCompile(`(?m)^[\t ]*`)
)

func isSpace(c byte) bool {
	return (c == ' ' || c == '\t')
}

//SplitLine splits a command-line text separated with any ' ' or '\t'
func SplitCommandLine(s string) []string {
	wordStart := 0
	inString := 0
	lastQuot := byte(0)

	out := make([]string, 0, len(s)/3)
	for i := 0; i < len(s); i++ {
		// consider " xxx 'yyy' zzz" as a single string
		// " xxxx yyyy " case, do not include "
		c := s[i]
		if c == '\'' || c == '"' && (inString%2 == 0 || lastQuot == c) {
			inString++
			lastQuot = c
		}

		if inString%2 == 0 && isSpace(c) {
			if wordStart == i { //escape continuous space
				wordStart += 1
			} else {
				out = append(out, s[wordStart:i])
				wordStart = i + 1
			}
		}
	}

	//NOTE: unclosed quotes will return the left parts
	if wordStart < len(s) {
		out = append(out, s[wordStart:len(s)])
	}

	return out
}

//FormatLineHead ensure all lines of s are lead with linehead string
func FormatLineHead(s, lineHead string) string {
	r := lineHeadExpr.ReplaceAllString(s, lineHead)
	return r
}
