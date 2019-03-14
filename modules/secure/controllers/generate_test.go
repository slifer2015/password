package controllers

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func findChar(a string, b string) int {
	var count int
	for i := range a {
	innerLoop:
		for j := range b {
			if a[i] == b[j] {
				count++
				break innerLoop
			}

		}
	}
	return count
}

func TestRandPass(t *testing.T) {
	convey.Convey("Rand pass check", t, func() {
		convey.Convey("check size of each chars", func() {
			size := 10
			special := 2
			num := 6
			res := RandPass(size, special, num)
			specialCount := findChar(res, specialChars)
			letterCount := findChar(res, letters)
			numCount := findChar(res, numbers)
			convey.So(len(res), convey.ShouldEqual, size)
			convey.So(specialCount, convey.ShouldEqual, special)
			convey.So(letterCount, convey.ShouldEqual, letterCount)
			convey.So(numCount, convey.ShouldEqual, num)
		})

	})
}
