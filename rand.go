package main

import (
	"errors"
	"math/rand"
	"time"
)

var (
	endLessThanCountErr = errors.New("luck-draw:rand num:end can not less than count")
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randNum(end, count int) (result []int, err error) {
	if end < count {
		return nil, endLessThanCountErr
	}
	m := make(map[int]interface{}, count)
	for len(m) < count {
		m[rand.Intn(end)] = nil
	}

	for k := range m {
		result = append(result, k)
	}

	return
}
