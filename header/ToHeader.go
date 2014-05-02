/**
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 * Module Name   : GoSIP Specification
 * File Name     : ToHeader.go
 * Author        : Rain Liu
 *~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */

package header

/**
 * The To header field first and foremost specifies the desired "logical"
 * recipient of the request, or the address-of-record of the user or resource
 * that is the target of this request.  This may or may not be the ultimate
 * recipient of the request. Requests and Responses must contain a ToHeader,
 * indicating the desired recipient of the Request. The UAS or redirect server
 * copies the ToHeader into its Response.
 * <p>
 * The To header field MAY contain a SIP or SIPS URI, but it may also make use
 * of other URI schemes i.e the telURL, when appropriate. All SIP
 * implementations MUST support the SIP URI scheme.  Any implementation that
 * supports TLS MUST support the SIPS URI scheme. Like the From header field,
 * it contains a URI and optionally a display name, encapsulated in a
 * {@link javax.sip.address.Address}.
 * <p>
 * A UAC may learn how to populate the To header field for a particular request
 * in a number of ways.  Usually the user will suggest the To header field
 * through a human interface, perhaps inputting the URI manually or selecting
 * it from some sort of address book.  Using the string to form the user part
 * of a SIP URI implies that the User Agent wishes the name to be resolved in the
 * domain to the right-hand side (RHS) of the at-sign in the SIP URI.  Using
 * the string to form the user part of a SIPS URI implies that the User Agent wishes to
 * communicate securely, and that the name is to be resolved in the domain to
 * the RHS of the at-sign. The RHS will frequently be the home domain of the
 * requestor, which allows for the home domain to process the outgoing request.
 * This is useful for features like "speed dial" that require interpretation of
 * the user part in the home domain.
 * <p>
 * The telURL may be used when the User Agent does not wish to specify the domain that
 * should interpret a telephone number that has been input by the user. Rather,
 * each domain through which the request passes would be given that opportunity.
 * As an example, a user in an airport might log in and send requests through
 * an outbound proxy in the airport.  If they enter "411" (this is the phone
 * number for local directory assistance in the United States), that needs to
 * be interpreted and processed by the outbound proxy in the airport, not the
 * user's home domain.  In this case, tel:411 would be the right choice.
 * <p>
 * Two To header fields are equivalent if their URIs match, and their
 * parameters match. Extension parameters in one header field, not present in
 * the other are ignored for the purposes of comparison. This means that the
 * display name and presence or absence of angle brackets do not affect
 * matching.
 * <ul>
 * <li> The "Tag" parameter - is used in the To and From header fields of SIP
 * messages.  It serves as a general mechanism to identify a dialog, which is
 * the combination of the Call-ID along with two tags, one from each
 * participant in the dialog.  When a UA sends a request outside of a dialog,
 * it contains a From tag only, providing "half" of the dialog ID. The dialog
 * is completed from the response(s), each of which contributes the second half
 * in the To header field. When a tag is generated by a UA for insertion into
 * a request or response, it MUST be globally unique and cryptographically
 * random with at least 32 bits of randomness. Besides the requirement for
 * global uniqueness, the algorithm for generating a tag is implementation
 * specific.  Tags are helpful in fault tolerant systems, where a dialog is to
 * be recovered on an alternate server after a failure.  A UAS can select the
 * tag in such a way that a backup can recognize a request as part of a dialog
 * on the failed server, and therefore determine that it should attempt to
 * recover the dialog and any other state associated with it.
 * </ul>
 * A request outside of a dialog MUST NOT contain a To tag; the tag in the To
 * field of a request identifies the peer of the dialog.  Since no dialog is
 * established, no tag is present.
 * <p>
 * For Example:<br>
 * <code>To: Carol sip:carol@jcp.org<br>
 * To: Duke sip:duke@jcp.org;tag=287447</code>
 *
 * @see AddressHeader
 * @version 1.1
 * @author Sun Microsystems
 */
type ToHeader interface {
	//Header
	AddressHeader
	ParametersHeader

	/**
	 * Sets the tag parameter of the ToHeader. The tag in the To field of a
	 * request identifies the peer of the dialog. If no dialog is established,
	 * no tag is present.
	 * <p>
	 * The To Header MUST contain a new "tag" parameter. When acting as a UAC
	 * the To "tag" is maintained by the SipProvider from the dialog layer,
	 * however whan acting as a UAS the To "tag" is assigned by the application.
	 * That is the tag assignment for outbound responses for messages in a
	 * dialog is only the responsibility of the application for the first
	 * outbound response. After dialog establishment, the stack will take care
	 * of the tag assignment.
	 *
	 * @param tag - the new tag of the To Header
	 * @throws ParseException which signals that an error has been reached
	 * unexpectedly while parsing the Tag value.
	 */
	SetTag(tag string) (ParseException error)

	/**
	 * Gets tag of ToHeader. The Tag parameter identified the Peer of the
	 * dialogue.
	 *
	 * @return the tag parameter of the ToHeader. Returns null if no Tag is
	 * present, i.e no dialogue is established.
	 */
	GetTag() string

	/**
	 * Name of the ToHeader
	 */
	//public final static String NAME = "To";
}
