package pacmaneffect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPacman(t *testing.T) {
	target := []int{1, 2, 3, 4, 5}
	p, _ := NewPacman(target)

	testApply(t, *p, "1", 2, "Test 1")

	testApply(t, *p, "-1", 5, "Test 2")

	testApply(t, *p, "1:3", []int{2, 3}, "Test 3")

	testApply(t, *p, "3:1", []int{}, "Test 4")

	testApply(t, *p, "3:3", []int{}, "Test 5")

	testApply(t, *p, "-1:3", []int{}, "Test 6")

	testApply(t, *p, "2:-1", []int{3, 4}, "Test 7")

	testApply(t, *p, "-3:-1", []int{3, 4}, "Test 8")

	testApply(t, *p, "-1:-3", []int{}, "Test 10")

	testApply(t, *p, "1:9", []int{2, 3, 4, 5}, "Test 11")

	testApply(t, *p, "-6:3", []int{1, 2, 3}, "Test 12")

	testApply(t, *p, "-6:9", []int{1, 2, 3, 4, 5}, "Test 13")

	testApply(t, *p, "9:-6", []int{}, "Test 14")

	testApply(t, *p, "2:", []int{3, 4, 5}, "Test 15")

	testApply(t, *p, ":2", []int{1, 2}, "Test 16")

	testApply(t, *p, ":9", []int{1, 2, 3, 4, 5}, "Test 17")

	testApply(t, *p, "-9:", []int{1, 2, 3, 4, 5}, "Test 18")

	testApply(t, *p, "-3:", []int{3, 4, 5}, "Test 19")

	testApply(t, *p, "9:", []int{}, "Test 20")

	testApply(t, *p, ":-3", []int{1, 2}, "Test 21")

	testApply(t, *p, ":-9", []int{}, "Test 22")

	testApply(t, *p, ":", []int{1, 2, 3, 4, 5}, "Test 23")

	testApply(t, *p, "0:4:2", []int{1, 3}, "Test 24")

	testApply(t, *p, "0:4:-2", []int{}, "Test 25")

	testApply(t, *p, "4:0:2", []int{}, "Test 26")

	testApply(t, *p, "4:0:-2", []int{5, 3}, "Test 27")

	testApply(t, *p, "0:4:1", []int{1, 2, 3, 4}, "Test 28")

	testApply(t, *p, "0:4:-1", []int{}, "Test 29")

	testApply(t, *p, "4:0:1", []int{}, "Test 30")

	testApply(t, *p, "4:0:-1", []int{5, 4, 3, 2}, "Test 31")

	testApply(t, *p, "-1:0:2", []int{}, "Test 32")

	testApply(t, *p, "-1:0:-2", []int{5, 3}, "Test 33")

	testApply(t, *p, "2:-1:2", []int{3}, "Test 34")

	testApply(t, *p, "-1:2:-2", []int{5}, "Test 35")

	testApply(t, *p, "2:-1:-2", []int{}, "Test 36")

	testApply(t, *p, "3:-1:1", []int{4}, "Test 37")

	testApply(t, *p, "3:-1:-1", []int{}, "Test 38")

	testApply(t, *p, "-1:3:1", []int{}, "Test 39")

	testApply(t, *p, "-1:3:-1", []int{5}, "Test 40")

	testApply(t, *p, "1::", []int{2, 3, 4, 5}, "Test 41")

	testApply(t, *p, "1:3:", []int{2, 3}, "Test 42")

	testApply(t, *p, "1::2", []int{2, 4}, "Test 43")

	testApply(t, *p, ":3:2", []int{1, 3}, "Test 44")

	testApply(t, *p, ":3:", []int{1, 2, 3}, "Test 45")

	testApply(t, *p, "::2", []int{1, 3, 5}, "Test 46")

	testApply(t, *p, "::", []int{1, 2, 3, 4, 5}, "Test 47")

	testApply(t, *p, "0::9", []int{1}, "Test 48")

	testApply(t, *p, "4::-9", []int{5}, "Test 49")

	testApply(t, *p, "-3::-9", []int{3}, "Test 50")
}

func testApply(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.Apply(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}
