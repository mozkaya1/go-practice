package salary

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCalculateSalary(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		actual := calculateSalary(10000.0, 1200.0, 1000.0)
		expected := float32(12200.0)
		assert.Equal(t, expected, actual)
	})
	t.Run("Test case 2", func(t *testing.T) {
		actual := calculateSalary(25000.0, 1500.0, 1000.0)
		expected := float32(27500.0)
		assert.Equal(t, expected, actual)
	})
}
