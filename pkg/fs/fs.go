package fs

import (
	"bufio"
	"fmt"
	"io"
)

func IterateOverFile(file io.Reader) func(func(string, error) bool) {
	scanner := bufio.NewScanner(file)
	return func(yield func(string, error) bool) {
		for scanner.Scan() {
			line := scanner.Text()
			if !yield(line, nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			yield("", fmt.Errorf("error while iteration over file %w", err))
		}
	}
}

