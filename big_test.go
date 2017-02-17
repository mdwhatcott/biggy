package biggy

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"strconv"
)

func TestBigFixture(t *testing.T) { gunit.Run(new(BigFixture), t) }

type BigFixture struct{ *gunit.Fixture }

func (this *BigFixture) TestParseAndDisplay() {
	this.So(parseAndDisplay("0"), should.Equal, "0")
	this.So(parseAndDisplay("12"), should.Equal, "12")
	this.So(parseAndDisplay("123"), should.Equal, "123")
	this.So(parseAndDisplay("1234"), should.Equal, "1,234")
	this.So(parseAndDisplay("12345"), should.Equal, "12,345")
	this.So(parseAndDisplay("123456"), should.Equal, "123,456")
	this.So(parseAndDisplay("1234567"), should.Equal, "1,234,567")
	this.So(parseAndDisplay("12345678"), should.Equal, "12,345,678")
	this.So(parseAndDisplay("123456789"), should.Equal, "123,456,789")
}

func parseAndDisplay(raw string) string {
	n, _ := strconv.Atoi(raw)
	return NewBigFromSmall(n).String()
}

func (this *BigFixture) TestAdd() {
	this.So(add("1", "1"), should.Equal, "2")
	this.So(add("2", "2"), should.Equal, "4")
	this.So(add("123", "456"), should.Equal, "579")
	this.So(add("998", "1"), should.Equal, "999")
	this.So(add("999", "1"), should.Equal, "1,000")
	this.So(add("1000", "1000"), should.Equal, "2,000")
	this.So(add("10000", "10000"), should.Equal, "20,000")
	this.So(add("1234", "4321"), should.Equal, "5,555")
	this.So(add("999", "999"), should.Equal, "1,998")
	this.So(add("123456789", "987654321"), should.Equal, "1,111,111,110")
	this.So(add("701,408,733", "1,134,903,170"), should.Equal, "1,836,311,903")
}

func add(a, b string) string {
	a = strings.Replace(a, ",", "", -1)
	a1, _ := strconv.Atoi(a)
	b = strings.Replace(b, ",", "", -1)
	b1, _ := strconv.Atoi(b)
	return NewBigFromSmall(a1).Plus(NewBigFromSmall(b1)).String()
}
