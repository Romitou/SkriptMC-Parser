package structures

type Task struct {
	Name string
	Cron string
	Exec func(ctx *ParserContext)
}
