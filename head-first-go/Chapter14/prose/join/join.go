package join

import "strings"

func JoinWithCommas(phrases []string) string {
	switch {
	case len(phrases) == 1:
		return phrases[0]
	
	case len(phrases) == 2:
		return phrases[0] + " and " + phrases[1]

	case len(phrases) >= 3:
		result := strings.Join(phrases[:len(phrases) - 1], ", ")
		result += ", and "
		result += phrases[len(phrases) - 1]
		return result

	default:
		return ""
	}
}