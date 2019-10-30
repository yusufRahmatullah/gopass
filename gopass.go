package gopass

import (
	"bytes"
	"math"
	"strings"
)

var (
	pool     []int
	width    = 256
	chunks   = 6
	digits   = 52
	nextMax  = []int{10, 26, 26, 10}
	nextChar = []string{
		"!@#$%^&*-_",
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
	}
	startdenom   = int64(math.Pow(float64(width), float64(chunks)))
	significance = int64(math.Pow(2, float64(digits)))
	overflow     = significance * 2
	mask         = width - 1
)

type ARC4 struct {
	i, j int
	S    []int
}

func NewArc4(key []int) *ARC4 {
	if len(key) == 0 {
		key = []int{0}
	}
	arc := &ARC4{}
	t := 0
	keylen := len(key)
	i := 0
	j := 0
	arc.i = 0
	arc.j = 0
	for ii := 0; ii < 256; ii++ {
		arc.S = append(arc.S, 0)
	}
	s := arc.S
	for i < width {
		s[i] = i
		i++
	}
	for ii := 0; ii < width; ii++ {
		t = s[ii]
		j = mask & (j + key[ii%keylen] + t)
		s[ii] = s[j]
		s[j] = t
	}
	arc.g(width)

	return arc
}

func (arc *ARC4) g(count int) int64 {
	t := 0
	r := int64(0)
	i := arc.i
	j := arc.j
	s := arc.S
	for count > 0 {
		i = mask & (i + 1)
		t = s[i]
		j = mask & (j + t)
		s[i] = s[j]
		s[j] = t
		k := s[i] + s[j]
		mk := mask & k
		rws := r*int64(width) + int64(s[mk])
		r = int64(float64(rws))
		count--
	}

	arc.i = i
	arc.j = j
	return r
}

func GenPass(seed string) string {
	prng := Seedrandom(seed)
	var result []string
	for i := 0; i < 64; i++ {
		nt := int(prng() * 4)
		nm := nextMax[nt]
		nc := int(prng() * float64(nm))
		result = append(result, string(nextChar[nt][nc]))
	}
	return strings.Join(result, "")
}

func GenPin(seed string) string {
	prng := Seedrandom(seed)
	var result []string
	for i := 0; i < 6; i++ {
		nc := int(prng() * 10)
		result = append(result, string(nextChar[3][nc]))
	}
	return strings.Join(result, "")
}

func flatten(d string, n int) string {
	return d
}

func getSmear(key []int, idx int) int {
	if idx >= len(key) {
		return 0
	}
	return key[idx] * 19
}

func mixkey(seed string, inputKey []int) (string, []int) {
	key := inputKey
	for i := 0; i < 255-len(inputKey); i++ {
		key = append(key, 0)
	}
	smear := 0
	j := 0
	for j < len(seed) {
		smear ^= getSmear(key, mask&j)
		key[mask&j] = mask & (smear + ord(string(seed[j])))
		j++
	}
	return toString(key), trimKey(key)
}

func ord(c string) int {
	return int([]rune(c)[0])
}

func Seedrandom(seed string) func() float64 {
	var key []int
	ff := flatten(seed, 3)
	_, key = mixkey(ff, key)
	arc := NewArc4(key)
	ss := toString(arc.S)
	_, pool = mixkey(ss, pool)

	return func() float64 {
		return sprng(arc)
	}
}

func sprng(arc *ARC4) float64 {
	n := arc.g(chunks)
	d := float64(startdenom)
	x := int64(0)
	for n < significance {
		n = (n + x) * int64(width)
		d *= float64(width)
		x = arc.g(1)
	}
	for n >= overflow {
		n /= 2
		d /= 2
		x = int64(int(x) >> 1)
	}
	nx := (n + x)
	nxd := float64(nx) / d
	return nxd
}

func toString(key []int) string {
	trimmed := trimKey(key)
	var chara []string
	bs := bytes.NewBufferString("")
	for _, v := range trimmed {
		chara = append(chara, string(v))
		bs.WriteByte(byte(v))
	}
	res := bs.String()
	return res
}

func trimKey(key []int) []int {
	var result []int
	for _, v := range key {
		if v != 0 {
			result = append(result, v)
		}
	}
	return result
}
