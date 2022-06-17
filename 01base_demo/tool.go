package base_demo

import (
	"crypto/md5"
	"math/rand"
)

func getMD5(s []byte) []byte {
	md := md5.New()
	md.Write(s)
	x := md.Sum([]byte(""))
	return x
}

func T() string {
	name := randomString()
	x := getMD5(name)
	if x[0] > 127 {
		return "A"
	} else {
		return "B"
	}
}

func randomString() []byte {
	var x []byte
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		x = append(x, byte(a+33))
	}
	return x
}
