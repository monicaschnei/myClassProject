package util

// Constants for all supported genders
const (
	FEMALE = "FEMALE"
	MALE   = "MALE"
)

// IsSupportedGender returns true if the gender is supported
func IsSupportedGender(gender string) bool {
	switch gender {
	case FEMALE, MALE:
		return true
	}
	return false
}
