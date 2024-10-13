package utils

func Escape(str string) string {
	escapedStr := ""
	for _, v := range str {
		if v == '\n' {
			escapedStr += "\\n"
			continue
		}
		if v < 32 {
			escapedStr += "\\" + string(v)
			continue
		}
		switch v {
		case '/', '\\', '"':
			escapedStr += "\\" + string(v)
		default:
			escapedStr += string(v)
		}
	}
	return escapedStr
}
