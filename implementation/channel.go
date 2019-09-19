package implementation

// Channel is a set of operations with channel
type Channel struct {
	data chan T
}

// Drop drops first n elements from channel c and returns a new channel with the rest.
// It returns channel do be unblocking. If you want array instead, wrap result into TakeAll.
func (c Channel) Drop(n int) chan T {
	result := make(chan T, 1)
	go func() {
		i := 0
		for el := range c.data {
			if i >= n {
				result <- el
			}
			i++
		}
		close(result)
	}()
	return result
}

// Each calls f for every element in the channel
func (c Channel) Each(f func(el T)) {
	for el := range c.data {
		f(el)
	}
}

// Filter returns a new channel with elements from input channel
// for which f returns true
func (c Channel) Filter(f func(el T) bool) chan T {
	result := make(chan T, 1)
	go func() {
		for el := range c.data {
			if f(el) {
				result <- el
			}
		}
	}()
	return result
}

// Take takes first n elements from channel c.
func (c Channel) Take(n int) []T {
	result := make([]T, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, <-c.data)
	}
	return result
}

// ToSlice returns slice with all elements from channel.
func (c Channel) ToSlice() []T {
	result := make([]T, 0)
	for val := range c.data {
		result = append(result, val)
	}
	return result
}
