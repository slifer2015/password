package controllers

import (
	"errors"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"test.com/test/services/framework"
	"test.com/test/services/framework/controller"
	"test.com/test/services/framework/middleware"
)

const (
	specialChars = `@%+\/'!#$^?:,(){}[]~-_.`
	letters      = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	numbers      = `1234567890`
)

type generatePayload struct {
	Size         int `json:"size" validate:"required" message:"size required"`
	SpecialChars int `json:"special_chars" validate:"required" message:"special_chars required"`
	Numbers      int `json:"numbers" validate:"required" message:"numbers required"`
	Options      int `json:"options"`
}

type generateResponse struct {
	Data []string `json:"data"`
}

func generate(w http.ResponseWriter, r *http.Request) {
	p, ok := middleware.GetPayload(r)
	if !ok {
		framework.JSONErr(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}
	payload, ok := p.(*generatePayload)
	if !ok {
		framework.JSONErr(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}
	validationErrs, fail := controller.Validate(payload)
	if fail {
		framework.JSON(w, http.StatusBadRequest, validationErrs)
		return
	}

	// validate request data logic
	if payload.SpecialChars+payload.Numbers > payload.Size {
		framework.JSONErr(w, http.StatusBadRequest, errors.New("invalid data"))
		return
	}
	var option = 1
	if payload.Options > 0 {
		option = payload.Options
	}
	var finalRes = make([]string, option)
	for i := 0; i < option; i++ {
		finalRes[i] = RandPass(payload.Size, payload.SpecialChars, payload.Numbers)
	}
	framework.JSON(w, http.StatusOK, generateResponse{
		Data: finalRes,
	})

}

// RandPass generate random passwords
func RandPass(size, special, num int) string {
	var sChars = make([]string, special)
	for i := 0; i < special; i++ {
		n := rand.Intn(len(specialChars))
		sChars[i] = string(specialChars[n])
	}
	var nums = make([]string, num)
	for i := 0; i < num; i++ {
		n := rand.Intn(len(numbers))
		nums[i] = string(numbers[n])
	}
	var res = make([]string, 0)
	res = append(res, sChars...)
	res = append(res, nums...)
	remain := size - special - num
	if remain > 0 {
		var alphabet = make([]string, remain)
		for i := 0; i < remain; i++ {
			n := rand.Intn(len(letters))
			alphabet[i] = string(letters[n])
		}
		res = append(res, alphabet...)
	}

	final := Shuffle(res)

	return strings.Join(final, "")
}

// Shuffle shuffle array
func Shuffle(in []string) []string {
	res := make([]string, len(in))
	n := len(in)
	for i := 0; i < n; i++ {
		randIndex := rand.Intn(len(in))
		res[i] = in[randIndex]
		in = append(in[:randIndex], in[randIndex+1:]...)
	}
	return res
}

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}
