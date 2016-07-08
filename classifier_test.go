package bayesian

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClassifier_Probability(t *testing.T) {
	// setup
	classifier := New()
	clsLaptop := "laptops"
	clsCom := "computers"
	cls := []string{clsLaptop, clsCom}
	classifier.AddClasses(cls)
	// - teach laptops class
	classifier.Learn(clsLaptop, "wifi")
	classifier.Learn(clsLaptop, "wifi")
	classifier.Learn(clsLaptop, "ram")
	classifier.Learn(clsLaptop, "hdd")
	// - teach computers class
	classifier.Learn(clsCom, "wifi")
	classifier.Learn(clsCom, "monitor")
	classifier.Learn(clsCom, "monitor")
	classifier.Learn(clsCom, "ram")
	classifier.Learn(clsCom, "hdd")
	// execute
	scores := classifier.Probability([]string{"wifi", "ram", "hdd"})
	// assert
	assert.Equal(t, 0.24242424242424246, scores["computers"])
	assert.Equal(t, 0.7575757575757576, scores["laptops"])
}
