package sublist

func Sublist(l1, l2 []int) Relation {
	short, long := l1, l2
	if len(l1) > len(l2) {
		short, long = l2, l1
	}

	switch isSubList(short, long) {
	case isEqual:
		return RelEqual
	case isUnequal:
		return RelUnequal
	default:
		if len(l1) > len(l2) {
			return RelSuperList
		}

		return RelSubList
	}
}

type resType int8

const (
	isSublist resType = iota
	isEqual
	isUnequal
)

type Relation string

const (
	RelEqual     Relation = "equal"
	RelUnequal   Relation = "unequal"
	RelSuperList Relation = "superlist"
	RelSubList   Relation = "sublist"
)

func isSubList(short, long []int) resType {
	if len(long) == 0 {
		return isEqual
	} else if len(short) == 0 {
		return isSublist
	}

	for i := 0; i <= len(long)-len(short); i++ {
		j := 0
		k := i
		for range short {
			if long[k] != short[j] {
				break
			}
			j++
			k++
		}

		if j == len(short) {
			if k == len(long) && len(short) == len(long) {
				return isEqual
			}
			return isSublist
		}
	}

	return isUnequal
}
