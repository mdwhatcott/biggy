package biggy

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
)

type Big struct{ triplets *list.List }

func ParseBig(raw string) *Big {
	triplets := list.New()

	for len(raw) > 0 {
		if len(raw) <= 3 {
			num, _ := strconv.Atoi(raw)
			triplets.PushBack(num)
			raw = ""
		} else {
			num, _ := strconv.Atoi(raw[len(raw)-3:])
			triplets.PushBack(num)
			raw = raw[:len(raw)-3]
		}
	}
	return &Big{triplets: triplets}
}

func (this *Big) Plus(that *Big) *Big {
	i, a := this.triplets.Front(), that.triplets.Front()
	var carry int
	for {
		sum := i.Value.(int) + a.Value.(int) + (carry)
		carry = 0
		for sum > 999 {
			carry += 1
			sum -= 1000
		}
		i.Value = sum
		i, a = i.Next(), a.Next()
		if i == nil && a == nil {
			break
		} else if i == nil {
			i = this.triplets.PushBack(0)
		} else if a == nil {
			a = that.triplets.PushBack(0)
		}
	}
	if carry > 0 {
		this.triplets.PushBack(carry)
	}
	return this
}

func (this *Big) String() string {
	buffer := new(bytes.Buffer)
	for e := this.triplets.Back(); e != nil; e = e.Prev() {

		if e.Value.(int) < 100 && e != this.triplets.Back() {
			fmt.Fprintf(buffer, "%03d", e.Value)
		} else {
			fmt.Fprintf(buffer, "%d", e.Value)
		}

		if e != this.triplets.Front() {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}
