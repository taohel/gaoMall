package conf

import "time"

type Token struct {
	Duration     time.Duration
	SymmetricKey string
}
