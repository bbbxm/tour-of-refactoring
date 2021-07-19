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

func statement(invoice invoice, plays map[string]play) string {
	strBuilder := strings.Builder{}
	totalAmount, volumeCredits := float64(0), 0
	strBuilder.WriteString(fmt.Sprintf("Statement for %s\n", invoice.Customer))
	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		thisAmount := float64(0)
		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += float64(1000 * (perf.Audience - 30))
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += float64(10000 + 500*(perf.Audience-20))
			}
			thisAmount += float64(300 * perf.Audience)
		default:
			log.Panicf("unknown type %s", play.Type)
		}
		// add volume credits
		volumeCredits += func() int {
			if perf.Audience-30 > 0 {
				return perf.Audience - 30
			}
			return 0
		}()
		// add extra credits for every ten comedy attendees
		if perf.PlayID == "comedy" {
			volumeCredits += perf.Audience / 5
		}
		// print line for this order
		strBuilder.WriteString(fmt.Sprintf("  %s:$%0.2f (%d)\n", play.Name, thisAmount/100, perf.Audience))
		totalAmount += thisAmount
	}
	strBuilder.WriteString(fmt.Sprintf("Amount owned is %0.2f\n", totalAmount/100))
	strBuilder.WriteString(fmt.Sprintf("You earned %d credits", volumeCredits))
	return strBuilder.String()
}
