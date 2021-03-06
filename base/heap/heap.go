package heap

type Cmp func(i, j int) bool
type Value interface{}

type Heap interface {
	InitWithCmp(cmp Cmp)
	Get(i int) Value

	Push(x Value)
	Pop() Value
	Peek() Value

	Len() int
	IndexOf(x Value) int
	Fix(i int)
	Remove(i int) Value
	Update(i int, value Value)
}

func New() Heap {
	return NewWithCap(0)
}

func NewWithCap(cap int) Heap {
	return &heapImp{slice: make([]Value, 0, cap)}
}

func NewWithSlice(slice []Value) Heap {
	return &heapImp{slice: slice}
}

type heapImp struct {
	cmp   Cmp
	slice []Value
}

func (h *heapImp) InitWithCmp(cmp Cmp) {
	h.cmp = cmp
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}
func (h *heapImp) Get(i int) Value {
	return h.slice[i]
}

func (h *heapImp) Push(x Value) {
	h.slice = append(h.slice, x)
	h.up(h.Len() - 1)
}
func (h *heapImp) Pop() Value {
	n := h.Len() - 1
	h.swap(0, n)
	h.down(0, n)
	result := h.slice[n]
	h.slice = h.slice[:n]
	return result
}
func (h *heapImp) Peek() Value {
	return h.slice[0]
}

func (h *heapImp) Len() int {
	return len(h.slice)
}
func (h *heapImp) IndexOf(x Value) int {
	for i, v := range h.slice {
		if v == x {
			return i
		}
	}
	return -1
}
func (h *heapImp) Fix(i int) {
	if !h.down(i, h.Len()) {
		h.up(i)
	}
}
func (h *heapImp) Remove(i int) Value {
	n := h.Len() - 1
	if n != i {
		h.swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	result := h.slice[n]
	h.slice = h.slice[:n]
	return result
}
func (h *heapImp) Update(i int, value Value) {
	h.slice[i] = value
	h.Fix(i)
}

func (h *heapImp) swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func (h *heapImp) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.cmp(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *heapImp) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.cmp(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.cmp(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}
