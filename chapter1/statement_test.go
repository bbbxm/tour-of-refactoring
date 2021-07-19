package chapter1

import "testing"

func Test_statement(t *testing.T) {
	type args struct {
		invoice invoice
		plays   map[string]play
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1",
			args{
				invoice: Invoices[0],
				plays:   Plays,
			},
			`Statement for BigCo
  Hamlet:$650.00 (55)
  As You Like It:$580.00 (35)
  Othello:$500.00 (40)
Amount owned is 1730.00
You earned 47 credits`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := statement(tt.args.invoice, tt.args.plays); got != tt.want {
				t.Errorf("statement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_statementHtml(t *testing.T) {
	type args struct {
		invoice invoice
		plays   map[string]play
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1",
			args{
				invoice: Invoices[0],
				plays:   Plays,
			},
			`Statement for BigCo
  Hamlet:$650.00 (55)
  As You Like It:$580.00 (35)
  Othello:$500.00 (40)
Amount owned is 1730.00
You earned 40 credits`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := statementHtml(tt.args.invoice, tt.args.plays); got != tt.want {
				t.Errorf("statement() = %v, want %v", got, tt.want)
			}
		})
	}
}
