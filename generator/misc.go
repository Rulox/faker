package generator

import (
	"time"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	"crypto/sha256"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Misc Generator. Generates random basic data that does not require locale configuration.
type MiscGenerator struct {}

func (mg *MiscGenerator) RandomInt() int {
	return rand.Int()
}

func (mg *MiscGenerator) RandomFloat32() float32 {
	return rand.Float32()
}

func (mg *MiscGenerator) RandomFloat64() float64 {
	return rand.Float64()
}

func (mg *MiscGenerator) RandomIntBelow(t int) int {
	return rand.Intn(t)
}

func (mg *MiscGenerator) UnixTime() int64 {
	return time.Now().Unix()
}

// Return a random NONSENSE array of bytes used for hashing functions
func (mg *MiscGenerator) randomString(s int) []byte {
	b := make([]byte, mg.RandomIntBelow(s))
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

// Hash a random string of 50 characters using md5.
func (mg *MiscGenerator) Md5() string {
	h := md5.New()
	h.Write(mg.randomString(50))
	return hex.EncodeToString(h.Sum(nil))
}

// Hash a random string of 50 characters using sha1.
func (mg *MiscGenerator) Sha1() string {
	h := sha1.New()
	h.Write(mg.randomString(50))
	return hex.EncodeToString(h.Sum(nil))
}

// Hash a random string of 50 characters using sha256.
func (mg *MiscGenerator) Sha256() string {
	h := sha256.New()
	h.Write(mg.randomString(50))
	return hex.EncodeToString(h.Sum(nil))
}
