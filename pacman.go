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
		return p.applyStart(e, false)
	case StartEndType:
		return p.applyStartEnd(e, false)
	case StartEndStepType:
		return p.applyStartEndStep(e, false)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

// ApplyUnbounded applies the effect to the slice as if it is unbounded
func (p Pacman) ApplyUnbounded(e Effect) (interface{}, error) {
	switch e.effectType {
	case StartType:
		return p.applyStart(e, true)
	case StartEndType:
		return p.applyStartEnd(e, true)
	case StartEndStepType:
		return p.applyStartEndStep(e, true)
	default:
		return nil, fmt.Errorf("Effect not valid: unknown type \"%s\"", reflect.TypeOf(e.effectType))
	}

}

func (p Pacman) applyStart(e Effect, unbounded bool) (interface{}, error) {
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

func (p Pacman) applyStartEnd(e Effect, unbounded bool) (interface{}, error) {
	switch reflect.TypeOf(p.slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(p.slice)
		len := s.Len()
		start, _, err := getOrDefault(e.start, 0)
		if err != nil {
			return nil, err
		}

		end, isEndDef, err := getOrDefault(e.end, len)
		if err != nil {
			return nil, err
		}

		if !unbounded {
			if start < 0 {
				start = len + start
			}
			if !isEndDef && end < 0 {
				end = len + end
			}
		}

		res := applySlice(p.slice, start, end, 1, unbounded)

		return res, nil
	default:
		return nil, fmt.Errorf("Invalid slice: \"%s\" is not a slice", reflect.TypeOf(p.slice))
	}
}

func (p Pacman) applyStartEndStep(e Effect, unbounded bool) (interface{}, error) {
	switch reflect.TypeOf(p.slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(p.slice)
		l := s.Len()
		start, _, err := getOrDefault(e.start, 0)
		if err != nil {
			return nil, err
		}

		step, _, err := getOrDefault(e.step, 1)
		if err != nil {
			return nil, err
		}

		var endDefault int
		if step > 0 {
			endDefault = l
		} else {
			endDefault = -1
		}

		end, isEndDef, err := getOrDefault(e.end, endDefault)
		if err != nil {
			return nil, err
		}

		if !unbounded {
			if start < 0 {
				start = l + start
			}
			if !isEndDef && end < 0 {
				end = l + end
			}
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
