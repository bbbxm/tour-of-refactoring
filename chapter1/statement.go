package chapter1

import (
	"fmt"
	"strings"
)

var (
	Plays = map[string]play{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
		"othello": {Name: "Othello", Type: "tragedy"},
	}
	Invoices = []invoice{
		{Customer: "BigCo", Performances: []performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}},
	}
)

type StatementData struct {
	Customer           string
	Performances       []performance
	TotalAmount        float64
	TotalVolumeCredits int
}

func statement(invoice invoice, plays map[string]play) string {
	return renderPlainText(createStatementData(invoice, plays))
}

func renderPlainText(data *StatementData) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf("Statement for %s\n", data.Customer))

	for _, perf := range data.Performances {
		// print line for this order
		strBuilder.WriteString(fmt.Sprintf("  %s:$%s (%d)\n", perf.Play.Name, usd(perf.Amount), perf.Audience))
	}

	strBuilder.WriteString(fmt.Sprintf("Amount owned is %s\n", usd(data.TotalAmount)))
	strBuilder.WriteString(fmt.Sprintf("You earned %d credits", data.TotalVolumeCredits))
	return strBuilder.String()
}

func statementHtml(invoice invoice, plays map[string]play) string {
	return renderHtml(createStatementData(invoice, plays))
}

func renderHtml(data *StatementData) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf("<h1>Statement for %s </h1>\n", data.Customer))
	strBuilder.WriteString("<table>\n")
	strBuilder.WriteString("<tr><th>play</th><th>seats</th><th>cost</th></tr>")
	for _, perf := range data.Performances {
		// print line for this order
		strBuilder.WriteString(fmt.Sprintf("  <tr><td>%s</td><td>$%s</td><td>%d</td></tr>\n", perf.Play.Name, usd(perf.Amount), perf.Audience))
	}
	strBuilder.WriteString("</table>\n")
	strBuilder.WriteString(fmt.Sprintf("<p>Amount owned is <em>%s</em></p>\n", usd(data.TotalAmount)))
	strBuilder.WriteString(fmt.Sprintf("<p>You earned <em>%d</em> credits</p>", data.TotalVolumeCredits))
	return strBuilder.String()
}

func usd(aNumber float64) string {
	return fmt.Sprintf("%0.2f", aNumber/100)
}
