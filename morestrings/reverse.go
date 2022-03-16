package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.

var service = Service{}

func init() {
	service.sameValueService = &SameValueService{}
}

func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	// fmt.Println(&service.sameValueService)
	service.sameValueService.Call("11")
	return string(r)
}

type Service struct {
	sameValueService interface {
		Call(string) string
	}
}
