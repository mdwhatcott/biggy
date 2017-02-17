package biggy

type FibonacciCounter struct {
	started bool

	a *Big
	b *Big
}

func NewFibonacciCounter() *FibonacciCounter {
	return &FibonacciCounter{
		a: NewBigFromSmall(0),
		b: NewBigFromSmall(1),
	}
}

func (this *FibonacciCounter) Next() string {
	if !this.started {
		this.started = true
		return "1"
	}
	this.a.Plus(this.b)
	this.a, this.b = this.b, this.a
	return this.b.String()
}

func (this *FibonacciCounter) Advance(n int) string {
	result := ""
	for x := 0; x < n; x++ {
		result = this.Next()
	}
	return result
}
