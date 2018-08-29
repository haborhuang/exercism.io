package perfect

import (
	"errors"
)

// Classification is result classification
type Classification int8

const (
	// ClassificationNull is the zero value
	ClassificationNull Classification = iota
	ClassificationPerfect
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("Support only positive")

func Classify(num int64) (c Classification, err error) {
	if num <= 0 {
		err = ErrOnlyPositive
		return
	}

	if num == 1 {
		return ClassificationDeficient, nil
	}

	var sum int64 = 1
	for i := int64(2); ; i++ {
		j := num / i
		if i > j {
			break
		}

		if num%i == 0 {
			if i != 1 && j != i {
				sum += j
			}
			sum += i
		}
	}

	c = ClassificationPerfect
	if sum < num {
		c = ClassificationDeficient
	} else if sum > num {
		c = ClassificationAbundant
	}

	return
}
