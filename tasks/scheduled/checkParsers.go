package scheduled

import "github.com/romitou/skriptmc-parser/structures"

func CheckParsers() *structures.Task {
	return &structures.Task{
		Name: "rebuildParsers",
		Cron: "* * * * *",
		Exec: func(ctx *structures.ParserContext) {
			// TODO: make the task
		},
	}
}
