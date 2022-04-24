package leetcode

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	return longUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return shortUrl
}
