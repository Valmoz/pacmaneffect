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

// Apply applies the effect to the slice
func (p Pacman) Apply(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applyStart(e)
	case StartEndType:
		return p.applyStartEnd(e)
	case StartEndStepType:
		return p.applyStartEndStep(e)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applyStart(e Effect) (interface{}, error) {
	start, err := strconv.Atoi(e.start)
	if err != nil {
		return nil, err
	}

	switch reflect.TypeOf(p.slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(p.slice)
		len := s.Len()
		if start >= 0 {
			return s.Index(start).Interface(), nil
		}
		return s.Index(len + start).Interface(), nil
	default:
		return nil, fmt.Errorf("Invalid slice: \"%s\" is not a slice", reflect.TypeOf(p.slice))
	}

}

func (p Pacman) applyStartEnd(e Effect) (interface{}, error) {
	switch reflect.TypeOf(p.slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(p.slice)
		len := s.Len()
		start, err := getOrDefault(e.start, 0)
		if err != nil {
			return nil, err
		}

		end, err := getOrDefault(e.end, len)
		if err != nil {
			return nil, err
		}

		if start < 0 {
			start = len + start
		}
		if end < 0 {
			end = len + end
		}

		start = max(0, start)
		end = min(len, end)

		res := applySlice(p.slice, start, end, 1)

		return res, nil
	default:
		return nil, fmt.Errorf("Invalid slice: \"%s\" is not a slice", reflect.TypeOf(p.slice))
	}
}

func (p Pacman) applyStartEndStep(e Effect) (interface{}, error) {
	switch reflect.TypeOf(p.slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(p.slice)
		len := s.Len()
		start, err := getOrDefault(e.start, 0)
		if err != nil {
			return nil, err
		}

		end, err := getOrDefault(e.end, len)
		if err != nil {
			return nil, err
		}

		step, err := getOrDefault(e.step, 1)
		if err != nil {
			return nil, err
		}

		if start < 0 {
			start = len + start
		}
		if end < 0 {
			end = len + end
		}

		start = max(0, start)
		end = min(len, end)

		res := applySlice(p.slice, start, end, step)

		return res, nil
	default:
		return nil, fmt.Errorf("Invalid slice: \"%s\" is not a slice", reflect.TypeOf(p.slice))
	}
}

func applySlice(slice interface{}, start, end, step int) interface{} {
	sliceType := reflect.TypeOf(slice)
	sliceValue := reflect.ValueOf(slice)

	res := reflect.MakeSlice(sliceType, 0, 0)
	if step == 0 { //|| (step > 0 && start >= end) || (step < 0 && start <= end) {
		return res.Interface()
	}

	cond := getLoopCondition(end, step)

	for i := start; cond(i); i = i + step {
		res = reflect.Append(res, sliceValue.Index(i))
	}

	return res.Interface()
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

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func getOrDefault(s string, def int) (int, error) {
	if s == "" {
		return def, nil
	}
	return strconv.Atoi(s)
}
