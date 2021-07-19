package chapter1

import (
	"fmt"
	"log"
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
	Customer     string
	Performances []performance
}

func statement(invoice invoice, plays map[string]play) string {
	playFor := func(aPerformance performance) play {
		return plays[aPerformance.PlayID]
	}
	enrichPerformance := func(aPerformance []performance) []performance {
		result := make([]performance, 0, len(aPerformance))
		for _, p := range aPerformance {
			p.Play = playFor(p)
			p.Amount = amountFor(p)
			result = append(result, p)
		}

		return result
	}
	statementData := new(StatementData)
	statementData.Customer = invoice.Customer
	statementData.Performances = enrichPerformance(invoice.Performances)
	return renderPlainText(statementData, plays)
}

func renderPlainText(data *StatementData, plays map[string]play) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf("Statement for %s\n", data.Customer))

	volumeCreditFor := func(perf performance) int {
		result := 0
		// add volume credits
		result += func() int {
			if perf.Audience-30 > 0 {
				return perf.Audience - 30
			}
			return 0
		}()
		// add extra credits for every ten comedy attendees
		if perf.PlayID == "comedy" {
			result += perf.Audience / 5
		}
		return result
	}

	totalVolumeCredits := func() int {
		var result int
		for _, perf := range data.Performances {
			result += volumeCreditFor(perf)
		}
		return result
	}
	totalAmount := func() float64 {
		var result float64
		for _, perf := range data.Performances {
			result += perf.Amount
		}
		return result
	}
	for _, perf := range data.Performances {
		// print line for this order
		strBuilder.WriteString(fmt.Sprintf("  %s:$%s (%d)\n", perf.Play.Name, usd(perf.Amount), perf.Audience))
	}

	strBuilder.WriteString(fmt.Sprintf("Amount owned is %s\n", usd(totalAmount())))
	strBuilder.WriteString(fmt.Sprintf("You earned %d credits", totalVolumeCredits()))
	return strBuilder.String()
}

func usd(aNumber float64) string {
	return fmt.Sprintf("%0.2f", aNumber/100)
}

func amountFor(perf performance) float64 {
	var result float64
	switch perf.Play.Type {
	case "tragedy":
		result = 40000
		if perf.Audience > 30 {
			result += float64(1000 * (perf.Audience - 30))
		}
	case "comedy":
		result = 30000
		if perf.Audience > 20 {
			result += float64(10000 + 500*(perf.Audience-20))
		}
		result += float64(300 * perf.Audience)
	default:
		log.Panicf("unknown type %s", perf.Play.Type)
	}
	return result
}
