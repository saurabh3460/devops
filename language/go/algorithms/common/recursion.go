package common

func RecureFact(v int) int {
	if v == 0 {
		return 1
	} else {
		return v * RecureFact(v-1)
	}
}
