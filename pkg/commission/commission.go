package commission

func Calculate(amount int) int {

	minAmount := 500000
	commis := 5
	bonus :=0
	if amount > minAmount {
		bonus = amount * commis / 100000
	} else {
		bonus = 0
	}
	return bonus
}