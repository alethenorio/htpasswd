// Package md5apr implements the MD5 APR hash algorithm
// This is a pretty ugly implementation of the hash.Hash api as I am not a crypto expert
// The algorithm implementation for MD5 APR1 is based on the one found here https://github.com/jimstudt/http-authentication/blob/master/basic/md5.go

package md5apr1

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"hash"
	"io"
	"log"
)

// The size of an MD% APR Salt
const PW_SALT_BYTES = 8

// The size of an MD5 APR1 checksum in bytes.
const Size = 22

type md5Apr struct {
	salt []byte
	pass []byte
}

var _ hash.Hash = &md5Apr{}

// New returns a new md5Apr1 hash computing the MD5Apr1 checksum based on a randomly generated salt
func New() hash.Hash {
	return newMd5Apr1([]byte{})
}

// FromSalt returns a new md5Apr1 hash computing the MD5Apr1 checksum based on a fixed salt value
// Take care to pass a randomly generated salt if using this
func FromSalt(salt []byte) hash.Hash {
	return newMd5Apr1(salt)
}

func newMd5Apr1(salt []byte) hash.Hash {

	var bs []byte

	if len(salt) == 0 {
		bs = generateSalt()
	} else {
		n := len(salt)
		if len(salt) > 8 {
			n = 8
		}

		bs = salt[:n]
	}

	return &md5Apr{
		salt: bs,
	}

}

func generateSalt() []byte {
	bs := make([]byte, PW_SALT_BYTES)
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

func (m *md5Apr) checksum() []byte {

	bin := make([]byte, len(m.pass))
	text := make([]byte, len(m.pass))

	copy(bin, m.pass)
	copy(text, m.pass)

	bin = append(bin, m.salt...)
	bin = append(bin, m.pass...)

	// start with a hash of password and salt
	initBin := md5.Sum(bin)

	text = append(text, "$apr1$"...)
	text = append(text, m.salt...)
	// begin an initial string with hash and salt
	initText := bytes.NewBuffer(text)

	// add crap to the string willy-nilly
	for i := len(m.pass); i > 0; i -= 16 {
		lim := i
		if lim > 16 {
			lim = 16
		}
		initText.Write(initBin[0:lim])
	}

	// add more crap to the string willy-nilly
	for i := len(m.pass); i > 0; i >>= 1 {
		if (i & 1) == 1 {
			initText.WriteByte(byte(0))
		} else {
			initText.WriteByte(m.pass[0])
		}
	}

	// Begin our hashing in earnest using our initial string
	h := md5.Sum(initText.Bytes())

	n := bytes.NewBuffer([]byte{})

	for i := 0; i < 1000; i++ {
		// prepare to make a new muddle
		n.Reset()

		// alternate password+crap+bin with bin+crap+password
		if (i & 1) == 1 {
			n.Write(m.pass)
		} else {
			n.Write(h[:])
		}

		// usually add the salt, but not always
		if i%3 != 0 {
			n.Write(m.salt)
		}

		// usually add the password but not always
		if i%7 != 0 {
			n.Write(m.pass)
		}

		// the back half of that alternation
		if (i & 1) == 1 {
			n.Write(h[:])
		} else {
			n.Write(m.pass)
		}

		// replace bin with the md5 of this muddle
		h = md5.Sum(n.Bytes())
	}

	// At this point we stop transliterating the PHP code and flip back to
	// reading the Apache source. The PHP uses their base64 library, but that
	// uses the wrong character set so needs to be repaired afterwards and reversed
	// and it is just really weird to read.

	result := bytes.NewBuffer([]byte{})

	// This is our own little similar-to-base64-but-not-quite filler
	fill := func(a byte, b byte, c byte) {
		v := (uint(a) << 16) + (uint(b) << 8) + uint(c) // take our 24 input bits

		for i := 0; i < 4; i++ { // and pump out a character for each 6 bits
			result.WriteByte("./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"[v&0x3f])
			v >>= 6
		}
	}

	// The order of these indices is strange, be careful
	fill(h[0], h[6], h[12])
	fill(h[1], h[7], h[13])
	fill(h[2], h[8], h[14])
	fill(h[3], h[9], h[15])
	fill(h[4], h[10], h[5]) // 5?  Yes.
	fill(0, 0, h[11])

	return result.Bytes()[0:22] // we wrote two extras since we only need 22.
}

func (d *md5Apr) Write(p []byte) (nn int, err error) {
	d.pass = p
	return len(p), nil
}

func (d *md5Apr) Reset() {
	// Should we regenerate the salt here?
	d.pass = d.pass[:0]
}

func (d *md5Apr) Size() int { return Size }

// This is unimplemented and needs further analysis so we panic if anybody tries to use it
func (d *md5Apr) BlockSize() int { panic("BlockSize is currently not implemented") }

func (d0 *md5Apr) Sum(in []byte) []byte {
	// Make a copy of d0 so that caller can keep writing and summing.
	d := *d0
	hash := d.checksum()
	return append(in, hash[:]...)
}

// Sum returns the MD5 APR1 checksum of the data.
func Sum(data []byte) []byte {
	var d md5Apr
	d.Reset()
	d.Write(data)
	return d.checksum()
}
