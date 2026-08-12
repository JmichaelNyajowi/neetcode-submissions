func isValid(s string) bool {
	// Initialize doubly-linked list from standard library
	stack := list.New()
	closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, c := range s {
		if open, exists := closeToOpen[c]; exists {
			// Equivalent to: !stack.Empty()
			if stack.Len() != 0 {
				// Remove front element (top of stack) and retrieve value
				top := stack.Remove(stack.Front())
				if top.(rune) != open {
					return false
				}
			} else {
				return false
			}
		} else {
			// Insert at front (top of stack)
			stack.PushFront(c)
		}
	}

	// Equivalent to: return stack.Empty()
	return stack.Len() == 0
}