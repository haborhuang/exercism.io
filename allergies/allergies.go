package allergies

func AllergicTo(score uint, allergen string) bool {
	return scores[allergen]&score > 0
}

func Allergies(score uint) []string {
	res := make([]string, 0, len(allergens))
	for i := 0; i < len(allergens) && score > 0; i++ {
		if score&1 > 0 {
			res = append(res, allergens[i])
		}
		score >>= 1
	}

	return res
}

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

var scores map[string]uint

func init() {
	scores = make(map[string]uint, len(allergens))
	for i, a := range allergens {
		scores[a] = 1 << uint(i)
	}
}
