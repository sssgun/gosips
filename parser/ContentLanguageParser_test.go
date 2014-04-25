package parser

import (
	"testing"
)

func TestContentLanguageParser(t *testing.T) {
	var tvs = []string{
		"Content-Language: fr \n",
		"Content-Language: fr,he \n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewContentLanguageParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Content-Language: fr \n",
          "Content-Language: fr , he \n"
      };

      for (int i = 0; i < r.length; i++ ) {
          ContentLanguageParser parser =
          new ContentLanguageParser(r[i]);
          ContentLanguageList e= (ContentLanguageList) parser.parse();
          System.out.println("encoded = " + e.encode());
      }
  }
*/