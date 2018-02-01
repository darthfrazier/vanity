package vanity

func validateEnum(field string, enums []string) bool {
	for _, b := range enums {
		if field == b {
			return true
		}
	}
	return false
}
