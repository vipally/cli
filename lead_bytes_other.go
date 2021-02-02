// +build !windows

package cli

func isFlagLeadByte(c byte) bool {
	return c == '-'
}

func isFlagLead(s string) bool {
	return s == "-" || s == "--"
}
