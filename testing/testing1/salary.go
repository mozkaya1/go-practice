package salary

func calculateSalary(basic float32, hra float32, allowance float32) float32 {
	totalSalary := basic + hra + allowance
	return totalSalary
}
func calculateBonus(basic float32, experience int) float32 {
	var bonus float32
	if experience >= 0 && experience < 1 {
		bonus = basic * float32(0.05)
	} else if experience >= 1 && experience < 2 {
		bonus = basic * float32(0.10)
	} else {
		bonus = basic * float32(0.20)
	}
	return bonus
}
