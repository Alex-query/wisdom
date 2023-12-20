package challenge

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"hash"
	"math"
	"strconv"
	"strings"
	"time"
	"wisdom/internal/infrastructure/config"
)

// Hash Hashier is the interface for a hashcash stamp generator.
type Hash struct {
	hasher  hash.Hash // SHA-1
	bits    uint      // Number of zero bits
	zeros   uint      // Number of zero digits
	saltLen uint      // Random salt length
	extra   string    // Extension to add to the minted stamp
}

// NewHashierSha1 creates a new Hashier using SHA-1.
func NewHashierSha1(config config.ChallengeConfig) *Hash {
	h := &Hash{
		hasher:  sha1.New(),
		bits:    config.GetNumberOfZeroBits(),
		saltLen: config.GetSaltLength(),
		extra:   config.GetExtra()}
	h.zeros = uint(math.Ceil(float64(h.bits) / 4.0))
	return h
}

// Date field format
const dateFormat = "060102"

// Mint a new hashcash stamp for resource.
func (h *Hash) Mint(prefixToken string) (string, error) {
	counter := 0
	var stamp string
	for {
		stamp = fmt.Sprintf("%s:%x",
			prefixToken, counter)
		if h.checkZeros(stamp) {
			return stamp, nil
		}
		counter++
	}
}

// GeneratePrefixToken generates a prefix token for a hashcash stamp.
func (h *Hash) GenerateRandomPrefixToken() (string, error) {
	date := time.Now().Format(dateFormat)
	resource := "wisdom"
	salt, err := h.getSalt()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("1:%d:%s:%s:%s:%s",
		h.bits, date, resource, h.extra, salt), nil
}

// Check whether a hashcash stamp is valid.
func (h *Hash) Check(stamp string, rightPrefix string) bool {
	if !strings.HasPrefix(stamp, rightPrefix) {
		return false
	}
	if h.checkDate(stamp) {
		return h.checkZeros(stamp)
	}
	return false
}

func (h *Hash) getSalt() (string, error) {
	buf := make([]byte, h.saltLen)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	salt := base64.StdEncoding.EncodeToString(buf)
	return salt[:h.saltLen], nil
}

func (h *Hash) checkZeros(stamp string) bool {
	h.hasher.Reset()
	h.hasher.Write([]byte(stamp))
	sum := h.hasher.Sum(nil)
	sumUint64 := binary.BigEndian.Uint64(sum)
	sumBits := strconv.FormatUint(sumUint64, 2)
	zeroes := 64 - len(sumBits)

	return uint(zeroes) >= h.bits
}

func (h *Hash) checkDate(stamp string) bool {
	fields := strings.Split(stamp, ":")
	if len(fields) != 7 {
		return false
	}
	then, err := time.Parse(dateFormat, fields[2])
	if err != nil {
		return false
	}
	duration := time.Since(then)
	return duration.Hours()*2 <= 48
}
