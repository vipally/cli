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
	escape := 0
	lastQuot := byte(0)

	const lenSep = 1
	out := make([]string, 0, len(s)/3)
	for i := 0; i+lenSep <= len(s); i++ {
		// consider " xxx 'yyy' zzz" as a single string
		// " xxxx yyyy " case, do not include "
		c := s[i]
		if c == '\'' || c == '"' {
			if inString%2 == 0 || lastQuot == c {
				inString++
				escape = 0
				lastQuot = c
			}
		} else {
			if !isSpace(c) {
				escape = 0
			}
		}

		if inString%2 == 0 && isSpace(c) {
			if wordStart == i { //escape continuous space
				wordStart += lenSep
			} else {
				out = append(out, s[wordStart+escape:i-escape])
				wordStart = i + lenSep
				i += (lenSep - 1)
			}
		}
	}

	if wordStart < len(s) {
		out = append(out, s[wordStart+escape:len(s)-escape])
	}

	return out
}

//FormatLineHead ensure all lines of s are lead with linehead string
func FormatLineHead(s, lineHead string) string {
	r := lineHeadExpr.ReplaceAllString(s, lineHead)
	return r
}
