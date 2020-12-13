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

	testApply(t, *p, "::-1", []int{5, 4, 3, 2, 1}, "Test 56")

	testApply(t, *p, "-10:10", []int{1, 2, 3, 4, 5}, "Test 57")

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

	testApplyUnbounded(t, *p, "0:4:0", []int{}, "Test Unbounded 55")

	testApplyUnbounded(t, *p, "::-1", []int{5, 4, 3, 2, 1}, "Test Unbounded 56")

	testApplyUnbounded(t, *p, "-10:10", []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, "Test Unbounded 57")

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
	resObj, _ := pObj.ApplyUnbounded(NewEffect("1:10"))
	assert.Equal(t, []interface{}{target, "ciao", map2, true, map1, target, "ciao", map2, true}, resObj, "test obj")
}

func TestPacmanApplyString(t *testing.T) {
	target := []string{"1", "2", "3", "4", "5"}
	p, _ := NewPacman(target)

	testApplyString(t, *p, "1", "2", "Test String 1")

	testApplyString(t, *p, "-1", "5", "Test String 2")

	testApplyString(t, *p, "1:3", []string{"2", "3"}, "Test String 3")

	testApplyString(t, *p, "3:1", []string{}, "Test String 4")

	testApplyString(t, *p, "3:3", []string{}, "Test String 5")

	testApplyString(t, *p, "-1:3", []string{}, "Test String 6")

	testApplyString(t, *p, "2:-1", []string{"3", "4"}, "Test String 7")

	testApplyString(t, *p, "-3:-1", []string{"3", "4"}, "Test String 8")

	testApplyString(t, *p, "-3:-3", []string{}, "Test String 9")

	testApplyString(t, *p, "-1:-3", []string{}, "Test String 10")

	testApplyString(t, *p, "1:9", []string{"2", "3", "4", "5"}, "Test String 11")

	testApplyString(t, *p, "-6:3", []string{"1", "2", "3"}, "Test String 12")

	testApplyString(t, *p, "-6:9", []string{"1", "2", "3", "4", "5"}, "Test String 13")

	testApplyString(t, *p, "9:-6", []string{}, "Test String 14")

	testApplyString(t, *p, "2:", []string{"3", "4", "5"}, "Test String 15")

	testApplyString(t, *p, ":2", []string{"1", "2"}, "Test String 16")

	testApplyString(t, *p, ":9", []string{"1", "2", "3", "4", "5"}, "Test String 17")

	testApplyString(t, *p, "-9:", []string{"1", "2", "3", "4", "5"}, "Test String 18")

	testApplyString(t, *p, "-3:", []string{"3", "4", "5"}, "Test String 19")

	testApplyString(t, *p, "9:", []string{}, "Test String 20")

	testApplyString(t, *p, ":-3", []string{"1", "2"}, "Test String 21")

	testApplyString(t, *p, ":-9", []string{}, "Test String 22")

	testApplyString(t, *p, ":", []string{"1", "2", "3", "4", "5"}, "Test String 23")

	testApplyString(t, *p, "0:4:2", []string{"1", "3"}, "Test String 24")

	testApplyString(t, *p, "0:4:-2", []string{}, "Test String 25")

	testApplyString(t, *p, "4:0:2", []string{}, "Test String 26")

	testApplyString(t, *p, "4:0:-2", []string{"5", "3"}, "Test String 27")

	testApplyString(t, *p, "0:4:1", []string{"1", "2", "3", "4"}, "Test String 28")

	testApplyString(t, *p, "0:4:-1", []string{}, "Test String 29")

	testApplyString(t, *p, "4:0:1", []string{}, "Test String 30")

	testApplyString(t, *p, "4:0:-1", []string{"5", "4", "3", "2"}, "Test String 31")

	testApplyString(t, *p, "-1:0:2", []string{}, "Test String 32")

	testApplyString(t, *p, "-1:0:-2", []string{"5", "3"}, "Test String 33")

	testApplyString(t, *p, "2:-1:2", []string{"3"}, "Test String 34")

	testApplyString(t, *p, "-1:2:-2", []string{"5"}, "Test String 35")

	testApplyString(t, *p, "2:-1:-2", []string{}, "Test String 36")

	testApplyString(t, *p, "3:-1:1", []string{"4"}, "Test String 37")

	testApplyString(t, *p, "3:-1:-1", []string{}, "Test String 38")

	testApplyString(t, *p, "-1:3:1", []string{}, "Test String 39")

	testApplyString(t, *p, "-1:3:-1", []string{"5"}, "Test String 40")

	testApplyString(t, *p, "1::", []string{"2", "3", "4", "5"}, "Test String 41")

	testApplyString(t, *p, "1:3:", []string{"2", "3"}, "Test String 42")

	testApplyString(t, *p, "1::2", []string{"2", "4"}, "Test String 43")

	testApplyString(t, *p, ":3:2", []string{"1", "3"}, "Test String 44")

	testApplyString(t, *p, ":3:", []string{"1", "2", "3"}, "Test String 45")

	testApplyString(t, *p, "::2", []string{"1", "3", "5"}, "Test String 46")

	testApplyString(t, *p, "::", []string{"1", "2", "3", "4", "5"}, "Test String 47")

	testApplyString(t, *p, "0::9", []string{"1"}, "Test String 48")

	testApplyString(t, *p, "4::-9", []string{"5"}, "Test String 49")

	testApplyString(t, *p, "-3::-9", []string{"3"}, "Test String 50")

	testApplyString(t, *p, "3::-1", []string{"4", "3", "2", "1"}, "Test String 51")

	testApplyString(t, *p, "4::-1", []string{"5", "4", "3", "2", "1"}, "Test String 52")

	testApplyString(t, *p, "3:-1:-1", []string{}, "Test String 53")

	testApplyString(t, *p, "6::-1", []string{"5", "4", "3", "2", "1"}, "Test String 54")

	testApplyString(t, *p, "0:4:0", []string{}, "Test String 55")

}

func TestPacmanApplyUnboundedString(t *testing.T) {
	target := []string{"1", "2", "3", "4", "5"}
	p, _ := NewPacman(target)

	testApplyUnboundedString(t, *p, "1", "2", "Test Unbounded String 1")

	testApplyUnboundedString(t, *p, "-1", "5", "Test Unbounded String 2")

	testApplyUnboundedString(t, *p, "1:3", []string{"2", "3"}, "Test Unbounded String 3")

	testApplyUnboundedString(t, *p, "3:1", []string{}, "Test Unbounded String 4")

	testApplyUnboundedString(t, *p, "3:3", []string{}, "Test Unbounded String 5")

	testApplyUnboundedString(t, *p, "-1:3", []string{"5", "1", "2", "3"}, "Test Unbounded String 6")

	testApplyUnboundedString(t, *p, "2:-1", []string{}, "Test Unbounded String 7")

	testApplyUnboundedString(t, *p, "-3:-1", []string{"3", "4"}, "Test Unbounded String 8")

	testApplyUnboundedString(t, *p, "-1:-3", []string{}, "Test Unbounded String 10")

	testApplyUnboundedString(t, *p, "1:9", []string{"2", "3", "4", "5", "1", "2", "3", "4"}, "Test Unbounded String 11")

	testApplyUnboundedString(t, *p, "-6:3", []string{"5", "1", "2", "3", "4", "5", "1", "2", "3"}, "Test Unbounded String 12")

	testApplyUnboundedString(t, *p, "-6:9", []string{"5", "1", "2", "3", "4", "5", "1", "2", "3", "4", "5", "1", "2", "3", "4"}, "Test Unbounded String 13")

	testApplyUnboundedString(t, *p, "9:-6", []string{}, "Test Unbounded String 14")

	testApplyUnboundedString(t, *p, "2:", []string{"3", "4", "5"}, "Test Unbounded String 15")

	testApplyUnboundedString(t, *p, ":2", []string{"1", "2"}, "Test Unbounded String 16")

	testApplyUnboundedString(t, *p, ":9", []string{"1", "2", "3", "4", "5", "1", "2", "3", "4"}, "Test Unbounded String 17")

	testApplyUnboundedString(t, *p, "-9:", []string{"2", "3", "4", "5", "1", "2", "3", "4", "5", "1", "2", "3", "4", "5"}, "Test Unbounded String 18")

	testApplyUnboundedString(t, *p, "-3:", []string{"3", "4", "5", "1", "2", "3", "4", "5"}, "Test Unbounded String 19")

	testApplyUnboundedString(t, *p, "9:", []string{}, "Test Unbounded String 20")

	testApplyUnboundedString(t, *p, ":-3", []string{}, "Test Unbounded String 21")

	testApplyUnboundedString(t, *p, ":-9", []string{}, "Test Unbounded String 22")

	testApplyUnboundedString(t, *p, ":", []string{"1", "2", "3", "4", "5"}, "Test Unbounded String 23")

	testApplyUnboundedString(t, *p, "0:4:2", []string{"1", "3"}, "Test Unbounded String 24")

	testApplyUnboundedString(t, *p, "0:4:-2", []string{}, "Test Unbounded String 25")

	testApplyUnboundedString(t, *p, "4:0:2", []string{}, "Test Unbounded String 26")

	testApplyUnboundedString(t, *p, "4:0:-2", []string{"5", "3"}, "Test Unbounded String 27")

	testApplyUnboundedString(t, *p, "0:4:1", []string{"1", "2", "3", "4"}, "Test Unbounded String 28")

	testApplyUnboundedString(t, *p, "0:4:-1", []string{}, "Test Unbounded String 29")

	testApplyUnboundedString(t, *p, "4:0:1", []string{}, "Test Unbounded String 30")

	testApplyUnboundedString(t, *p, "4:0:-1", []string{"5", "4", "3", "2"}, "Test Unbounded String 31")

	testApplyUnboundedString(t, *p, "-1:0:2", []string{"5"}, "Test Unbounded String 32")

	testApplyUnboundedString(t, *p, "-1:0:-2", []string{}, "Test Unbounded String 33")

	testApplyUnboundedString(t, *p, "2:-1:2", []string{}, "Test Unbounded String 34")

	testApplyUnboundedString(t, *p, "-1:2:-2", []string{}, "Test Unbounded String 35")

	testApplyUnboundedString(t, *p, "2:-1:-2", []string{"3", "1"}, "Test Unbounded String 36")

	testApplyUnboundedString(t, *p, "3:-1:1", []string{}, "Test Unbounded String 37")

	testApplyUnboundedString(t, *p, "3:-1:-1", []string{"4", "3", "2", "1"}, "Test Unbounded String 38")

	testApplyUnboundedString(t, *p, "-1:3:1", []string{"5", "1", "2", "3"}, "Test Unbounded String 39")

	testApplyUnboundedString(t, *p, "-1:3:-1", []string{}, "Test Unbounded String 40")

	testApplyUnboundedString(t, *p, "1::", []string{"2", "3", "4", "5"}, "Test Unbounded String 41")

	testApplyUnboundedString(t, *p, "1:3:", []string{"2", "3"}, "Test Unbounded String 42")

	testApplyUnboundedString(t, *p, "1::2", []string{"2", "4"}, "Test Unbounded String 43")

	testApplyUnboundedString(t, *p, ":3:2", []string{"1", "3"}, "Test Unbounded String 44")

	testApplyUnboundedString(t, *p, ":3:", []string{"1", "2", "3"}, "Test Unbounded String 45")

	testApplyUnboundedString(t, *p, "::2", []string{"1", "3", "5"}, "Test Unbounded String 46")

	testApplyUnboundedString(t, *p, "::", []string{"1", "2", "3", "4", "5"}, "Test Unbounded String 47")

	testApplyUnboundedString(t, *p, "0::9", []string{"1"}, "Test Unbounded String 48")

	testApplyUnboundedString(t, *p, "4::-9", []string{"5"}, "Test Unbounded String 49")

	testApplyUnboundedString(t, *p, "-3::-9", []string{}, "Test Unbounded String 50")

	testApplyUnboundedString(t, *p, "3::-1", []string{"4", "3", "2", "1"}, "Test Unbounded String 51")

	testApplyUnboundedString(t, *p, "4::-1", []string{"5", "4", "3", "2", "1"}, "Test Unbounded String 52")

	testApplyUnboundedString(t, *p, "3:-1:-1", []string{"4", "3", "2", "1"}, "Test Unbounded String 53")

	testApplyUnboundedString(t, *p, "6::-1", []string{"2", "1", "5", "4", "3", "2", "1"}, "Test Unbounded String 54")
}

func TestPacmanApplyInt(t *testing.T) {
	target := []int{1, 2, 3, 4, 5}
	p, _ := NewPacman(target)

	testApplyInt(t, *p, "1", 2, "Test Int 1")

	testApplyInt(t, *p, "-1", 5, "Test Int 2")

	testApplyInt(t, *p, "1:3", []int{2, 3}, "Test Int 3")

	testApplyInt(t, *p, "3:1", []int{}, "Test Int 4")

	testApplyInt(t, *p, "3:3", []int{}, "Test Int 5")

	testApplyInt(t, *p, "-1:3", []int{}, "Test Int 6")

	testApplyInt(t, *p, "2:-1", []int{3, 4}, "Test Int 7")

	testApplyInt(t, *p, "-3:-1", []int{3, 4}, "Test Int 8")

	testApplyInt(t, *p, "-3:-3", []int{}, "Test Int 9")

	testApplyInt(t, *p, "-1:-3", []int{}, "Test Int 10")

	testApplyInt(t, *p, "1:9", []int{2, 3, 4, 5}, "Test Int 11")

	testApplyInt(t, *p, "-6:3", []int{1, 2, 3}, "Test Int 12")

	testApplyInt(t, *p, "-6:9", []int{1, 2, 3, 4, 5}, "Test Int 13")

	testApplyInt(t, *p, "9:-6", []int{}, "Test Int 14")

	testApplyInt(t, *p, "2:", []int{3, 4, 5}, "Test Int 15")

	testApplyInt(t, *p, ":2", []int{1, 2}, "Test Int 16")

	testApplyInt(t, *p, ":9", []int{1, 2, 3, 4, 5}, "Test Int 17")

	testApplyInt(t, *p, "-9:", []int{1, 2, 3, 4, 5}, "Test Int 18")

	testApplyInt(t, *p, "-3:", []int{3, 4, 5}, "Test Int 19")

	testApplyInt(t, *p, "9:", []int{}, "Test Int 20")

	testApplyInt(t, *p, ":-3", []int{1, 2}, "Test Int 21")

	testApplyInt(t, *p, ":-9", []int{}, "Test Int 22")

	testApplyInt(t, *p, ":", []int{1, 2, 3, 4, 5}, "Test Int 23")

	testApplyInt(t, *p, "0:4:2", []int{1, 3}, "Test Int 24")

	testApplyInt(t, *p, "0:4:-2", []int{}, "Test Int 25")

	testApplyInt(t, *p, "4:0:2", []int{}, "Test Int 26")

	testApplyInt(t, *p, "4:0:-2", []int{5, 3}, "Test Int 27")

	testApplyInt(t, *p, "0:4:1", []int{1, 2, 3, 4}, "Test Int 28")

	testApplyInt(t, *p, "0:4:-1", []int{}, "Test Int 29")

	testApplyInt(t, *p, "4:0:1", []int{}, "Test Int 30")

	testApplyInt(t, *p, "4:0:-1", []int{5, 4, 3, 2}, "Test Int 31")

	testApplyInt(t, *p, "-1:0:2", []int{}, "Test Int 32")

	testApplyInt(t, *p, "-1:0:-2", []int{5, 3}, "Test Int 33")

	testApplyInt(t, *p, "2:-1:2", []int{3}, "Test Int 34")

	testApplyInt(t, *p, "-1:2:-2", []int{5}, "Test Int 35")

	testApplyInt(t, *p, "2:-1:-2", []int{}, "Test Int 36")

	testApplyInt(t, *p, "3:-1:1", []int{4}, "Test Int 37")

	testApplyInt(t, *p, "3:-1:-1", []int{}, "Test Int 38")

	testApplyInt(t, *p, "-1:3:1", []int{}, "Test Int 39")

	testApplyInt(t, *p, "-1:3:-1", []int{5}, "Test Int 40")

	testApplyInt(t, *p, "1::", []int{2, 3, 4, 5}, "Test Int 41")

	testApplyInt(t, *p, "1:3:", []int{2, 3}, "Test Int 42")

	testApplyInt(t, *p, "1::2", []int{2, 4}, "Test Int 43")

	testApplyInt(t, *p, ":3:2", []int{1, 3}, "Test Int 44")

	testApplyInt(t, *p, ":3:", []int{1, 2, 3}, "Test Int 45")

	testApplyInt(t, *p, "::2", []int{1, 3, 5}, "Test Int 46")

	testApplyInt(t, *p, "::", []int{1, 2, 3, 4, 5}, "Test Int 47")

	testApplyInt(t, *p, "0::9", []int{1}, "Test Int 48")

	testApplyInt(t, *p, "4::-9", []int{5}, "Test Int 49")

	testApplyInt(t, *p, "-3::-9", []int{3}, "Test Int 50")

	testApplyInt(t, *p, "3::-1", []int{4, 3, 2, 1}, "Test Int 51")

	testApplyInt(t, *p, "4::-1", []int{5, 4, 3, 2, 1}, "Test Int 52")

	testApplyInt(t, *p, "3:-1:-1", []int{}, "Test Int 53")

	testApplyInt(t, *p, "6::-1", []int{5, 4, 3, 2, 1}, "Test Int 54")

	testApplyInt(t, *p, "0:4:0", []int{}, "Test Int 55")
}

func TestPacmanApplyUnboundedInt(t *testing.T) {
	target := []int{1, 2, 3, 4, 5}
	p, _ := NewPacman(target)

	testApplyUnboundedInt(t, *p, "1", 2, "Test Unbounded Int 1")

	testApplyUnboundedInt(t, *p, "-1", 5, "Test Unbounded Int 2")

	testApplyUnboundedInt(t, *p, "1:3", []int{2, 3}, "Test Unbounded Int 3")

	testApplyUnboundedInt(t, *p, "3:1", []int{}, "Test Unbounded Int 4")

	testApplyUnboundedInt(t, *p, "3:3", []int{}, "Test Unbounded Int 5")

	testApplyUnboundedInt(t, *p, "-1:3", []int{5, 1, 2, 3}, "Test Unbounded Int 6")

	testApplyUnboundedInt(t, *p, "2:-1", []int{}, "Test Unbounded Int 7")

	testApplyUnboundedInt(t, *p, "-3:-1", []int{3, 4}, "Test Unbounded Int 8")

	testApplyUnboundedInt(t, *p, "-1:-3", []int{}, "Test Unbounded Int 10")

	testApplyUnboundedInt(t, *p, "1:9", []int{2, 3, 4, 5, 1, 2, 3, 4}, "Test Unbounded Int 11")

	testApplyUnboundedInt(t, *p, "-6:3", []int{5, 1, 2, 3, 4, 5, 1, 2, 3}, "Test Unbounded Int 12")

	testApplyUnboundedInt(t, *p, "-6:9", []int{5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4}, "Test Unbounded Int 13")

	testApplyUnboundedInt(t, *p, "9:-6", []int{}, "Test Unbounded Int 14")

	testApplyUnboundedInt(t, *p, "2:", []int{3, 4, 5}, "Test Unbounded Int 15")

	testApplyUnboundedInt(t, *p, ":2", []int{1, 2}, "Test Unbounded Int 16")

	testApplyUnboundedInt(t, *p, ":9", []int{1, 2, 3, 4, 5, 1, 2, 3, 4}, "Test Unbounded Int 17")

	testApplyUnboundedInt(t, *p, "-9:", []int{2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, "Test Unbounded Int 18")

	testApplyUnboundedInt(t, *p, "-3:", []int{3, 4, 5, 1, 2, 3, 4, 5}, "Test Unbounded Int 19")

	testApplyUnboundedInt(t, *p, "9:", []int{}, "Test Unbounded Int 20")

	testApplyUnboundedInt(t, *p, ":-3", []int{}, "Test Unbounded Int 21")

	testApplyUnboundedInt(t, *p, ":-9", []int{}, "Test Unbounded Int 22")

	testApplyUnboundedInt(t, *p, ":", []int{1, 2, 3, 4, 5}, "Test Unbounded Int 23")

	testApplyUnboundedInt(t, *p, "0:4:2", []int{1, 3}, "Test Unbounded Int 24")

	testApplyUnboundedInt(t, *p, "0:4:-2", []int{}, "Test Unbounded Int 25")

	testApplyUnboundedInt(t, *p, "4:0:2", []int{}, "Test Unbounded Int 26")

	testApplyUnboundedInt(t, *p, "4:0:-2", []int{5, 3}, "Test Unbounded Int 27")

	testApplyUnboundedInt(t, *p, "0:4:1", []int{1, 2, 3, 4}, "Test Unbounded Int 28")

	testApplyUnboundedInt(t, *p, "0:4:-1", []int{}, "Test Unbounded Int 29")

	testApplyUnboundedInt(t, *p, "4:0:1", []int{}, "Test Unbounded Int 30")

	testApplyUnboundedInt(t, *p, "4:0:-1", []int{5, 4, 3, 2}, "Test Unbounded Int 31")

	testApplyUnboundedInt(t, *p, "-1:0:2", []int{5}, "Test Unbounded Int 32")

	testApplyUnboundedInt(t, *p, "-1:0:-2", []int{}, "Test Unbounded Int 33")

	testApplyUnboundedInt(t, *p, "2:-1:2", []int{}, "Test Unbounded Int 34")

	testApplyUnboundedInt(t, *p, "-1:2:-2", []int{}, "Test Unbounded Int 35")

	testApplyUnboundedInt(t, *p, "2:-1:-2", []int{3, 1}, "Test Unbounded Int 36")

	testApplyUnboundedInt(t, *p, "3:-1:1", []int{}, "Test Unbounded Int 37")

	testApplyUnboundedInt(t, *p, "3:-1:-1", []int{4, 3, 2, 1}, "Test Unbounded Int 38")

	testApplyUnboundedInt(t, *p, "-1:3:1", []int{5, 1, 2, 3}, "Test Unbounded Int 39")

	testApplyUnboundedInt(t, *p, "-1:3:-1", []int{}, "Test Unbounded Int 40")

	testApplyUnboundedInt(t, *p, "1::", []int{2, 3, 4, 5}, "Test Unbounded Int 41")

	testApplyUnboundedInt(t, *p, "1:3:", []int{2, 3}, "Test Unbounded Int 42")

	testApplyUnboundedInt(t, *p, "1::2", []int{2, 4}, "Test Unbounded Int 43")

	testApplyUnboundedInt(t, *p, ":3:2", []int{1, 3}, "Test Unbounded Int 44")

	testApplyUnboundedInt(t, *p, ":3:", []int{1, 2, 3}, "Test Unbounded Int 45")

	testApplyUnboundedInt(t, *p, "::2", []int{1, 3, 5}, "Test Unbounded Int 46")

	testApplyUnboundedInt(t, *p, "::", []int{1, 2, 3, 4, 5}, "Test Unbounded Int 47")

	testApplyUnboundedInt(t, *p, "0::9", []int{1}, "Test Unbounded Int 48")

	testApplyUnboundedInt(t, *p, "4::-9", []int{5}, "Test Unbounded Int 49")

	testApplyUnboundedInt(t, *p, "-3::-9", []int{}, "Test Unbounded Int 50")

	testApplyUnboundedInt(t, *p, "3::-1", []int{4, 3, 2, 1}, "Test Unbounded Int 51")

	testApplyUnboundedInt(t, *p, "4::-1", []int{5, 4, 3, 2, 1}, "Test Unbounded Int 52")

	testApplyUnboundedInt(t, *p, "3:-1:-1", []int{4, 3, 2, 1}, "Test Unbounded Int 53")

	testApplyUnboundedInt(t, *p, "6::-1", []int{2, 1, 5, 4, 3, 2, 1}, "Test Unbounded Int 54")
}

func TestPacmanApplyUint(t *testing.T) {
	target := []uint{1, 2, 3, 4, 5}
	p, _ := NewPacman(target)

	testApplyUint(t, *p, "1", uint(2), "Test Uint 1")

	testApplyUint(t, *p, "1:5", []uint{2, 3, 4, 5}, "Test Uint 2")

	testApplyUint(t, *p, "5:1:-1", []uint{5, 4, 3}, "Test Uint 3")

	testApplyUint(t, *p, "-1", uint(5), "Test Uint 4")

	testApplyUint(t, *p, "-1:-3:-1", []uint{5, 4}, "Test Uint 5")

	testApplyUint(t, *p, "5:1:0", []uint{}, "Test Uint 6")

	testApplyUint(t, *p, "-10:10", []uint{1, 2, 3, 4, 5}, "Test Uint 7")
}

func TestPacmanApplyUnboundedUint(t *testing.T) {
	target := []uint{1, 2, 3, 4, 5}
	p, _ := NewPacman(target)

	testApplyUnboundedUint(t, *p, "1", uint(2), "Test Unbounded Uint 1")

	testApplyUnboundedUint(t, *p, "1:5", []uint{2, 3, 4, 5}, "Test Unbounded Uint 2")

	testApplyUnboundedUint(t, *p, "5:1:-1", []uint{1, 5, 4, 3}, "Test Unbounded Uint 3")

	testApplyUnboundedUint(t, *p, "-1", uint(5), "Test Unbounded Uint 4")

	testApplyUnboundedUint(t, *p, "-1:-3:-1", []uint{5, 4}, "Test Unbounded Uint 5")

	testApplyUnboundedUint(t, *p, "5:1:0", []uint{}, "Test Unbounded Uint 6")

	testApplyUnboundedUint(t, *p, "-10:10", []uint{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, "Test Unbounded Uint 7")
}

func TestPacmanApplyBool(t *testing.T) {
	target := []bool{false, true, true, true, false, true, true, false}
	p, _ := NewPacman(target)

	testApplyBool(t, *p, "1", true, "Test Bool 1")

	testApplyBool(t, *p, "1:5", []bool{true, true, true, false}, "Test Bool 2")

	testApplyBool(t, *p, "5:1:-1", []bool{true, false, true, true}, "Test Bool 3")

	testApplyBool(t, *p, "-1", false, "Test Bool 4")

	testApplyBool(t, *p, "-1:-3:-1", []bool{false, true}, "Test Bool 5")

	testApplyBool(t, *p, "5:1:0", []bool{}, "Test Bool 6")

	testApplyBool(t, *p, "-10:10", []bool{false, true, true, true, false, true, true, false}, "Test Bool 7")
}

func TestPacmanApplyUnboundedBool(t *testing.T) {
	target := []bool{false, true, true, true, false, true, true, false}
	p, _ := NewPacman(target)

	testApplyUnboundedBool(t, *p, "1", true, "Test Unbounded Bool 1")

	testApplyUnboundedBool(t, *p, "1:5", []bool{true, true, true, false}, "Test Unbounded Bool 2")

	testApplyUnboundedBool(t, *p, "5:1:-1", []bool{true, false, true, true}, "Test Unbounded Bool 3")

	testApplyUnboundedBool(t, *p, "-1", false, "Test Unbounded Bool 4")

	testApplyUnboundedBool(t, *p, "-1:-3:-1", []bool{false, true}, "Test Unbounded Bool 5")

	testApplyUnboundedBool(t, *p, "5:1:0", []bool{}, "Test Unbounded Bool 6")

	testApplyUnboundedBool(t, *p, "-10:10", []bool{true, false, false, true, true, true, false, true, true, false, false, true, true, true, false, true, true, false, false, true}, "Test Unbounded Bool 7")
}

func TestPacmanApplyByte(t *testing.T) {
	target := []byte("Bergamo")
	p, _ := NewPacman(target)

	testApplyByte(t, *p, "1", byte('e'), "Test Byte 1")

	testApplyByte(t, *p, "1:5", []byte("erga"), "Test Byte 2")

	testApplyByte(t, *p, "5:1:-1", []byte("magr"), "Test Byte 3")

	testApplyByte(t, *p, "-1", byte('o'), "Test Byte 4")

	testApplyByte(t, *p, "-1:-3:-1", []byte("om"), "Test Byte 5")

	testApplyByte(t, *p, "5:1:0", []byte{}, "Test Byte 6")

	testApplyByte(t, *p, "-10:10", []byte("Bergamo"), "Test Unbounded Byte 7")
}

func TestPacmanApplyUnboundedByte(t *testing.T) {
	target := []byte("Bergamo")
	p, _ := NewPacman(target)

	testApplyUnboundedByte(t, *p, "1", byte('e'), "Test Unbounded Byte 1")

	testApplyUnboundedByte(t, *p, "1:5", []byte("erga"), "Test Unbounded Byte 2")

	testApplyUnboundedByte(t, *p, "5:1:-1", []byte("magr"), "Test Unbounded Byte 3")

	testApplyUnboundedByte(t, *p, "-1", byte('o'), "Test Unbounded Byte 4")

	testApplyUnboundedByte(t, *p, "-1:-3:-1", []byte("om"), "Test Unbounded Byte 5")

	testApplyUnboundedByte(t, *p, "5:1:0", []byte{}, "Test Unbounded Byte 6")

	testApplyUnboundedByte(t, *p, "-10:10", []byte("amoBergamoBergamoBer"), "Test Unbounded Byte 7")
}

func TestPacmanApplyRune(t *testing.T) {
	target := []rune{'♛', '♠', '♧', '♡', '♬'}
	p, _ := NewPacman(target)

	testApplyRune(t, *p, "1", rune('♠'), "Test Rune 1")

	testApplyRune(t, *p, "1:5", []rune{'♠', '♧', '♡', '♬'}, "Test Rune 2")

	testApplyRune(t, *p, "5:1:-1", []rune{'♬', '♡', '♧'}, "Test Rune 3")

	testApplyRune(t, *p, "-1", rune('♬'), "Test Rune 4")

	testApplyRune(t, *p, "-1:-3:-1", []rune{'♬', '♡'}, "Test Rune 5")

	testApplyRune(t, *p, "5:1:0", []rune{}, "Test Rune 6")

	testApplyRune(t, *p, "-10:10", []rune{'♛', '♠', '♧', '♡', '♬'}, "Test Rune 7")
}

func TestPacmanApplyUnboundedRune(t *testing.T) {
	target := []rune{'♛', '♠', '♧', '♡', '♬'}
	p, _ := NewPacman(target)

	testApplyUnboundedRune(t, *p, "1", rune('♠'), "Test Unbounded Rune 1")

	testApplyUnboundedRune(t, *p, "1:5", []rune{'♠', '♧', '♡', '♬'}, "Test Unbounded Rune 2")

	testApplyUnboundedRune(t, *p, "5:1:-1", []rune{'♛', '♬', '♡', '♧'}, "Test Unbounded Rune 3")

	testApplyUnboundedRune(t, *p, "-1", rune('♬'), "Test Unbounded Rune 4")

	testApplyUnboundedRune(t, *p, "-1:-3:-1", []rune{'♬', '♡'}, "Test Unbounded Rune 5")

	testApplyUnboundedRune(t, *p, "5:1:0", []rune{}, "Test Unbounded Rune 6")

	testApplyUnboundedRune(t, *p, "-10:10", []rune{'♛', '♠', '♧', '♡', '♬', '♛', '♠', '♧', '♡', '♬', '♛', '♠', '♧', '♡', '♬', '♛', '♠', '♧', '♡', '♬'}, "Test Unbounded Rune 7")
}

func TestPacmanApplyFloat32(t *testing.T) {
	target := []float32{1.11, 2.22, 3.33, 4.44, 5.55}
	p, _ := NewPacman(target)

	testApplyFloat32(t, *p, "1", float32(2.22), "Test Float32 1")

	testApplyFloat32(t, *p, "1:5", []float32{2.22, 3.33, 4.44, 5.55}, "Test Float32 2")

	testApplyFloat32(t, *p, "5:1:-1", []float32{5.55, 4.44, 3.33}, "Test Float32 3")

	testApplyFloat32(t, *p, "-1", float32(5.55), "Test Float32 4")

	testApplyFloat32(t, *p, "-1:-3:-1", []float32{5.55, 4.44}, "Test Float32 5")

	testApplyFloat32(t, *p, "5:1:0", []float32{}, "Test Float32 6")

	testApplyFloat32(t, *p, "-10:10", []float32{1.11, 2.22, 3.33, 4.44, 5.55}, "Test Float32 7")
}

func TestPacmanApplyUnboundedFloat32(t *testing.T) {
	target := []float32{1.11, 2.22, 3.33, 4.44, 5.55}
	p, _ := NewPacman(target)

	testApplyUnboundedFloat32(t, *p, "1", float32(2.22), "Test Unbounded Float32 1")

	testApplyUnboundedFloat32(t, *p, "1:5", []float32{2.22, 3.33, 4.44, 5.55}, "Test Unbounded Float32 2")

	testApplyUnboundedFloat32(t, *p, "5:1:-1", []float32{1.11, 5.55, 4.44, 3.33}, "Test Unbounded Float32 3")

	testApplyUnboundedFloat32(t, *p, "-1", float32(5.55), "Test Unbounded Float32 4")

	testApplyUnboundedFloat32(t, *p, "-1:-3:-1", []float32{5.55, 4.44}, "Test Unbounded Float32 5")

	testApplyUnboundedFloat32(t, *p, "5:1:0", []float32{}, "Test Unbounded Float32 6")

	testApplyUnboundedFloat32(t, *p, "-10:10", []float32{1.11, 2.22, 3.33, 4.44, 5.55, 1.11, 2.22, 3.33, 4.44, 5.55, 1.11, 2.22, 3.33, 4.44, 5.55, 1.11, 2.22, 3.33, 4.44, 5.55}, "Test Unbounded Float32 7")
}

func testApply(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.Apply(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUnbounded(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUnbounded(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyString(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyString(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUnboundedString(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUnboundedString(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyInt(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyInt(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUnboundedInt(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUnboundedInt(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUint(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUint(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUnboundedUint(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUnboundedUint(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyBool(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyBool(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUnboundedBool(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUnboundedBool(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyByte(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyByte(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUnboundedByte(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUnboundedByte(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyRune(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyRune(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUnboundedRune(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUnboundedRune(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyFloat32(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyFloat32(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}

func testApplyUnboundedFloat32(t *testing.T, p Pacman, effect string, expected interface{}, message string) {
	result, _ := p.ApplyUnboundedFloat32(NewEffect(effect))
	assert.Equal(t, expected, result, message)
}
