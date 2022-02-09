package standalone

import (
	"github.com/romitou/skriptmc-parser/structures"
)

func BuildParsers() *structures.Task {
	return &structures.Task{
		Name: "rebuildParsers",
		Exec: func(ctx *structures.ParserContext) {
			// TODO: make the task
		},
	}
}
