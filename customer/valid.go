package customer

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Valid
func (this *Customer) Valid(echostr, signature, timestamp, nonce string) string {
	if this.checkSign(signature, timestamp, nonce) {
		return echostr
	}
	return echostr + this.ContactToken
}

// checkSign
func (this *Customer) checkSign(signature, timestamp, nonce string) bool {
	sl := []string{this.ContactToken, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	if fmt.Sprintf("%x", s.Sum(nil)) == signature {
		return true
	}
	return false
}
