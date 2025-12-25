package salary

import "testing"

func TestCalculateSalary(t *testing.T) {
	actual := calculateSalary(10000.0, 1200.0, 5000.0)
	expected := 16500.0
	if actual != float32(expected) {
		t.Errorf("Incorrect result: expected=%f, actual=%f", expected, actual)
	}
}
func TestCalculateBonus(t *testing.T) {
	actual := calculateBonus(10000.0, 1)
	expected := 1000.0
	if actual != float32(expected) {
		t.Errorf("Incorrect result: expected=%f, actual=%f", expected, actual)
	}
}
