package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/ByteFlinger/htpasswd/crypto/md5apr"
	"hash"
	"io"
	"log"
)

type htpasswd struct {
	entries map[string][]byte // maps username to password byte slice.
}

type HashProvider struct {
	provider  hash.Hash
	prefix    string
	formatter func([]byte) []byte
}

const (
	SHA1 = iota
	MD5
)

var algoNames = map[string]int{
	"SHA1": SHA1,
	"MD5":  MD5,
}

var saltFunc = generateSalt

func NewHash(algo string) (*HashProvider, error) {

	var hashProvider *HashProvider
	switch algoNames[algo] {
	case SHA1:
		hashProvider = &HashProvider{
			provider: sha1.New(),
			prefix:   "{SHA}",
			formatter: func(src []byte) []byte {
				dstStr := base64.StdEncoding.EncodeToString(src)
				return []byte(dstStr)
			},
		}
	case MD5:
		salt := saltFunc()
		hashProvider = &HashProvider{
			provider: md5apr1.FromSalt(salt),
			prefix:   "$apr1$" + string(salt) + "$",
			formatter: func(src []byte) []byte {
				return src
			},
		}
	default:
		return nil, errors.New("Unsupported algorithm " + algo)
	}

	return hashProvider, nil
}

func generateSalt() []byte {
	bs := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, bs)
	if err != nil {
		log.Fatal(err)
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(bs)))
	base64.RawStdEncoding.Encode(dst, bs)
	n := len(dst)
	if n > 8 {
		n = 8
	}
	return dst[:n]
}

// Hashes a given password
func (h *HashProvider) hash(password string) string {
	h.provider.Write([]byte(password))
	return string(h.formatter(h.provider.Sum(nil)))
}

// Returns a properly formatted .htpasswd string based on the chosen hashing algorithm
func (h *HashProvider) FormattedAuth(username, password string) string {
	return username + ":" + h.prefix + h.hash(password)
}
