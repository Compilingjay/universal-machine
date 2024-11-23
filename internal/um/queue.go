package um

import "fmt"

type Queue []uint32

func (s *Queue) push(v uint32) {
	*s = append(*s, v)
}

func (s *Queue) pop() (uint32, error) {
	if len(*s) == 0 {
		return 0xffffffff, fmt.Errorf("no value in stack")
	}

	v := (*s)[0]
	*s = (*s)[1:]
	return v, nil
}
