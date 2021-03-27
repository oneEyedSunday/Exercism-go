package strain

type Ints []int
type Lists [][]int
type Strings []string

func (input Ints) Keep(pred func(int) bool) Ints {
	if input == nil {
		return nil
	}
	vals := make(Ints, 0)
	for _, item := range input {
		if pred(item) {
			vals = append(vals, item)
		}
	}

	return vals
}

func (input Lists) Keep(pred func([]int) bool) Lists {
	if input == nil {
		return nil
	}
	vals := make(Lists, 0)
	for _, item := range input {
		if pred(item) {
			vals = append(vals, item)
		}
	}

	return vals
}

func (input Strings) Keep (pred func(string) bool) Strings {
	if input == nil {
		return nil
	}
	vals := make(Strings, 0)
	for _, item := range input {
		if pred(item) {
			vals = append(vals, item)
		}
	}

	return vals
}

func (input Ints) Discard(pred func(int) bool) Ints {
	if input == nil {
		return nil
	}
	vals := make(Ints, 0)
	for _, item := range input {
		if !pred(item) {
			vals = append(vals, item)
		}
	}

	return vals
}

//    (Ints) Keep(func(int) bool) Ints
//    (Ints) Discard(func(int) bool) Ints
//    (Lists) Keep(func([]int) bool) Lists
//    (Strings) Keep(func(string) bool) Strings

