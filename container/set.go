package container

type Set interface {
	Size() int
	Contains(e int) bool
	Add(e int) bool
	Remove(e int) bool
}

type HashSet map[int]bool

func (h HashSet) Size() int {
	return len(h)
}

func (h HashSet) Contains(e int) bool {
	_, ok := (h)[e]
	return ok
}

func (h HashSet) Add(e int) bool {
	if _, ok := h[e]; ok {
		return false
	}
	h[e] = true
	return true
}

func (h HashSet) Remove(e int) bool {
	delete(h, e)
	return true
}

func NewHashSet() Set {
	return &HashSet{}
}
