package leetcode

// BrowserHistory ...
type BrowserHistory struct {
	stack []string
	cur   int
}

// Constructor ...
func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		stack: []string{homepage},
		cur:   0,
	}
}

// Visit ...
func (h *BrowserHistory) Visit(url string) {
	h.stack = append(h.stack[:h.cur+1], url)
	h.cur++
}

// Back ...
func (h *BrowserHistory) Back(steps int) string {
	h.cur -= steps
	if h.cur < 0 {
		h.cur = 0
	}
	return h.stack[h.cur]
}

// Forward ...
func (h *BrowserHistory) Forward(steps int) string {
	h.cur += steps
	if h.cur > len(h.stack)-1 {
		h.cur = len(h.stack) - 1
	}
	return h.stack[h.cur]
}
