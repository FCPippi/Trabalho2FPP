package simulador

import (
	"math/rand"
	"time"
)

func SimularFalhaRede() bool {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Float32() < 0.1
}

func SimularAtraso() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}
