func isAlphaNumeric(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		leftChar := rune(s[left])
		rightChar := rune(s[right])

		// 1. Skip invalid characters from the left
		if !isAlphaNumeric(leftChar) {
			left++
			continue
		}

		// 2. Skip invalid characters from the right
		if !isAlphaNumeric(rightChar) {
			right--
			continue
		}

		// 3. Normalize case to lowercase
		leftChar = unicode.ToLower(leftChar)
		rightChar = unicode.ToLower(rightChar)

		// 4. Compare characters
		if leftChar != rightChar {
			return false
		}

		// 5. Advance pointers inward
		left++
		right--
	}

	return true
}