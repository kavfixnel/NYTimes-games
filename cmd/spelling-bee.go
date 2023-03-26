package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	spellingBee "github.com/kavfixnel/nytimes-games/pkg/spelling-bee"
)

var (
	requiredCharacters string
	extraCharacters    string
)

// preprocessArgs takes the string arguments for required and extra runes, checks if they are not
// empty strings and parses them into rune sets.
// It returns the two run sets and any errors encountered.
func preprocessArgs(required, extra string) (map[rune]struct{}, map[rune]struct{}, error) {
	if len(required) == 0 {
		return map[rune]struct{}{}, map[rune]struct{}{}, fmt.Errorf("argument required cannot be 0 characters")
	}
	if len(extra) == 0 {
		return map[rune]struct{}{}, map[rune]struct{}{}, fmt.Errorf("argument extra cannot be 0 characters")
	}

	requiredRuneSet := make(map[rune]struct{}, len(required))
	extraRuneSet := make(map[rune]struct{}, len(extra))
	for _, r := range required {
		requiredRuneSet[r] = struct{}{}
	}
	for _, r := range extra {
		extraRuneSet[r] = struct{}{}
	}

	return requiredRuneSet, extraRuneSet, nil
}

var spellingBeeCmd = &cobra.Command{
	Use:     "spelling-bee",
	Aliases: []string{"sp"},
	Short:   "Gives solutions to the spelling-bee game",
	RunE: func(cmd *cobra.Command, args []string) error {
		requiredRuneSet, extraRuneSet, err := preprocessArgs(requiredCharacters, extraCharacters)
		if err != nil {
			return err
		}

		wordList, err := spellingBee.GetConstructableWordList(requiredRuneSet, extraRuneSet)
		if err != nil {
			return err
		}

		for _, word := range wordList {
			fmt.Println(word)
		}

		return nil
	},
}

func init() {
	flags := spellingBeeCmd.Flags()

	// TODO: Revers engineer the NY-Times spelling bee api to get these parameters automatically
	flags.StringVarP(&requiredCharacters, "required", "r", "", `A concatinated string of required characters.
Usually only one character`)
	flags.StringVarP(&extraCharacters, "extra", "e", "", `A concatinated string of required letters.
Usually six letters`)

	cobra.MarkFlagRequired(flags, "required")
	cobra.MarkFlagRequired(flags, "extra")

	rootCmd.AddCommand(spellingBeeCmd)
}
