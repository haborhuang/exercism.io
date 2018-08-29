package brackets

func Bracket(input string) (bool, error) {
	s := newStack(len(input) / 2)
	for _, c := range input {
		switch c {
		case '[', '{', '(':
			s.push(bracketType(c))
		case ']', '}', ')':
			if !s.pop(getBracketType(c)) {
				return false, nil
			}
		}
	}

	return s.isEmpty(), nil
}

func getBracketType(right rune) bracketType {
	switch right {
	case ']':
		return bracketTypeBracket
	case '}':
		return bracketTypeBrace
	case ')':
		return bracketTypeParenthese
	}
	panic("Unexpected error")
}

type bracketType rune

const (
	bracketTypeBracket    = '['
	bracketTypeBrace      = '{'
	bracketTypeParenthese = '('
)

type stack struct {
	lefts []bracketType
}

func newStack(cap int) *stack {
	return &stack{
		lefts: make([]bracketType, 0, cap),
	}
}

func (s *stack) push(t bracketType) {
	s.lefts = append(s.lefts, t)
}

func (s *stack) pop(t bracketType) bool {
	if len(s.lefts) < 1 || s.lefts[len(s.lefts)-1] != t {
		return false
	}

	s.lefts = s.lefts[:len(s.lefts)-1]
	return true

}

func (s *stack) isEmpty() bool {
	return len(s.lefts) == 0
}
