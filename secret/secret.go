package secret

const reverseCode = 0x10

var secCodes = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) []string {
	res := make([]string, 0, len(secCodes))

	reversed := code&reverseCode == reverseCode

	for i := 0; i < len(secCodes); i++ {
		j := i
		if reversed {
			j = len(secCodes) - 1 - i
		}

		sc := uint(1 << uint(j))
		if code < sc {
			break
		}

		if sc&code == sc {
			res = append(res, secCodes[j])
		}
	}

	return res
}
