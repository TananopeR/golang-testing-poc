package morestrings

import (
	"fmt"
)

type SameValueService struct {
}

func (p *SameValueService) Call(s string) string {
	fmt.Println(s + " real")
	return s
}
