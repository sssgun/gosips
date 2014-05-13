package parser

import (
	"gosips/core"
	"gosips/sip/header"
	"strconv"
)

/** SIPParser for SIP Expires SIPParser. Converts from SIP Date to the
* internal storage (Calendar).
*
*@version  JAIN-SIP-1.1
*
*@author M. Ranganathan <mranga@nist.gov>  <br/>
*
*<a href="{@docRoot}/uncopyright.html">This code is in the public domain.</a>
*
 */
type ExpiresParser struct {
	HeaderParser
}

/** protected constructor.
 *@param text is the text of the header to parse
 */
func NewExpiresParser(text string) *ExpiresParser {
	this := &ExpiresParser{}
	this.HeaderParser.super(text)
	return this
}

/** constructor.
 *@param lexer is the lexer passed in from the enclosing parser.
 */
func NewExpiresParserFromLexer(lexer core.Lexer) *ExpiresParser {
	this := &ExpiresParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

/** Parse the header.
 */
func (this *ExpiresParser) Parse() (sh header.Header, ParseException error) {
	expires := header.NewExpires()
	// if (debug) dbg_enter("parse");
	//        try {
	lexer := this.GetLexer()
	lexer.Match(TokenTypes_EXPIRES)
	lexer.SPorHT()
	lexer.Match(':')
	lexer.SPorHT()
	nextId := lexer.GetNextId()
	lexer.Match('\n')
	//try {
	delta, ParseException := strconv.ParseInt(nextId, 10, 32)
	expires.SetExpires(int(delta))
	return expires, ParseException
	//           } catch (NumberFormatException ex) {
	// throw createParseException("bad integer format");
	//    } catch (InvalidArgumentException ex) {
	// throw createParseException(ex.getMessage());
	//    }
	//       } finally  {
	// if (debug) dbg_leave("parse");
	//       }

}