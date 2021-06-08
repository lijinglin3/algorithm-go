package leetcode

// AnimalShelf ...
type AnimalShelf struct {
	cat, dog [][]int
}

// Constructor ...
func Constructor() AnimalShelf {
	return AnimalShelf{}
}

// Enqueue ...
func (s *AnimalShelf) Enqueue(animal []int) {
	if animal[1] == 0 {
		s.cat = append(s.cat, animal)
		return
	}
	s.dog = append(s.dog, animal)
}

// DequeueAny ...
func (s *AnimalShelf) DequeueAny() []int {
	ret := []int{-1, -1}
	switch {
	case len(s.cat) != 0 && len(s.dog) != 0 && s.cat[0][0] > s.dog[0][0]:
		ret := s.dog[0]
		s.dog = s.dog[1:]
		return ret
	case len(s.cat) != 0 && len(s.dog) != 0 && s.cat[0][0] < s.dog[0][0]:
		ret := s.cat[0]
		s.cat = s.cat[1:]
		return ret
	case len(s.dog) != 0 && len(s.cat) == 0:
		ret := s.dog[0]
		s.dog = s.dog[1:]
		return ret
	case len(s.dog) == 0 && len(s.cat) != 0:
		ret = s.cat[0]
		s.cat = s.cat[1:]
	}
	return ret
}

// DequeueDog ...
func (s *AnimalShelf) DequeueDog() []int {
	ret := []int{-1, -1}
	if len(s.dog) != 0 {
		ret = s.dog[0]
		s.dog = s.dog[1:]
	}
	return ret
}

// DequeueCat ...
func (s *AnimalShelf) DequeueCat() []int {
	ret := []int{-1, -1}
	if len(s.cat) != 0 {
		ret = s.cat[0]
		s.cat = s.cat[1:]
	}
	return ret
}
