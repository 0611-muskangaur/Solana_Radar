package res

import "regexp"

// IsValidUUID validates if a string is a valid UUID.
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// IsValidHash checks if a given transaction hash is valid.
func IsValidHash(hash string) bool {
	return len(hash) > 0 && len(hash) <= 255
}
