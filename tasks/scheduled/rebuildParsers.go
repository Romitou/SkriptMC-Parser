package scheduled

import (
	"github.com/romitou/skriptmc-parser/structures"
)

func RebuildParsers() *structures.Task {
	return &structures.Task{
		Name: "rebuildParsers",
		Cron: "0 0 * * *",
		Exec: func(ctx *structures.ParserContext) {
			// TODO: make the task
		},
	}
}
