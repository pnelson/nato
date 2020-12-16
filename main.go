package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var (
	help = flag.Bool("h", false, "show this usage information")
	n    = flag.Bool("n", false, "output with newlines")
)

var m = map[rune]string{
	'a': "alfa",
	'b': "bravo",
	'c': "charlie",
	'd': "delta",
	'e': "echo",
	'f': "foxtrot",
	'g': "golf",
	'h': "hotel",
	'i': "india",
	'j': "juliett",
	'k': "kilo",
	'l': "lima",
	'm': "mike",
	'n': "november",
	'o': "oscar",
	'p': "papa",
	'q': "quebec",
	'r': "romeo",
	's': "sierra",
	't': "tango",
	'u': "uniform",
	'v': "victor",
	'w': "whisky",
	'x': "x-ray",
	'y': "yankee",
	'z': "zulu",
	'0': "zero",
	'1': "one",
	'2': "two",
	'3': "three",
	'4': "four",
	'5': "five",
	'6': "six",
	'7': "seven",
	'8': "eight",
	'9': "nine",
}

func init() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] STRING\n\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "%s accepts exactly one argument\n", os.Args[0])
		os.Exit(1)
		return
	}
	arg := flag.Arg(0)
	nato := make([]string, len(arg))
	for i, r := range arg {
		l := unicode.ToLower(r)
		s, ok := m[l]
		if !ok {
			fmt.Fprintf(os.Stderr, "%s accepts only alphanumeric characters\n", os.Args[0])
			os.Exit(1)
			return
		}
		if unicode.IsUpper(r) {
			nato[i] = strings.ToUpper(s)
		} else {
			nato[i] = s
		}
	}
	out := ""
	if *n {
		out = strings.Join(nato, "\n")
	} else {
		out = strings.Join(nato, " ")
	}
	fmt.Println(out)
}
