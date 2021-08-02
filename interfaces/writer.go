package interfaces

import "github.com/masu-mi/test_gomod_replace_dep_a"

type Writer interface {
	Write(payload string) (num test_gomod_replace_dep_a.OriginalInt)
}
