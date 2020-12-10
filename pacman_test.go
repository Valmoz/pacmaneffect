package pacmaneffect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testType struct {
	a, b string
	c    bool
}

func TestPacmanApply(t *testing.T) {
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

	testApply(t, *p, "-3:-3", []int{}, "Test 9")

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

	testApply(t, *p, "3::-1", []int{4, 3, 2, 1}, "Test 51")

	testApply(t, *p, "4::-1", []int{5, 4, 3, 2, 1}, "Test 52")

	testApply(t, *p, "3:-1:-1", []int{}, "Test 53")

	testApply(t, *p, "6::-1", []int{5, 4, 3, 2, 1}, "Test 54")

	testApply(t, *p, "0:4:0", []int{}, "Test 55")

	map1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	map2 := map[string]testType{
		"z": {"zorro", "zurich", false},
		"w": {"wario", "warsaw", true},
		"b": {"bugs bunny", "bergamo", true},
	}

	targetObj := []interface{}{map1, target, "ciao", map2, true}
	pObj, _ := NewPacman(targetObj)
	resObj, _ := pObj.Apply(NewEffect("1:4"))
	assert.Equal(t, []interface{}{target, "ciao", map2}, resObj, "test obj")
}

func TestPacmanApplyUnbounded(t *testing.T) {
	target := []int{1, 2, 3, 4, 5}
	p, _ := NewPacman(target)

	testApplyUnbounded(t, *p, "1", 2, "Test Unbounded 1")

	testApplyUnbounded(t, *p, "-1", 5, "Test Unbounded 2")

	testApplyUnbounded(t, *p, "1:3", []int{2, 3}, "Test Unbounded 3")

	testApplyUnbounded(t, *p, "3:1", []int{}, "Test Unbounded 4")

	testApplyUnbounded(t, *p, "3:3", []int{}, "Test Unbounded 5")

	testApplyUnbounded(t, *p, "-1:3", []int{5, 1, 2, 3}, "Test Unbounded 6")

	testApplyUnbounded(t, *p, "2:-1", []int{}, "Test Unbounded 7")

	testApplyUnbounded(t, *p, "-3:-1", []int{3, 4}, "Test Unbounded 8")

	testApplyUnbounded(t, *p, "-1:-3", []int{}, "Test Unbounded 10")

	testApplyUnbounded(t, *p, "1:9", []int{2, 3, 4, 5, 1, 2, 3, 4}, "Test Unbounded 11")

	testApplyUnbounded(t, *p, "-6:3", []int{5, 1, 2, 3, 4, 5, 1, 2, 3}, "Test Unbounded 12")

	testApplyUnbounded(t, *p, "-6:9", []int{5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4}, "Test Unbounded 13")

	testApplyUnbounded(t, *p, "9:-6", []int{}, "Test Unbounded 14")

	testApplyUnbounded(t, *p, "2:", []int{3, 4, 5}, "Test Unbounded 15")

	testApplyUnbounded(t, *p, ":2", []int{1, 2}, "Test Unbounded 16")

	testApplyUnbounded(t, *p, ":9", []int{1, 2, 3, 4, 5, 1, 2, 3, 4}, "Test Unbounded 17")

	testApplyUnbounded(t, *p, "-9:", []int{2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, "Test Unbounded 18")

	testApplyUnbounded(t, *p, "-3:", []int{3, 4, 5, 1, 2, 3, 4, 5}, "Test Unbounded 19")

	testApplyUnbounded(t, *p, "9:", []int{}, "Test Unbounded 20")

	testApplyUnbounded(t, *p, ":-3", []int{}, "Test Unbounded 21")

	testApplyUnbounded(t, *p, ":-9", []int{}, "Test Unbounded 22")

	testApplyUnbounded(t, *p, ":", []int{1, 2, 3, 4, 5}, "Test Unbounded 23")

	testApplyUnbounded(t, *p, "0:4:2", []int{1, 3}, "Test Unbounded 24")

	testApplyUnbounded(t, *p, "0:4:-2", []int{}, "Test Unbounded 25")

	testApplyUnbounded(t, *p, "4:0:2", []int{}, "Test Unbounded 26")

	testApplyUnbounded(t, *p, "4:0:-2", []int{5, 3}, "Test Unbounded 27")

	testApplyUnbounded(t, *p, "0:4:1", []int{1, 2, 3, 4}, "Test Unbounded 28")

	testApplyUnbounded(t, *p, "0:4:-1", []int{}, "Test Unbounded 29")

	testApplyUnbounded(t, *p, "4:0:1", []int{}, "Test Unbounded 30")

	testApplyUnbounded(t, *p, "4:0:-1", []int{5, 4, 3, 2}, "Test Unbounded 31")

	testApplyUnbounded(t, *p, "-1:0:2", []int{5}, "Test Unbounded 32")

	testApplyUnbounded(t, *p, "-1:0:-2", []int{}, "Test Unbounded 33")

	testApplyUnbounded(t, *p, "2:-1:2", []int{}, "Test Unbounded 34")

	testApplyUnbounded(t, *p, "-1:2:-2", []int{}, "Test Unbounded 35")

	testApplyUnbounded(t, *p, "2:-1:-2", []int{3, 1}, "Test Unbounded 36")

	testApplyUnbounded(t, *p, "3:-1:1", []int{}, "Test Unbounded 37")

	testApplyUnbounded(t, *p, "3:-1:-1", []int{4, 3, 2, 1}, "Test Unbounded 38")

	testApplyUnbounded(t, *p, "-1:3:1", []int{5, 1, 2, 3}, "Test Unbounded 39")

	testApplyUnbounded(t, *p, "-1:3:-1", []int{}, "Test Unbounded 40")

	testApplyUnbounded(t, *p, "1::", []int{2, 3, 4, 5}, "Test Unbounded 41")

	testApplyUnbounded(t, *p, "1:3:", []int{2, 3}, "Test Unbounded 42")

	testApplyUnbounded(t, *p, "1::2", []int{2, 4}, "Test Unbounded 43")

	testApplyUnbounded(t, *p, ":3:2", []int{1, 3}, "Test Unbounded 44")

	testApplyUnbounded(t, *p, ":3:", []int{1, 2, 3}, "Test Unbounded 45")

	testApplyUnbounded(t, *p, "::2", []int{1, 3, 5}, "Test Unbounded 46")

	testApplyUnbounded(t, *p, "::", []int{1, 2, 3, 4, 5}, "Test Unbounded 47")

	testApplyUnbounded(t, *p, "0::9", []int{1}, "Test Unbounded 48")

	testApplyUnbounded(t, *p, "4::-9", []int{5}, "Test Unbounded 49")

	testApplyUnbounded(t, *p, "-3::-9", []int{}, "Test Unbounded 50")

	testApplyUnbounded(t, *p, "3::-1", []int{4, 3, 2, 1}, "Test Unbounded 51")

	testApplyUnbounded(t, *p, "4::-1", []int{5, 4, 3, 2, 1}, "Test Unbounded 52")

	testApplyUnbounded(t, *p, "3:-1:-1", []int{4, 3, 2, 1}, "Test Unbounded 53")

	testApplyUnbounded(t, *p, "6::-1", []int{2, 1, 5, 4, 3, 2, 1}, "Test Unbounded 54")
}

func testApply(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.Apply(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUnbounded(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUnbounded(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}
