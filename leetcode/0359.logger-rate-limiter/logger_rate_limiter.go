package leetcode

// Logger ...
type Logger struct {
	data map[string]int
}

// Constructor initialize your data structure here
func Constructor() Logger {
	return Logger{data: map[string]int{}}
}

// ShouldPrintMessage returns true if the message should be printed in the given timestamp,
// otherwise returns false.
// If this method returns false, the message will not be printed.
// The timestamp is in seconds granularity. */
func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, ok := l.data[message]; ok {
		if timestamp-l.data[message] < 10 {
			return false
		}
	}
	l.data[message] = timestamp
	return true
}
