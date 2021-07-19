package chapter1

func createStatementData(invoice invoice, plays map[string]play) *StatementData {
	playFor := func(aPerformance performance) play {
		return plays[aPerformance.PlayID]
	}
	enrichPerformance := func(aPerformance []performance) []performance {
		result := make([]performance, 0, len(aPerformance))
		for _, p := range aPerformance {
			calculator := NewCalculator(p, playFor(p))
			p.Play = calculator.Play()
			p.Amount = calculator.Amount()
			p.VolumeCredits = calculator.VolumeCredits()
			result = append(result, p)
		}

		return result
	}

	statementData := new(StatementData)
	statementData.Customer = invoice.Customer
	statementData.Performances = enrichPerformance(invoice.Performances)
	statementData.TotalAmount = totalAmount(*statementData)
	statementData.TotalVolumeCredits = totalVolumeCredits(*statementData)
	return statementData
}

type Calculator interface {
	Play() play
	Amount() float64
	VolumeCredits() int
}

func NewCalculator(aPerformance performance, aPlay play) Calculator {
	switch aPlay.Type {
	case "tragedy":
		return Tragedy{base{
			perf: aPerformance,
			play: aPlay,
		}}
	case "comedy":
		return Comedy{base{
			perf: aPerformance,
			play: aPlay,
		}}
	default:
		panic("subclass responsibility")
	}
}

type base struct {
	perf performance
	play play
}

func (b base) Play() play {
	return b.play
}

func (b base) VolumeCredits() int {
	return func() int {
		if b.perf.Audience-30 > 0 {
			return b.perf.Audience - 30
		}
		return 0
	}()
}

type Tragedy struct {
	base
}

func (t Tragedy) Amount() float64 {
	var result float64 = 40000
	if t.perf.Audience > 30 {
		result += float64(1000 * (t.perf.Audience - 30))
	}
	return result
}

type Comedy struct {
	base
}

func (c Comedy) Amount() float64 {
	var result float64 = 30000
	if c.perf.Audience > 20 {
		result += float64(10000 + 500*(c.perf.Audience-20))
	}
	result += float64(300 * c.perf.Audience)
	return result
}
func (c Comedy) VolumeCredits() int {
	return c.base.VolumeCredits() + c.perf.Audience/5
}

func totalVolumeCredits(data StatementData) int {
	var result int
	for _, perf := range data.Performances {
		result += perf.VolumeCredits
	}
	return result
}
func totalAmount(data StatementData) float64 {
	var result float64
	for _, perf := range data.Performances {
		result += perf.Amount
	}
	return result
}
