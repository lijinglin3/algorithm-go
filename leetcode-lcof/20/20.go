package leetcode

const (
	stateInitial int = iota
	stateIntSign
	stateInteger
	statePoint
	statePointWithoutInt
	stateFraction
	stateExp
	stateExpSign
	stateExpNumber
	stateEnd
)

const (
	charNumber int = iota
	charExp
	charPoint
	charSign
	charSpace
	charIllegal
)

var (
	transfer = map[int]map[int]int{
		stateInitial: {
			charSpace:  stateInitial,
			charNumber: stateInteger,
			charPoint:  statePointWithoutInt,
			charSign:   stateIntSign,
		},
		stateIntSign: {
			charNumber: stateInteger,
			charPoint:  statePointWithoutInt,
		},
		stateInteger: {
			charNumber: stateInteger,
			charExp:    stateExp,
			charPoint:  statePoint,
			charSpace:  stateEnd,
		},
		statePoint: {
			charNumber: stateFraction,
			charExp:    stateExp,
			charSpace:  stateEnd,
		},
		statePointWithoutInt: {
			charNumber: stateFraction,
		},
		stateFraction: {
			charNumber: stateFraction,
			charExp:    stateExp,
			charSpace:  stateEnd,
		},
		stateExp: {
			charNumber: stateExpNumber,
			charSign:   stateExpSign,
		},
		stateExpSign: {
			charNumber: stateExpNumber,
		},
		stateExpNumber: {
			charNumber: stateExpNumber,
			charSpace:  stateEnd,
		},
		stateEnd: {
			charSpace: stateEnd,
		},
	}
)

func isNumber(s string) bool {
	state := stateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		}
		state = transfer[state][typ]
	}
	return state == stateInteger ||
		state == statePoint ||
		state == stateFraction ||
		state == stateExpNumber ||
		state == stateEnd
}

func toCharType(ch byte) int {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return charNumber
	case 'e', 'E':
		return charExp
	case '.':
		return charPoint
	case '+', '-':
		return charSign
	case ' ':
		return charSpace
	default:
		return charIllegal
	}
}
