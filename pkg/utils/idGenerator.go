package utils

import (
	"fmt"
	"time"
)

func GenerateTransactionID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
