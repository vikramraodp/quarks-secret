package inmemorygenerator

import (
	"go.uber.org/zap"
	"os"
	"strconv"
)

// InMemoryGenerator represents a secret generator that generates everything
// by itself, using no 3rd party tools
type InMemoryGenerator struct {
	Bits      int    // Key bits
	Expiry    int    // Expiration (days)
	Algorithm string // Algorithm type

	log *zap.SugaredLogger
}

// NewInMemoryGenerator creates a default InMemoryGenerator
func NewInMemoryGenerator(log *zap.SugaredLogger) *InMemoryGenerator {
	EXP,err := strconv.ParseUint(os.Getenv("CERT_EXPIRY_DAYS"), 10, 32)
	_ = err
	KEY_BITS,err2 := strconv.ParseUint(os.Getenv("CERT_KEY_BITS"), 10, 32)
	_ = err2
	return &InMemoryGenerator{Bits: int(KEY_BITS), Expiry: int(EXP), Algorithm: os.Getenv("CERT_ALGORITHM"), log: log}
}
