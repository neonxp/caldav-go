package ixml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/neonxp/caldav-go/lib"
)

const (
	DAV_NS     = "DAV:"
	CALDAV_NS  = "urn:ietf:params:xml:ns:caldav"
	CALSERV_NS = "http://calendarserver.org/ns/"
)

var NS_PREFIXES = map[string]string{
	DAV_NS:     "D",
	CALDAV_NS:  "C",
	CALSERV_NS: "CS",
}

var (
	CALENDAR_TG                         = xml.Name{Space: CALDAV_NS, Local: "calendar"}
	CALENDAR_DATA_TG                    = xml.Name{Space: CALDAV_NS, Local: "calendar-data"}
	CALENDAR_HOME_SET_TG                = xml.Name{Space: CALDAV_NS, Local: "calendar-home-set"}
	CALENDAR_QUERY_TG                   = xml.Name{Space: CALDAV_NS, Local: "calendar-query"}
	CALENDAR_MULTIGET_TG                = xml.Name{Space: CALDAV_NS, Local: "calendar-multiget"}
	CALENDAR_USER_ADDRESS_SET_TG        = xml.Name{Space: CALDAV_NS, Local: "calendar-user-address-set"}
	COLLECTION_TG                       = xml.Name{Space: DAV_NS, Local: "collection"}
	CURRENT_USER_PRINCIPAL_TG           = xml.Name{Space: DAV_NS, Local: "current-user-principal"}
	DISPLAY_NAME_TG                     = xml.Name{Space: DAV_NS, Local: "displayname"}
	GET_CONTENT_LENGTH_TG               = xml.Name{Space: DAV_NS, Local: "getcontentlength"}
	GET_CONTENT_TYPE_TG                 = xml.Name{Space: DAV_NS, Local: "getcontenttype"}
	GET_CTAG_TG                         = xml.Name{Space: CALSERV_NS, Local: "getctag"}
	GET_ETAG_TG                         = xml.Name{Space: DAV_NS, Local: "getetag"}
	GET_LAST_MODIFIED_TG                = xml.Name{Space: DAV_NS, Local: "getlastmodified"}
	HREF_TG                             = xml.Name{Space: DAV_NS, Local: "href"}
	OWNER_TG                            = xml.Name{Space: DAV_NS, Local: "owner"}
	PRINCIPAL_TG                        = xml.Name{Space: DAV_NS, Local: "principal"}
	PRINCIPAL_COLLECTION_SET_TG         = xml.Name{Space: DAV_NS, Local: "principal-collection-set"}
	PRINCIPAL_URL_TG                    = xml.Name{Space: DAV_NS, Local: "principal-URL"}
	RESOURCE_TYPE_TG                    = xml.Name{Space: DAV_NS, Local: "resourcetype"}
	STATUS_TG                           = xml.Name{Space: DAV_NS, Local: "status"}
	SUPPORTED_CALENDAR_COMPONENT_SET_TG = xml.Name{Space: CALDAV_NS, Local: "supported-calendar-component-set"}
)

// Namespaces returns the default XML namespaces in for CalDAV contents.
func Namespaces() string {
	bf := new(lib.StringBuffer)
	bf.Write(`xmlns:%s="%s" `, NS_PREFIXES[DAV_NS], DAV_NS)
	bf.Write(`xmlns:%s="%s" `, NS_PREFIXES[CALDAV_NS], CALDAV_NS)
	bf.Write(`xmlns:%s="%s"`, NS_PREFIXES[CALSERV_NS], CALSERV_NS)

	return bf.String()
}

// Tag returns a XML tag as string based on the given tag name and content. It
// takes in consideration the namespace and also if it is an empty content or not.
func Tag(xmlName xml.Name, content string) string {
	name := xmlName.Local
	ns := NS_PREFIXES[xmlName.Space]

	if ns != "" {
		ns = ns + ":"
	}

	if content != "" {
		return fmt.Sprintf("<%s%s>%s</%s%s>", ns, name, content, ns, name)
	} else {
		return fmt.Sprintf("<%s%s/>", ns, name)
	}
}

// HrefTag returns a DAV <D:href> tag with the given href path.
func HrefTag(href string) (tag string) {
	return Tag(HREF_TG, href)
}

// StatusTag returns a DAV <D:status> tag with the given HTTP status. The
// status is translated into a label, e.g.: HTTP/1.1 404 NotFound.
func StatusTag(status int) string {
	statusText := fmt.Sprintf("HTTP/1.1 %d %s", status, http.StatusText(status))
	return Tag(STATUS_TG, statusText)
}

// EscapeText escapes any special character in the given text and returns the result.
func EscapeText(text string) string {
	buffer := bytes.NewBufferString("")
	xml.EscapeText(buffer, []byte(text))

	return buffer.String()
}
