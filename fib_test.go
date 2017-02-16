package biggy

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestFibonacciFixture(t *testing.T) {
	gunit.Run(new(FibonacciFixture), t)
}

type FibonacciFixture struct {
	*gunit.Fixture

	fibonacci *FibonacciCounter
}

func (this *FibonacciFixture) Setup() {
	this.fibonacci = NewFibonacciCounter()
}

func (this *FibonacciFixture) TestSeries() {
	this.So(this.fibonacci.Next(), should.Equal, "1")
	this.So(this.fibonacci.Next(), should.Equal, "1")
	this.So(this.fibonacci.Next(), should.Equal, "2")
	this.So(this.fibonacci.Next(), should.Equal, "3")
	this.So(this.fibonacci.Next(), should.Equal, "5")
	this.So(this.fibonacci.Next(), should.Equal, "8")
	this.So(this.fibonacci.Next(), should.Equal, "13")
	this.So(this.fibonacci.Next(), should.Equal, "21")
	this.So(this.fibonacci.Next(), should.Equal, "34")
	this.So(this.fibonacci.Next(), should.Equal, "55")
}

func (this *FibonacciFixture) TestAdvance() {
	this.So(this.fibonacci.Advance(100), should.Equal, "354,224,848,179,261,915,075")
	this.So(this.fibonacci.Advance(900), should.Equal, "43,466,557,686,937,456,435,688,527,675,040,625,802,564,660,517,371,780,402,481,729,089,536,555,417,949,051,890,403,879,840,079,255,169,295,922,593,080,322,634,775,209,689,623,239,873,322,471,161,642,996,440,906,533,187,938,298,969,649,928,516,003,704,476,137,795,166,849,228,875")
}
