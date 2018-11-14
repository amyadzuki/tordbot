package main

import (
	"math/rand"
	"time"
)

var PRG *rand.Rand

func init() {
	PRG = rand.New(rand.NewSource(time.Now().UnixNano()))
}
