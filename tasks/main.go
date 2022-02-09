package tasks

import (
	"github.com/mileusna/crontab"
	"github.com/romitou/skriptmc-parser/structures"
	"github.com/romitou/skriptmc-parser/tasks/scheduled"
	"github.com/romitou/skriptmc-parser/tasks/standalone"
	"log"
)

func InitTasks(ctx *structures.ParserContext) {
	InitCrontab(ctx)
	RegisterTasks(ctx,
		scheduled.RebuildParsers(),
		scheduled.CheckParsers(),
		standalone.BuildParsers(),
	)
}

func InitCrontab(ctx *structures.ParserContext) {
	ctx.Crontab = crontab.New()
}

func RegisterTasks(ctx *structures.ParserContext, tasks ...*structures.Task) {
	for i := range tasks {
		RegisterTask(ctx, tasks[i])
	}
}

func RegisterTask(ctx *structures.ParserContext, task *structures.Task) {
	if task.Cron == "" {
		return
	}
	err := ctx.Crontab.AddJob(task.Cron, task.Exec, ctx)
	if err != nil {
		log.Fatal("[Skript-MC] An error occurred when registering "+task.Name+" task: ", err)
	}
	ctx.Tasks[task.Name] = task
	log.Println("[Skript-MC] Task " + task.Name + " successfully registered.")
}
