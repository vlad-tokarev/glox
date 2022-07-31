package scanner

func (s *Scanner) advance() rune {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

func (s *Scanner) peekNext() rune {
	if int(s.current+1) >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= int64(len(s.source))
}
