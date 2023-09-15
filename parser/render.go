package parser

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
)

func RenderOutput(res *response) {
	t := table.NewWriter()
	t.AppendRow(table.Row{"minute", res.minutes})
	t.AppendRow(table.Row{"hour", res.hour})
	t.AppendRow(table.Row{"day of month", res.dayOfMonth})
	t.AppendRow(table.Row{"month", res.month})
	t.AppendRow(table.Row{"day of week", res.dayOfWeek})
	t.AppendRow(table.Row{"command", res.command})
	fmt.Println(t.Render())
}
