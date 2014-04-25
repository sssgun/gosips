package parser

import (
	"testing"
)

func TestWWWAuthenticateParser(t *testing.T) {
	var tvs = []string{
		"WWW-Authenticate: Digest realm=\"MCI WorldCom SIP\"," +
			"domain=\"sip:ss2.wcom.com\",nonce=\"ea9c8e88df84f1cec4341ae6cbe5a359\"," +
			"opaque=\"\",stale=FALSE,algorithm=\"MD5\"\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewWWWAuthenticateParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}