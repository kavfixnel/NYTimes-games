package spellingBee

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	requiredRunes = map[rune]struct{}{'f': {}}
	extraRunes    = map[rune]struct{}{
		't': {},
		'p': {},
		'a': {},
		'y': {},
		'e': {},
		'c': {},
	}
	allRunes = map[rune]struct{}{
		'f': {},
		't': {},
		'p': {},
		'a': {},
		'y': {},
		'e': {},
		'c': {},
	}

	validWords   = []string{"Face", "catface", "caffeate", "feat"}
	invalidWords = []string{"Faced", "grommet", "clicker"}
)

func TestSanitizeRunes(t *testing.T) {
	t.Run("can correctly sanitize a rune set", func(t *testing.T) {
		assert.True(t, reflect.DeepEqual(
			sanitizeRunes(
				map[rune]struct{}{'H': {}, 'e': {}, 'l': {}, 'o': {}},
				defaultRuneSanitizer,
			),
			map[rune]struct{}{'h': {}, 'e': {}, 'l': {}, 'o': {}},
		))
	})
}

func TestWordToRunes(t *testing.T) {
	t.Run("can correctly convert a word to it's runes", func(t *testing.T) {
		assert.True(t, reflect.DeepEqual(
			wordToRunes("hello"),
			map[rune]struct{}{'h': {}, 'e': {}, 'l': {}, 'o': {}},
		))
	})
}

func TestIsValidWord(t *testing.T) {
	t.Run("can correctly identify valid words", func(t *testing.T) {
		for _, validWord := range validWords {
			assert.True(t, isValidWord(validWord, requiredRunes, extraRunes))
		}
	})

	t.Run("can correctly reject invalid words", func(t *testing.T) {
		for _, invalidWord := range invalidWords {
			assert.False(t, isValidWord(invalidWord, requiredRunes, extraRunes))
		}
	})
}

func TestFilterWordList(t *testing.T) {
	t.Run("returns empty list when passed an empty list", func(t *testing.T) {
		assert.Len(t, filterWordList(requiredRunes, extraRunes, []string{}), 0)
	})

	t.Run("correctly filters a list of words", func(t *testing.T) {
		assert.ElementsMatch(t,
			validWords,
			filterWordList(requiredRunes, extraRunes,
				append(append([]string{}, validWords[:]...), invalidWords[:]...),
			),
		)
	})

	t.Run("returns empty list if no words are found", func(t *testing.T) {
		assert.ElementsMatch(t,
			[]string{},
			filterWordList(requiredRunes, extraRunes, invalidWords),
		)
	})
}

func TestWordScore(t *testing.T) {
	t.Run("word score is 0 on a words shorter than 4", func(t *testing.T) {
		shortWords := []string{"", "a", "aa", "aaa"}

		for _, shortWord := range shortWords {
			assert.Equal(t, 0, wordScore(shortWord, allRunes))
		}
	})

	t.Run("word score is correct for words without the pangram bonus", func(t *testing.T) {
		shortWords := map[string]int{
			"face":      1, // 4 letter words are worth 1 point
			"facet":     5, // 5+ letter words are worth 1 point per rune
			"effect":    6,
			"catface":   7,
			"caffeate":  8,
			"affectate": 9,
		}

		for shortWord, expectedScore := range shortWords {
			assert.Equal(t, expectedScore, wordScore(shortWord, allRunes))
		}
	})

	t.Run("word score is correct for words with the pangram bonus", func(t *testing.T) {
		shortWords := map[string]int{
			"ftpayec":    14,
			"ftpayecpay": 17,
		}

		for shortWord, expectedScore := range shortWords {
			assert.Equal(t, expectedScore, wordScore(shortWord, allRunes))
		}
	})
}
