package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Time() int64 {
	return time.Now().Unix()
}

func Rand(min, max int64) int64 {
	if min > max {
		max, min = min, max
	} else if min == max {
		return min
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := r.Float64()*float64(max-min) + float64(min)
	return int64(math.Floor(t + 0.5))
}

func Salt() string {
	return fmt.Sprintf("%06s", strconv.FormatInt(Rand(1, 2176782336), 36))
}
