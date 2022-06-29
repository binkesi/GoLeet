package classdesign

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type CodecUrl struct {
	urlMap map[string]string
}

func ConstructorUrl() CodecUrl {
	return CodecUrl{urlMap: make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *CodecUrl) encode(longUrl string) string {
	var s strings.Builder
	rand.Seed(time.Now().Unix())
	for i := 0; i < 6; i++ {
		b := byte('a' + rand.Intn(25))
		s.WriteByte(b)
	}
	for {
		if _, ok := this.urlMap[s.String()]; !ok {
			break
		} else {
			x := byte('a' + rand.Intn(25))
			s.WriteByte(x)
		}
	}
	fmt.Println(s.String())
	this.urlMap[s.String()] = longUrl
	return s.String()
}

// Decodes a shortened URL to its original URL.
func (this *CodecUrl) decode(shortUrl string) string {
	return this.urlMap[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
