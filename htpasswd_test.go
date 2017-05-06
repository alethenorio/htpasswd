package main

import "testing"

func fixedSalt() []byte {
	return []byte("abcdefgh")
}

var hashtests = []struct {
	username string
	algo     string
	plain    string
	line     string
}{
	{"user", "SHA1", "helloworld", "user:{SHA}at+xg6SiyUovktq1redipHiJpaE="},
	{"user", "MD5", "helloworld", "user:$apr1$abcdefgh$xa95zAChgav3pqV63ZBU.1"},
}

func TestNewHash(t *testing.T) {
	saltFunc = fixedSalt
	defer func() { saltFunc = generateSalt }()

	for _, d := range hashtests {
		provider, err := NewHash(d.algo)
		if err != nil {
			t.Error(err)
		}

		if o := provider.FormattedAuth(d.username, d.plain); o != d.line {
			t.Errorf("%q hash(%q) => %q, want %q", d.algo, d.plain, o, d.line)
		}
	}
}
