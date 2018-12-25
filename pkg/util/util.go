package util

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
)

var (
	seededRand *rand.Rand
)

// GenerateSalt is generate salt
func GenerateSalt(length int) string {
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GeneratePassword is generate password
func GeneratePassword(src, salt string) string {
	sa := Sha256String(src)
	return Sha256String(sa + salt)
}

// HashAndSalt is
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if nil != err {
		return ""
	}
	return string(hash)
}

// Sha256String is sha256 string
func Sha256String(str string) string {
	m := sha256.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
