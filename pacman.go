package pacmaneffect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Pacman contains the slice
type Pacman struct {
	slice interface{}
}

// Effect describes the transformation to apply to the slice
type Effect struct {
	start, end, step string
	effectType       EffectType
}

// EffectType describes the possible types of effects to apply
type EffectType int

const (
	// StartType describes an effect with only "start" parameter
	StartType EffectType = iota

	// StartEndType describes an effect with "start" and "end" parameter
	StartEndType

	// StartEndStepType describes an effect with "start", "end" and "step" parameter
	StartEndStepType
)

// NewPacman creates a new Pacman instance containing the slice
func NewPacman(slice interface{}) (*Pacman, error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		return &Pacman{slice}, nil
	default:
		return nil, fmt.Errorf("Invalid slice: \"%s\" is not a slice", reflect.TypeOf(slice))
	}
}

// NewEffect creates a new Effect instance
func NewEffect(effect string) Effect {
	indexes := strings.Split(effect, ":")
	start := indexes[0]
	end := ""
	step := ""
	effectType := StartType
	if len(indexes) >= 2 {
		end = indexes[1]
		effectType = StartEndType
		if len(indexes) == 3 {
			step = indexes[2]
			effectType = StartEndStepType
		}
	}
	return Effect{start, end, step, effectType}
}

/*
 * Interface methods
 */

// Apply applies the effect to the slice
func (p Pacman) Apply(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingle(e, false)
	case StartEndType:
		return p.applyMultiple(e, false)
	case StartEndStepType:
		return p.applyMultiple(e, false)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

// ApplyUnbounded applies the effect to the slice as if it is unbounded
func (p Pacman) ApplyUnbounded(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingle(e, true)
	case StartEndType:
		return p.applyMultiple(e, true)
	case StartEndStepType:
		return p.applyMultiple(e, true)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applySingle(e Effect, unbounded bool) (interface{}, error) {
	start, err := strconv.Atoi(e.start)
	if err != nil {
		return nil, err
	}

	switch reflect.TypeOf(p.slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(p.slice)
		l := s.Len()
		k := start
		if unbounded {
			k = k % l
		}
		if k >= 0 {
			return s.Index(k).Interface(), nil
		}
		return s.Index(l + k).Interface(), nil
	default:
		return nil, fmt.Errorf("Invalid slice: \"%s\" is not a slice", reflect.TypeOf(p.slice))
	}

}

func (p Pacman) applyMultiple(e Effect, unbounded bool) (interface{}, error) {
	switch reflect.TypeOf(p.slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(p.slice)
		l := s.Len()

		start, end, step, err := getLoopParams(e, l, unbounded)
		if err != nil {
			return nil, err
		}

		res := applySlice(p.slice, start, end, step, unbounded)

		return res, nil
	default:
		return nil, fmt.Errorf("Invalid slice: \"%s\" is not a slice", reflect.TypeOf(p.slice))
	}
}

func applySlice(slice interface{}, start, end, step int, unbounded bool) interface{} {
	sliceType := reflect.TypeOf(slice)
	sliceValue := reflect.ValueOf(slice)
	l := sliceValue.Len()

	res := reflect.MakeSlice(sliceType, 0, 0)
	if step == 0 {
		return res.Interface()
	}

	cond := getLoopCondition(end, step)

	for i := start; cond(i); i = i + step {

		if unbounded || (i >= 0 && i < l) {
			k := i % l
			if k < 0 {
				res = reflect.Append(res, sliceValue.Index(l+k))
			} else {
				res = reflect.Append(res, sliceValue.Index(k))
			}
		}
	}

	return res.Interface()
}

/*
 * String methods
 */

// ApplyString applies the effect to the string slice
func (p Pacman) ApplyString(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleString(e, false)
	case StartEndType:
		return p.applyMultipleString(e, false)
	case StartEndStepType:
		return p.applyMultipleString(e, false)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

// ApplyUnboundedString applies the effect to the string slice as if it is unbounded
func (p Pacman) ApplyUnboundedString(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleString(e, true)
	case StartEndType:
		return p.applyMultipleString(e, true)
	case StartEndStepType:
		return p.applyMultipleString(e, true)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applySingleString(e Effect, unbounded bool) (string, error) {
	start, err := strconv.Atoi(e.start)
	if err != nil {
		return "", err
	}

	switch p.slice.(type) {
	case []string:
		s := p.slice.([]string)
		l := len(s)
		k := start
		if unbounded {
			k = k % l
		}
		if k >= 0 {
			return s[k], nil
		}
		return s[l+k], nil
	default:
		return "", fmt.Errorf("Not a []string slice: \"%s\"", reflect.TypeOf(p.slice))
	}

}

func (p Pacman) applyMultipleString(e Effect, unbounded bool) ([]string, error) {
	switch p.slice.(type) {
	case []string:
		s := p.slice.([]string)
		l := len(s)

		start, end, step, err := getLoopParams(e, l, unbounded)
		if err != nil {
			return nil, err
		}

		res := applySliceString(s, start, end, step, unbounded)

		return res, nil
	default:
		return nil, fmt.Errorf("Not a []string slice: \"%s\"", reflect.TypeOf(p.slice))
	}
}

func applySliceString(slice []string, start, end, step int, unbounded bool) []string {
	l := len(slice)

	res := []string{}
	if step == 0 {
		return res
	}

	cond := getLoopCondition(end, step)

	for i := start; cond(i); i = i + step {

		if unbounded || (i >= 0 && i < l) {
			k := i % l
			if k < 0 {
				res = append(res, slice[l+k])
			} else {
				res = append(res, slice[k])
			}
		}
	}

	return res
}

/*
 * Int methods
 */

// ApplyInt applies the effect to the int slice
func (p Pacman) ApplyInt(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleInt(e, false)
	case StartEndType:
		return p.applyMultipleInt(e, false)
	case StartEndStepType:
		return p.applyMultipleInt(e, false)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

// ApplyUnboundedInt applies the effect to the int slice as if it is unbounded
func (p Pacman) ApplyUnboundedInt(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleInt(e, true)
	case StartEndType:
		return p.applyMultipleInt(e, true)
	case StartEndStepType:
		return p.applyMultipleInt(e, true)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applySingleInt(e Effect, unbounded bool) (int, error) {
	start, err := strconv.Atoi(e.start)
	if err != nil {
		return 0, err
	}

	switch p.slice.(type) {
	case []int:
		s := p.slice.([]int)
		l := len(s)
		k := start
		if unbounded {
			k = k % l
		}
		if k >= 0 {
			return s[k], nil
		}
		return s[l+k], nil
	default:
		return 0, fmt.Errorf("Not a []int slice: \"%s\"", reflect.TypeOf(p.slice))
	}

}

func (p Pacman) applyMultipleInt(e Effect, unbounded bool) ([]int, error) {
	switch p.slice.(type) {
	case []int:
		s := p.slice.([]int)
		l := len(s)

		start, end, step, err := getLoopParams(e, l, unbounded)
		if err != nil {
			return nil, err
		}

		res := applySliceInt(s, start, end, step, unbounded)

		return res, nil
	default:
		return nil, fmt.Errorf("Not a []int slice: \"%s\"", reflect.TypeOf(p.slice))
	}
}

func applySliceInt(slice []int, start, end, step int, unbounded bool) []int {
	l := len(slice)

	res := []int{}
	if step == 0 {
		return res
	}

	cond := getLoopCondition(end, step)

	for i := start; cond(i); i = i + step {

		if unbounded || (i >= 0 && i < l) {
			k := i % l
			if k < 0 {
				res = append(res, slice[l+k])
			} else {
				res = append(res, slice[k])
			}
		}
	}

	return res
}

/*
 * Uint methods
 */

// ApplyUint applies the effect to the uint slice
func (p Pacman) ApplyUint(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleUint(e, false)
	case StartEndType:
		return p.applyMultipleUint(e, false)
	case StartEndStepType:
		return p.applyMultipleUint(e, false)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

// ApplyUnboundedUint applies the effect to the uint slice as if it is unbounded
func (p Pacman) ApplyUnboundedUint(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleUint(e, true)
	case StartEndType:
		return p.applyMultipleUint(e, true)
	case StartEndStepType:
		return p.applyMultipleUint(e, true)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applySingleUint(e Effect, unbounded bool) (uint, error) {
	start, err := strconv.Atoi(e.start)
	if err != nil {
		return 0, err
	}

	switch p.slice.(type) {
	case []uint:
		s := p.slice.([]uint)
		l := len(s)
		k := start
		if unbounded {
			k = k % l
		}
		if k >= 0 {
			return s[k], nil
		}
		return s[l+k], nil
	default:
		return 0, fmt.Errorf("Not a []uint slice: \"%s\"", reflect.TypeOf(p.slice))
	}

}

func (p Pacman) applyMultipleUint(e Effect, unbounded bool) ([]uint, error) {
	switch p.slice.(type) {
	case []uint:
		s := p.slice.([]uint)
		l := len(s)

		start, end, step, err := getLoopParams(e, l, unbounded)
		if err != nil {
			return nil, err
		}

		res := applySliceUint(s, start, end, step, unbounded)

		return res, nil
	default:
		return nil, fmt.Errorf("Not a []uint slice: \"%s\"", reflect.TypeOf(p.slice))
	}
}

func applySliceUint(slice []uint, start, end, step int, unbounded bool) []uint {
	l := len(slice)

	res := []uint{}
	if step == 0 {
		return res
	}

	cond := getLoopCondition(end, step)

	for i := start; cond(i); i = i + step {

		if unbounded || (i >= 0 && i < l) {
			k := i % l
			if k < 0 {
				res = append(res, slice[l+k])
			} else {
				res = append(res, slice[k])
			}
		}
	}

	return res
}

/*
 * Bool methods
 */

// ApplyBool applies the effect to the bool slice
func (p Pacman) ApplyBool(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleBool(e, false)
	case StartEndType:
		return p.applyMultipleBool(e, false)
	case StartEndStepType:
		return p.applyMultipleBool(e, false)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

// ApplyUnboundedBool applies the effect to the bool slice as if it is unbounded
func (p Pacman) ApplyUnboundedBool(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleBool(e, true)
	case StartEndType:
		return p.applyMultipleBool(e, true)
	case StartEndStepType:
		return p.applyMultipleBool(e, true)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applySingleBool(e Effect, unbounded bool) (bool, error) {
	start, err := strconv.Atoi(e.start)
	if err != nil {
		return false, err
	}

	switch p.slice.(type) {
	case []bool:
		s := p.slice.([]bool)
		l := len(s)
		k := start
		if unbounded {
			k = k % l
		}
		if k >= 0 {
			return s[k], nil
		}
		return s[l+k], nil
	default:
		return false, fmt.Errorf("Not a []bool slice: \"%s\"", reflect.TypeOf(p.slice))
	}

}

func (p Pacman) applyMultipleBool(e Effect, unbounded bool) ([]bool, error) {
	switch p.slice.(type) {
	case []bool:
		s := p.slice.([]bool)
		l := len(s)

		start, end, step, err := getLoopParams(e, l, unbounded)
		if err != nil {
			return nil, err
		}

		res := applySliceBool(s, start, end, step, unbounded)

		return res, nil
	default:
		return nil, fmt.Errorf("Not a []bool slice: \"%s\"", reflect.TypeOf(p.slice))
	}
}

func applySliceBool(slice []bool, start, end, step int, unbounded bool) []bool {
	l := len(slice)

	res := []bool{}
	if step == 0 {
		return res
	}

	cond := getLoopCondition(end, step)

	for i := start; cond(i); i = i + step {

		if unbounded || (i >= 0 && i < l) {
			k := i % l
			if k < 0 {
				res = append(res, slice[l+k])
			} else {
				res = append(res, slice[k])
			}
		}
	}

	return res
}

/*
 * Byte methods
 */

// ApplyByte applies the effect to the byte slice
func (p Pacman) ApplyByte(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleByte(e, false)
	case StartEndType:
		return p.applyMultipleByte(e, false)
	case StartEndStepType:
		return p.applyMultipleByte(e, false)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

// ApplyUnboundedByte applies the effect to the byte slice as if it is unbounded
func (p Pacman) ApplyUnboundedByte(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleByte(e, true)
	case StartEndType:
		return p.applyMultipleByte(e, true)
	case StartEndStepType:
		return p.applyMultipleByte(e, true)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applySingleByte(e Effect, unbounded bool) (byte, error) {
	start, err := strconv.Atoi(e.start)
	if err != nil {
		return 0, err
	}

	switch p.slice.(type) {
	case []byte:
		s := p.slice.([]byte)
		l := len(s)
		k := start
		if unbounded {
			k = k % l
		}
		if k >= 0 {
			return s[k], nil
		}
		return s[l+k], nil
	default:
		return 0, fmt.Errorf("Not a []byte slice: \"%s\"", reflect.TypeOf(p.slice))
	}

}

func (p Pacman) applyMultipleByte(e Effect, unbounded bool) ([]byte, error) {
	switch p.slice.(type) {
	case []byte:
		s := p.slice.([]byte)
		l := len(s)

		start, end, step, err := getLoopParams(e, l, unbounded)
		if err != nil {
			return nil, err
		}

		res := applySliceByte(s, start, end, step, unbounded)

		return res, nil
	default:
		return nil, fmt.Errorf("Not a []byte slice: \"%s\"", reflect.TypeOf(p.slice))
	}
}

func applySliceByte(slice []byte, start, end, step int, unbounded bool) []byte {
	l := len(slice)

	res := []byte{}
	if step == 0 {
		return res
	}

	cond := getLoopCondition(end, step)

	for i := start; cond(i); i = i + step {

		if unbounded || (i >= 0 && i < l) {
			k := i % l
			if k < 0 {
				res = append(res, slice[l+k])
			} else {
				res = append(res, slice[k])
			}
		}
	}

	return res
}

/*
 * Rune methods
 */

// ApplyRune applies the effect to the rune slice
func (p Pacman) ApplyRune(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleRune(e, false)
	case StartEndType:
		return p.applyMultipleRune(e, false)
	case StartEndStepType:
		return p.applyMultipleRune(e, false)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

// ApplyUnboundedRune applies the effect to the rune slice as if it is unbounded
func (p Pacman) ApplyUnboundedRune(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleRune(e, true)
	case StartEndType:
		return p.applyMultipleRune(e, true)
	case StartEndStepType:
		return p.applyMultipleRune(e, true)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applySingleRune(e Effect, unbounded bool) (rune, error) {
	start, err := strconv.Atoi(e.start)
	if err != nil {
		return 0, err
	}

	switch p.slice.(type) {
	case []rune:
		s := p.slice.([]rune)
		l := len(s)
		k := start
		if unbounded {
			k = k % l
		}
		if k >= 0 {
			return s[k], nil
		}
		return s[l+k], nil
	default:
		return 0, fmt.Errorf("Not a []rune slice: \"%s\"", reflect.TypeOf(p.slice))
	}

}

func (p Pacman) applyMultipleRune(e Effect, unbounded bool) ([]rune, error) {
	switch p.slice.(type) {
	case []rune:
		s := p.slice.([]rune)
		l := len(s)

		start, end, step, err := getLoopParams(e, l, unbounded)
		if err != nil {
			return nil, err
		}

		res := applySliceRune(s, start, end, step, unbounded)

		return res, nil
	default:
		return nil, fmt.Errorf("Not a []rune slice: \"%s\"", reflect.TypeOf(p.slice))
	}
}

func applySliceRune(slice []rune, start, end, step int, unbounded bool) []rune {
	l := len(slice)

	res := []rune{}
	if step == 0 {
		return res
	}

	cond := getLoopCondition(end, step)

	for i := start; cond(i); i = i + step {

		if unbounded || (i >= 0 && i < l) {
			k := i % l
			if k < 0 {
				res = append(res, slice[l+k])
			} else {
				res = append(res, slice[k])
			}
		}
	}

	return res
}

/*
 * Float32 methods
 */

// ApplyFloat32 applies the effect to the float32 slice
func (p Pacman) ApplyFloat32(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleFloat32(e, false)
	case StartEndType:
		return p.applyMultipleFloat32(e, false)
	case StartEndStepType:
		return p.applyMultipleFloat32(e, false)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

// ApplyUnboundedFloat32 applies the effect to the float32 slice as if it is unbounded
func (p Pacman) ApplyUnboundedFloat32(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applySingleFloat32(e, true)
	case StartEndType:
		return p.applyMultipleFloat32(e, true)
	case StartEndStepType:
		return p.applyMultipleFloat32(e, true)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applySingleFloat32(e Effect, unbounded bool) (float32, error) {
	start, err := strconv.Atoi(e.start)
	if err != nil {
		return 0, err
	}

	switch p.slice.(type) {
	case []float32:
		s := p.slice.([]float32)
		l := len(s)
		k := start
		if unbounded {
			k = k % l
		}
		if k >= 0 {
			return s[k], nil
		}
		return s[l+k], nil
	default:
		return 0, fmt.Errorf("Not a []float32 slice: \"%s\"", reflect.TypeOf(p.slice))
	}

}

func (p Pacman) applyMultipleFloat32(e Effect, unbounded bool) ([]float32, error) {
	switch p.slice.(type) {
	case []float32:
		s := p.slice.([]float32)
		l := len(s)

		start, end, step, err := getLoopParams(e, l, unbounded)
		if err != nil {
			return nil, err
		}

		res := applySliceFloat32(s, start, end, step, unbounded)

		return res, nil
	default:
		return nil, fmt.Errorf("Not a []float32 slice: \"%s\"", reflect.TypeOf(p.slice))
	}
}

func applySliceFloat32(slice []float32, start, end, step int, unbounded bool) []float32 {
	l := len(slice)

	res := []float32{}
	if step == 0 {
		return res
	}

	cond := getLoopCondition(end, step)

	for i := start; cond(i); i = i + step {

		if unbounded || (i >= 0 && i < l) {
			k := i % l
			if k < 0 {
				res = append(res, slice[l+k])
			} else {
				res = append(res, slice[k])
			}
		}
	}

	return res
}

/*
 * Common methods
 */

func getLoopParams(e Effect, length int, unbounded bool) (int, int, int, error) {
	start, _, err := getOrDefault(e.start, 0)
	if err != nil {
		return 0, 0, 0, err
	}

	step, _, err := getOrDefault(e.step, 1)
	if err != nil {
		return 0, 0, 0, err
	}

	var endDefault int
	if step > 0 {
		endDefault = length
	} else {
		endDefault = -1
	}

	end, isEndDef, err := getOrDefault(e.end, endDefault)
	if err != nil {
		return 0, 0, 0, err
	}

	if !unbounded {
		if start < 0 {
			start = length + start
		}
		if !isEndDef && end < 0 {
			end = length + end
		}
	}
	return start, end, step, err
}

func getLoopCondition(lim, step int) func(int) bool {
	if step < 0 {
		return func(i int) bool {
			return i > lim
		}
	}
	return func(i int) bool {
		return i < lim
	}
}

func getOrDefault(s string, def int) (int, bool, error) {
	if s == "" {
		return def, true, nil
	}
	res, err := strconv.Atoi(s)
	return res, false, err
}
