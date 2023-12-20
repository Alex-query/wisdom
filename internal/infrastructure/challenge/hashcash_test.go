package challenge

import (
	"strconv"
	"strings"
	"testing"
	"time"
	"wisdom/internal/infrastructure/config"
)

var stampTests = []struct {
	bits     uint
	saltLen  uint
	extra    string
	resource string
}{
	{20, 8, "", "wisdom"},
	{10, 10, "rara", "wisdom"},
	{20, 10, "ddd", "wisdom"},
	{15, 4, "", "wisdom"},
}

func TestStampFormat(t *testing.T) {
	expectedDate := time.Now().Format("060102")
	for _, tt := range stampTests {
		conf := config.NewChallengeConfig(tt.bits, tt.saltLen, tt.extra)
		h := NewHashierSha1(conf)
		prefixToken, err := h.GenerateRandomPrefixToken()
		if err != nil {
			t.Errorf("GenerateRandomPrefixToken failed with error %v", err)
		}
		stamp, err := h.Mint(prefixToken)
		if err != nil {
			t.Errorf("Mint failed for %s with error %v", tt.resource, err)
		}
		fields := strings.Split(stamp, ":")
		if len(fields) != 7 {
			t.Errorf("Expected 7 fields got %d", len(fields))
		}
		ver, err := strconv.Atoi(fields[0])
		if err != nil {
			t.Errorf("Expected version 1, got error %v", err)
		}
		if ver != 1 {
			t.Errorf("Expected version 1, got %d", ver)
		}
		bits, err := strconv.ParseUint(fields[1], 10, 32)
		if err != nil {
			t.Errorf("Expected %d bits, got error %v", tt.bits, err)
		}
		if uint(bits) != tt.bits {
			t.Errorf("Expected %d bits, got %d", tt.bits, bits)
		}
		date := fields[2]
		if date != expectedDate {
			t.Errorf("Expected %s date, got %s", expectedDate, date)
		}
		resource := fields[3]
		if resource != tt.resource {
			t.Errorf("Expected %s resource, got %s", tt.resource, resource)
		}
		extra := fields[4]
		if extra != tt.extra {
			t.Errorf("Expected %s extra, got %s", tt.extra, extra)
		}
		salt := fields[5]
		if uint(len(salt)) != tt.saltLen {
			t.Errorf("Expected %d salt chars, got %d", tt.saltLen, len(salt))
		}
		counter := fields[6]
		if counter == "" {
			t.Errorf("Counter field is empty")
		}
	}
}

func TestMintAndCheck(t *testing.T) {
	conf := config.NewChallengeConfig(20, 8, "")
	h := NewHashierSha1(conf)
	for i := 0; i < 10; i++ {
		prefix, err := h.GenerateRandomPrefixToken()
		if err != nil {
			t.Errorf("GenerateRandomPrefixToken failed with error %v", err)
		}
		stamp, err := h.Mint(prefix)
		if err != nil {
			t.Errorf("Mint failed for %s with error %v", prefix, err)
		}
		if !h.Check(stamp, prefix) {
			t.Errorf("Check failed for %s , stamp %s", prefix, stamp)
		}
	}
}
