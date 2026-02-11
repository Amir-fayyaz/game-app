package validators

import "strconv"

func IsPhoneNumber(phone string) bool {
	//TODO use regex for full validation
	if len(phone) != 11 {
		return false
	}

	if phone[0:2] != "09" {
		return false
	}

	if _, err := strconv.Atoi(phone[2:]); err != nil {
		return false
	}

	return true
}
