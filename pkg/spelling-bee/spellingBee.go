package spellingBee

import (
	"sort"
	"unicode"

	"github.com/kavfixnel/words"
)

const (
	// minWordLength is the minimum length of the words.
	// This is defined by the rules of the game
	minWordLength = 4
)

var (
	// defaultRuneSanitizer is the default way the code will sanitizes runes
	defaultRuneSanitizer = unicode.ToLower
)

// sanitizeRunes sanitizes a map of map[rune]struct{} by applying a sanitizer function that maps runes to runes
// It returns the map of sanitized runes
func sanitizeRunes(runes map[rune]struct{}, sanitizer func(rune) rune) map[rune]struct{} {
	sanitizedRunes := make(map[rune]struct{}, len(runes))

	for r := range runes {
		sanitizedRunes[sanitizer(r)] = struct{}{}
	}

	return sanitizedRunes
}

// wordToRunes takes a string word and converts it into a rune set
// It return the run set of all the runes in the word
func wordToRunes(word string) map[rune]struct{} {
	wordRunes := make(map[rune]struct{}, len(word))
	for _, r := range word {
		wordRunes[r] = struct{}{}
	}
	return wordRunes
}

// isValidWord checks if a single word is valid based on the following parameters:
//   - The word must contain all required runes from requiredRunes
//   - The word must only contain runes from the superset of {requiredRunes..., extraRunes...}
//
// Note that with respect to runes, case does not matter. The function will take care of sanitization for you.
// It returns the validness of the word.
func isValidWord(word string, requiredRunes, extraRunes map[rune]struct{}) bool {
	if len(word) < minWordLength {
		return false
	}

	requiredRunes = sanitizeRunes(requiredRunes, defaultRuneSanitizer)
	extraRunes = sanitizeRunes(extraRunes, defaultRuneSanitizer)
	wordRunes := sanitizeRunes(wordToRunes(word), defaultRuneSanitizer)

	// Check that the word contains all required runes
	for r := range requiredRunes {
		if _, ok := wordRunes[r]; !ok {
			return false
		}
	}
	// Check that word only contains valid runes
	for r := range wordRunes {
		_, ok1 := requiredRunes[r]
		_, ok2 := extraRunes[r]
		if !(ok1 || ok2) {
			return false
		}
	}

	return true
}

// filterWordList takes a list of words and rune sets and filters the list down via the following parameters:
//   - The word must contain all required runes from requiredRunes
//   - The word must only contain runes from the superset of {requiredRunes..., extraRunes...}
//
// It returns the filtered list.
func filterWordList(requiredRunes, extraRunes map[rune]struct{}, wordList []string) []string {
	var validWords []string

	for _, word := range wordList {
		if isValidWord(word, requiredRunes, extraRunes) {
			validWords = append(validWords, word)
		}
	}

	return validWords
}

// wordScore takes a word and returns the score of a word based on the NY Times' "spelling-bee" game.
// https://www.nytimes.com/puzzles/spelling-bee
//   - Words shorter than 4 letters are woth 0 points
//   - 4-letter words are worth 1 point each.
//   - Longer words earn 1 point per letter.
//   - Each puzzle includes at least one “pangram” which uses every letter. These are worth 7 extra points!
//
// Note that this function does not check if the word is valid.
// It returns the score of the word.
func wordScore(word string, allRunes map[rune]struct{}) int {
	if len(word) < minWordLength {
		return 0
	}

	if len(word) == 4 {
		return 1
	}

	score := len(word)

	wordRunes := wordToRunes(word)
	// Award an an extra 7 points if the word contains all runes in the
	isPangram := true
	for r := range allRunes {
		if _, ok := wordRunes[r]; !ok {
			isPangram = false
			break
		}
	}
	if isPangram {
		score += 7
	}

	return score
}

// GetConstructableWordList takes a required and an extra set of runes, and returns a list of words that
// can be constructed with the folling parameters:
//   - The word must contain all required runes from requiredRunes
//   - The word must only contain runes from the superset of {requiredRunes..., extraRunes...}
//
// It returns a []string of
func GetConstructableWordList(requiredRunes, extraRunes map[rune]struct{}) ([]string, error) {
	wordList, err := words.NewWordList(nil)
	if err != nil {
		return []string{}, err
	}

	wordList = filterWordList(requiredRunes, extraRunes, wordList)

	allRunes := make(map[rune]struct{}, len(requiredRunes)+len(extraRunes))
	for _, runeSet := range []map[rune]struct{}{requiredRunes, extraRunes} {
		for k, v := range runeSet {
			allRunes[k] = v
		}
	}
	// Sort the words according each words wordScore and alphabetical if equal
	sort.Slice(wordList, func(i, j int) bool {
		iScore := wordScore(wordList[i], allRunes)
		jScore := wordScore(wordList[j], allRunes)
		if iScore == jScore {
			return wordList[i] < wordList[j]
		}
		return iScore < jScore
	})

	return wordList, nil
}
