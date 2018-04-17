package params

import "regexp"

var mailReg = regexp.MustCompile(`^[a-zA-Z0-9\-.]+@[a-zA-Z0-9\-.]+$`)
var creditReg = regexp.MustCompile(`^[0-9]{10}$`)
var zipReg = regexp.MustCompile(`^[0-9]{7}$`)

func isValidateError(value, typ string) bool {
	switch typ {
	case "mail":
		return !mailReg.MatchString(value)
	case "credit":
		return !creditReg.MatchString(value)
	case "zip":
		return !zipReg.MatchString(value)
	default:
		return false
	}
}
