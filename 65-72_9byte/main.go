package main

import (
	"bufio"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"log"
	mrand "math/rand"
	"os"
	"regexp"
	"sort"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/ripemd160"
)

var (
	filePath *string
	addrType *string
	data     [][]byte
)

func findAddr(d [][]byte, s []byte) int {
	// defer timer("findAddr")()
	l := len(d)
	i := sort.Search(l, func(i int) bool { return bytes.Compare(d[i], s) >= 0 })
	if i < l && bytes.Equal(d[i], s) {
		return i
	}
	return -1
}

func sortByteMatrix(d [][]byte) [][]byte {
	sort.Slice(d, func(i, j int) bool {
		return bytes.Compare(d[i], d[j]) < 0
	})
	return d
}

func truncateDoubles(d [][]byte) [][]byte {
	d = sortByteMatrix(d)
	for i := 0; i < len(d); i++ {
		if i == 0 {
			continue
		}
		if bytes.Equal(d[i], d[i-1]) {
			log.Printf("Found double at %d\n", i)
			d[i] = d[len(d)-1]     // Copy last element to index i.
			d[len(d)-1] = []byte{} // Erase last element (write zero value).
			d = d[:len(d)-1]       // Truncate slice.
		}
	}
	return d
}

func importFromFile(f, t *string) [][]byte {

	var p *regexp.Regexp

	file, err := os.Open(*f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	switch {
	case *t == "p2pkh":
		p, _ = regexp.Compile(`.?(1[1-9A-HJ-NP-Za-km-z]{32,34}).?`)
	case *t == "p2sh":
		p, _ = regexp.Compile(`([^1-9A-HJ-NP-Za-km-z](3[1-9A-HJ-NP-Za-km-z]{32,34})[^1-9A-HJ-NP-Za-km-z])\n`)
	}

	s := bufio.NewScanner(file)
	result := make([][]byte, 0)
	for s.Scan() {
		if p.MatchString(s.Text()) {
			found := string(p.FindStringSubmatch(s.Text())[1])
			decoded, _, err := base58.CheckDecode(found)
			if err != nil {
				log.Println(err)
				log.Printf("Err with: %s\n", found)
				continue
			}
			result = append(result, decoded)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s imported total:\t%d\n", *addrType, len(result))
	return result
}

func getPubKey(k []byte) []byte {
	// defer timer("getPubKey")()
	var e ecdsa.PrivateKey
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(k) // took 121.734µs
	return elliptic.MarshalCompressed(secp256k1.S256(), e.X, e.Y)      // took 11.714µs
}

func timer(name string, i *int) func() {
	start := time.Now()
	return func() {
		log.Printf("%s took %v, i=%d\n", name, time.Since(start), *i)
	}
}

func sha256rmd160(p []byte) []byte {
	// defer timer("sha256rmd160")()
	h := sha256.New()
	h.Write(p)
	r := ripemd160.New()
	r.Write(h.Sum(nil))
	return r.Sum(nil)
}

func p2pkhWIF(p []byte) string {
	// defer timer("p2pkhWIF")()
	enc := base58.CheckEncode(p, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return enc
}

func privateKeyToWIF(p []byte) string {
	// defer timer("privateKeyToWIF")()
	p = append(p, 0x01)
	enc := base58.CheckEncode(p, 0x80)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return enc
}

func PWIFTtoKey(s string) []byte {
	// defer timer("PWIFTtoKey")()
	k, _, err := base58.CheckDecode(s)
	if err != nil {
		log.Fatal(err)
	}
	return k
}

func getRandomKey() []byte {
	// defer timer("getRandomKey")()
	k := make([]byte, 32)
	_, err := rand.Read(k)
	if err != nil {
		log.Fatalf("error while generating random string: %s", err)
	}
	return k
}

func getRandomKeyByRange9byte() []byte { // 67,68,69,71,72, exclude 65,66,70
	// defer timer("getRandomKeyByRange9byte")()
	//                                               9 8 7 6 5 4 3 2 1  // byte
	// 0000000000000000000000000000000000000000000000010000000000000000 // 65 Solved
	// 000000000000000000000000000000000000000000000001ffffffffffffffff //
	// 0000000000000000000000000000000000000000000000020000000000000000 // 66 Solved
	// 000000000000000000000000000000000000000000000003ffffffffffffffff //
	// 0000000000000000000000000000000000000000000000040000000000000000 // 67 Solved
	// 000000000000000000000000000000000000000000000007ffffffffffffffff //
	// 0000000000000000000000000000000000000000000000080000000000000000 // 68
	// 00000000000000000000000000000000000000000000000fffffffffffffffff //
	// 0000000000000000000000000000000000000000000000100000000000000000 // 69
	// 00000000000000000000000000000000000000000000001fffffffffffffffff //
	// 0000000000000000000000000000000000000000000000200000000000000000 // 70 Solved
	// 00000000000000000000000000000000000000000000003fffffffffffffffff //
	// 0000000000000000000000000000000000000000000000400000000000000000 // 71
	// 00000000000000000000000000000000000000000000007fffffffffffffffff //
	// 0000000000000000000000000000000000000000000000800000000000000000 // 72
	// 0000000000000000000000000000000000000000000000ffffffffffffffffff //
	r := make([]byte, 23, 32)
	k := make([]byte, 9)
	for {
		_, err := rand.Read(k)
		if err != nil {
			log.Fatalf("error while generating random string: %s", err)
			continue
		}
		if (k[0] >= 0x20 && k[0] <= 0x3f) || (k[0] >= 0x01 && k[0] <= 0x07) { // exclude solved (9 byte)
			// log.Printf("Ignored key: %x", k)
			continue
		}
		break
	}
	r = append(r, k...)
	// log.Printf("Key: %x", r)
	return r
}

func getKnownKey() []byte {
	keys := []string{
		// https://btcpuzzle.info/puzzlelist
		"000000000000000000000000000000000000000af55fc59c335c8ec67ed24826", // 100
		"000000000000000000000000000000000000016f14fc2054cd87ee6396b33df3", // 105
		"00000000000000000000000000000000000035c0d7234df7deb0f20cf7062444", // 110
		"0000000000000000000000000000000000060f4d11574f5deee49961d9609ac6", // 115
	}
	key, _ := hex.DecodeString(keys[mrand.Intn(len(keys))])
	return key
}

func init() {
	filePath = flag.String("file", "./btc_addreses.txt", "Path to file")
	addrType = flag.String("address-type", "p2pkh", "Type of addresses to import. Support p2pkh, p2sh")
}

func main() {

	flag.Parse()

	data = importFromFile(filePath, addrType)
	data = truncateDoubles(data)
	data = sortByteMatrix(data)

	for {
		// key := getKnownKey()
		// key := getRandomKey()
		key := getRandomKeyByRange9byte()
		// key, _ := hex.DecodeString("00000000000000000000000000000000000000000000005906a7d53cdab02d94")
		// log.Printf("Key is %x\n", key)
		pubkey := getPubKey(key)
		rawAddr := sha256rmd160(pubkey)
		pos := findAddr(data, rawAddr) // findAddr took 6.368µs
		if pos != -1 {
			log.Printf("Found at %d\n", pos)
			log.Printf("Key: %x\n", key)
			log.Printf("Addr (%s): %s\n", *addrType, p2pkhWIF(rawAddr))
		}
	}
}
