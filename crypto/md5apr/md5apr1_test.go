package md5apr1

import (
	"testing"
)

type md5Apr1Test struct {
	out  string
	in   string
	salt string
}

// The data has been generated from `openssl passwd -apr1`
// The values have been chosen from the test in the standard library for MD5 hashing
var data = []md5Apr1Test{
	/*{"74UOCdrQgBQA4n1PBS.s71", "", "somesalt"},*/
	{"ikdeSwhXrwbHiQ/BzWjMc.", "a", "somesalt"},
	{"y9H1RoK9FTsnWW0.xPKoJ/", "ab", "somesalt"},
	{"1Tg85DBvZh.3pP9Oj3ci90", "abc", "somesalt"},
	{"ArWhd5yIxGleNwpFdtPR21", "abcd", "somesalt"},
	{"O0sH/nfGq7kzVcX2zMEX3.", "abcde", "somesalt"},
	{"jPf5tJLBpTOAoKz61S0YZ1", "abcdef", "somesalt"},
	{"yfJUh3G9dhq6B2gcazxr00", "abcdefg", "somesalt"},
	{"PjJx0sxmp1IW0JfGEQ7m81", "abcdefgh", "somesalt"},
	{"Fm1ORSkI3HDoGRNJLDc4J1", "abcdefghi", "somesalt"},
	{"M63WucNpkwiK6QHzfH63/0", "abcdefghij", "somesalt"},
	{"N1CpryXUkvtSx4i.gNTxt0", "Discard medicine more than two years old.", "somesalt"},
	{"pGkBU9DwwMzFMUZ/h9yvb.", "He who has a shady past knows that nice guys finish last.", "somesalt"},
	{"T1knbizA.0dO0aXeGpvVS.", "I wouldn't marry him with a ten foot pole.", "somesalt"},
	{"eLIBY16s04n7RQKUo9mdT1", "The days of the digital watch are numbered.  -Tom Stoppard", "somesalt"},
	{"tE.n3sYviUSdIbsiZG4GY1", "Nepal premier won't resign.", "somesalt"},
	{"jpwnQk/X/lHhVQsm3ZzmX/", "For every action there is an equal and opposite government program.", "somesalt"},
	{"BSXVjZZbnk1WDXPXF9s6z1", "His money is twice tainted: 'taint yours and 'taint mine.", "somesalt"},
	{"XDUQ6RVH77IL485sg3QbL/", "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977", "somesalt"},
	{"eO4HCJ22nne7IcVjiPSPc/", "It's a tiny change to the code and not completely disgusting. - Bob Manchek", "somesalt"},
	{"A444c20bA.PJArCCVFSIK1", "size:  a.out:  bad magic", "somesalt"},
	{"S72eTPlywyL2reArsQ6JY1", "The major problem is with sendmail.  -Mark Horton", "somesalt"},
	{"PLtyCcVy.LknM6tJyEAUj0", "Give me a rock, paper and scissors and I will move the world.  CCFestoon", "somesalt"},
	{".fI.FnrKVQN3N1cHMURjo0", "If the enemy is within range, then so are you.", "somesalt"},
	{"YFC7yQ0M.2FT3EFn4/wbb/", "It's well we cannot hear the screams/That we create in others' dreams.", "somesalt"},
	{"nXkxd9JaYFekdoRusdmt4.", "You remind me of a TV show, but that's all right: I watch it anyway.", "somesalt"},
	{"hkfl4dDz1P2SC5Ja23aWZ0", "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley", "somesalt"},
	{"i.ZcGl0cAu1OrgKSsVSYc1", "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule", "somesalt"},
	{"OPMdaznlSUV9XuqzB/tqe.", "How can you write a big system without C++?  -Paul Glick", "somesalt"},
	// Shells have some issues with the exclamation mark as it is a special character and I couldn't confirm if escaping it impacted the result of openssl or not so disabled these for now
	/*{"bDAnqgg1u/BVMyXhXBuOI.", "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave", "somesalt"},
	{"469b13a78ebf297ecda64d4723655154", "C is as portable as Stonehedge!!", "somesalt"},*/
}

func TestMd5Apr_checksum(t *testing.T) {
	for i := 0; i < len(data); i++ {
		g := data[i]
		md5apr := FromSalt([]byte(g.salt))
		md5apr.Write([]byte(g.in))
		o := string(md5apr.Sum(nil))
		if o != g.out {
			t.Errorf("md5apr(%s, %s) = %s want %s", g.salt, g.in, o, g.out)
		}

	}
}
