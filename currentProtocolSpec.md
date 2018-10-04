# The QRZ XML Interface Specification

![QRZ Logo](https://www.qrz.com/gifs/qrzcom200.gif)

|          |                                    |
| -------- | ---------------------------------- |
| Title:   | QRZ XML Logbook Data Specification |
| Version: | 1.33                               |
| Date:    | September 18, 2014                 |
| Author:  | Fred Lloyd, AA7BQ                  |
| Email:   | flloyd@qrz.com                     |

## What's New in Version 1.33

- **Bug Fix Release** - A bug fix to make sure that the image URL points to the same picture as shown on the user's QRZ web page. Also, this document has been revised to indicate that the correct service URL for the QRZ server is **xmldata.qrz.com**

## Overview

This document describes the interface specification for access to QRZ's XML subscription data service. This service provides real-time access to information from the QRZ.COM servers and databases.

Access to this service requires user authentication through the use of a valid QRZ.COM username and password. While any QRZ user may login to the service, an active QRZ Logbook Data subscription is required to access most of its features. Non-subscriber access limits the data fields that are returned and is primarily intended for testing and troubleshooting purposes only.

A description of subscription plans and rates is available [on the QRZ website](http://www.qrz.com/i/subscriptions.html).

## Access Methodology

QRZ's XML service is implemented using standard HTTP protocols. Requests are made as query string arguments to a standard HTTP URL. Results are returned over HTTP as XML formatted data.

All post-login interactions with the server require that a dynamic **session key** be included with each request. The session key is sent by specifying the s= parameter. Note that requests to the XML service port may utilize either HTTP GET or HTTP POST methods.

Session keys are dynamically managed by the server and have no guaranteed lifetime. When a session key expires or becomes invalidated, a new key may be issued upon request. Client programs should cache all session keys provided by the server and reuse them until they expire. This practice maximizes client performance and serves to minimize the load on QRZ's servers.

As a general programming pattern, clients should expect to perform only one login operation per session. The session key returned from a successful login should then be locally cached and reused as many times as possible. All clients should monitor the status response returned in each transaction and be prepared to login again whenever indicated.

A session key is valid only for a single user and may become immediately invalidated if it is detected that the user's IP address or other identifying information changes after login has been completed. The interface does not use HTTP cookies, Javascript, or HTML.

## QRZ's XML Service Address

The QRZ XML service is located at `http://xmldata.qrz.com/xml/`. If greater security is desired to protect login credentials, HTTPS may be used as `https://xmldata.qrz.com/xml/` with an HTTP POST method. The session key that is returned from an HTTPS login is valid on both the HTTP and HTTPS interfaces. We request that you limit your HTTPS interactions to login (session establishment) operations and not use HTTPS for routine callsign lookups. Session keys are cryptographically secure and non-transferrable and thus do not need to be protected or wrapped with additional encryption protocols.

## Versioned URL Interface

Beginning with version **1.25**, access to QRZ's XML service is versioned. Versioning provides the ability for QRZ to periodically extend or make significant changes to this specification without sacrificing client compatibility with earlier releases.

This version introduces an extension to the service URL that indicates which version of the interface should be used to service the request. The general format of the service URL is:

**http://xmldata.qrz.com/xml/*version_identifier*/?\<query parameters\>**

The version_identifier indicates which version of the interface will be used.
Valid examples are:

### Version Identifier Values

| Identifier | use                                           |
| ---------- | --------------------------------------------- |
| (none)     | Missing identifier: use version 1.24 (legacy) |
| 1.xx       | Use version 1.xx                              |
| current    | Use the latest available version.             |

#### Note:

If no version identifier is present, the interface specification follows the legacy 1.24 specification to maintain compatibility with existing client programs.

### Examples:

1. http://xmldata.qrz.com/xml/?username=xx1xxx;password=abcdef
2. http://xmldata.qrz.com/xml/1.31/?username=xx1xxx;password=abcdef
3. http://xmldata.qrz.com/xml/current/?username=xx1xxx;password=abcdef

In the first example, no version identifier was specified and so the lowest version (1.24) will be used. In the second example, the version **1.31** interface will be used. In the third example, the most current version avaialable will be used. Note that the trailing slash (/) character following either the 'xml' or the 'version' identifier is highly recommended. Also, the dot (.) in the version identifier may be omitted, e.g. either '1.31' or '131' may be used.

If a version identifier is used, it must be either 'current' or a supported release number. The use of any other value will result in an error.

## XML Data Structure

The QRZ XML data structure follows the standard proposed by WWW.W3C.ORG. QRZ's top-level node is the `<QRZDatabase>` element. All responses from the QRZ server are prefaced with this element. Three child node types are defined for this level. They are:

- `<Session>`
- `<Callsign>`
- `<DXCC>`

Each of these nodes is described below. An important design pattern must be observed in order for your client program to maintain its highest level of reliability:

*The XML data supplied by this interface may be extended at any time in a forward compatible manner. QRZ may add new XML nodes and/or attributes at any time. Therefore, it is critically important that client programs carefully parse the data received in an "object=attribute" manner while ignoring any unknown nodes or attributes, and do so without raising an error condition. Your decoding strategy must make no assumptions regarding the number of nodes returned, the order in which nodes are returned, or the presence of nodes or attributes not previously defined.*

## The Session Node

The Session node contains information pertaining to the status of the user's session, error codes and informational messages relating to the request.

### Setting up a Session

To perform a simple login and obtain a session key, the client program sends the username and password using either an HTTP GET or POST operation. A typical login exchange looks like this:

```
http://xmldata.qrz.com/xml/current/?username=xx1xxx;password=abcdef;agent=q5.0
```

The URI may use either the ampersand (&) or the semicolon (;) as the parameter separator.

### Session Input Fields:

- username - a valid QRZ user name
- password - the correct password for the username
- agent - a string that contains the product name and version of the client program

The required parameters are **username** and **password**. The use of the **agent** parameter is strongly recommended as it assists in troubleshooting and end user support. You may use any agent identifier you wish, however, it will be most useful if you maintain a unique prefix for your product (for example, Ham Radio Deluxe uses 'HRDv.x'). Please keep it short and simple.

Note: if the **agent** parameter is not used, the XML server will attempt to use the HTTP_USER_AGENT string obtained from the HTTP protocol.

If the username and password are accepted, the system will respond with:

```xml
<?xml version="1.0" ?>
<QRZDatabase version="1.33">
  <Session>
    <Key>2331uf894c4bd29f3923f3bacf02c532d7bd9</Key>
    <Count>123</Count>
    <SubExp>Wed Jan 1 12:34:03 2013</SubExp>
    <GMTime>Sun Aug 16 03:51:47 2012</GMTime>
  </Session>
</QRZDatabase>
```

As with all server responses, the return begins with **QRZDatabase**. Note that the QRZDatabase node also contains two attributes, `version` and `xmlns`. The `version` attribute represents the QRZ XML version currently in use as defined by the release number of this specification. The `xmlns` attribute identifies the XML namespace of this product.

The possible elements in a session node are:

#### Session section fields

| Field | Description |
| - | - |
| Key | a valid user session key |
| Count | Number of lookups performed by this user in the current 24 hour period |
| SubExp | time and date that the users subscription will expire - or - "non-subscriber" |
| GMTime | Time stamp for this message |
| Message | An informational message for the user |
| Error | XML system error message |

Both `Error` and `Message` responses should be presented to the end user. Typical `Error` responses include things such as "password incorrect", "session timeout", and "callsign not found". Typical `Message` responses return notices like: "A subscription is required to obtain the complete data".

A user session is established whenever a session key is returned. Any response from the server that does not contain the `Key` element indicates that no valid session exists and that a re-login is required to continue.

The `Count`, `SubExp`, and `GMTime` nodes are simply informational and may be shown to the end user if deemed to be useful.

## Callsign Lookups

To make a callsign query, simply pass the current session key in the `s=` parameter, followed by a `callsign=` parameter.

A typical request might look like this:

```
http://xmldata.qrz.com/xml/current/?s=f894c4bd29f3923f3bacf02c532d7bd9;callsign=aa7bq
```

This returns the following data:

```xml
<?xml version="1.0" ?> 
<QRZDatabase version="1.33">
  <Callsign>
      <call>AA7BQ</call> 
      <aliases>N6UFT,KJ6RK,DL/AA7BQ</aliases> 
      <dxcc>291</dxcc> 
      <fname>FRED L</fname> 
      <name>LLOYD</name> 
      <addr1>8711 E PINNACLE PEAK RD 193</addr1> 
      <addr2>SCOTTSDALE</addr2> 
      <state>AZ</state> 
      <zip>85255</zip> 
      <country>United States</country> 
      <ccode>291</ccode> 
      <lat>34.23456</lat> 
      <lon>-112.34356</lon> 
      <grid>DM32af</grid> 
      <county>Maricopa</county> 
      <fips>04013</fips> 
      <land>USA</land> 
      <efdate>2000-01-20</efdate> 
      <expdate>2010-01-20</expdate> 
      <p_call>KJ6RK</p_call> 
      <class>E</class> 
      <codes>HAI</codes> 
      <qslmgr>NONE</qslmgr> 
      <email>flloyd@qrz.com</email> 
      <url>http://www.qrz.com/db/aa7bq</url> 
      <u_views>115336</u_views> 
      <bio>3937/2003-11-04</bio> 
      <image>http://files.qrz.com/q/aa7bq/aa7bq.jpg</image> 
      <serial>3626</serial> 
      <moddate>2003-11-04 19:37:02</moddate> 
      <MSA>6200</MSA> 
      <AreaCode>602</AreaCode> 
      <TimeZone>Mountain</TimeZone> 
      <GMTOffset>-7</GMTOffset> 
      <DST>N</DST> 
      <eqsl>Y</eqsl> 
      <mqsl>Y</mqsl> 
      <cqzone>3</cqzone> 
      <ituzone>2</ituzone> 
      <geoloc>user</geoloc> 
      <born>1953</born> 
  </Callsign>
  <Session>
      <Key>2331uf894c4bd29f3923f3bacf02c532d7bd9</Key> 
      <Count>123</Count> 
      <SubExp>Wed Jan 1 12:34:03 2013</SubExp> 
      <GMTime>Sun Nov 16 04:13:46 2012</GMTime> 
  </Session>
</QRZDatabase>
```

The possible callsign data fields are listed below. Not all fields may be returned with each request. The field ordering is arbitrary and subject to change.

### Callsign node fields

| Field | Description |
| ----- | ----------- |
| call | callsign |
| xref | Cross reference: the query callsign that returned this record |
| aliases | Other callsigns that resolve to this record |
| dxcc | DXCC entity ID (country code) for the callsign |
| fname | first name |
| name | last name |
| addr1 | address line 1 (i.e. house # and street) |
| addr2 | address line 2 (i.e, city name) |
| state | state (USA Only) |
| zip | Zip/postal code |
| country | country name for the QSL mailing address |
| ccode | dxcc entity code for the mailing address country |
| lat | lattitude of address **(signed decimal) S < 0 > N** |
| lon | longitude of address **(signed decimal) W < 0 > E** |
| grid | grid locator |
| county | county name (USA) |
| fips | FIPS county identifier (USA) |
| land | DXCC country name of the callsign |
| efdate | license effective date (USA) |
| expdate | license expiration date (USA) |
| p_call | previous callsign |
| class | license class |
| codes | license type codes (USA) |
| qslmgr | QSL manager info |
| email | email address |
| url | web page address |
| u_views | QRZ web page views |
| bio | approximate length of the bio HTML in bytes |
| biodate | date of the last bio update |
| image | full URL of the callsign's primary image |
| imageinfo | height:width:size in bytes, of the image file |
| serial | QRZ db serial number |
| moddate | QRZ callsign last modified date |
| MSA | Metro Service Area (USPS) |
| AreaCode | Telephone Area Code (USA) |
| TimeZone | Time Zone (USA) |
| GMTOffset | GMT Time Offset |
| DST | Daylight Saving Time Observed |
| eqsl | Will accept e-qsl (0/1 or blank if unknown) |
| mqsl | Will return paper QSL (0/1 or blank if unknown) |
| cqzone | CQ Zone identifier |
| ituzone | ITU Zone identifier |
| born | operator's year of birth |
| user | User who manages this callsign on QRZ |
| lotw | Will accept LOTW (0/1 or blank if unknown) |
| iota | IOTA Designator (blank if unknown) |
| geoloc | Describes source of lat/long data |

## About Lat/Long and Grid Data

The lat/long and grid data returned varies by the method by which it was obtained by QRZ. Whenever a lat/long coordinate pair exists for the queried callsign, its value is used to calculate the appropriate Maidenhead locator (grid) value. In some rare cases, a given QRZ record may be missing its lat/long value but has a user supplied grid locator. When this is the case, the station's lat/long values are stated as the center the given grid locator.

In most cases, when neither a lat/long pair or grid is known for a USA record, it will be derived from its postal address. For USA callsigns, the lat/long pair is determined using a geocoding procedure derived from the US Census Tiger/Line dataset. Not all addresses can be successfully geocoded and subsequently a Zip Code database is used to return an approximate location that roughly centers on the zip code. Failing all of the above, for USA callsigns, the geographic coordinates of the state's approximate center are used.

For non-USA callsigns, no geocoding is available so unless the user has specifically input lat/long or grid coordinates, the approximate center of the DXCC entity (country) is used.

The net result is that nearly every callsign returned by QRZ will include geographic coordinates. While in some cases they may not be exact, they are generally close enough for DX antenna positioning operations.

The `geoloc` field describes the source of the returned lat/long data. The possible string values for `geoloc` are:

- `user` - the value was input by the user
- `geocode` - the value was derived from USA Geocoding data
- `grid` - the value was derived from a user supplied grid square
- `zip` - the value was derived from the callsign's USA Zip Code
- `state` - the value was derived from the callsign's USA State
- `dxcc` - the value was derived from the callsign's DXCC entity (country)
- `none` - no value could be determined

## Biography Data

The `<bio>` field is present in the callsign record **if and only if** a biography record exists for the target callsign on the server. This field indicates the approximate size of the bio HTML data and the `<biodate>` field indicates the date of last bio update.

To fetch the bio HTML, issue a separate request using the `html=` query parameter for the same callsign.

### Bio Fetch Example:

```
http://xmldata.qrz.com/xml/current/?s=d0cf9d7b3b937ed5f5de28ddf5a0122d;html=xx1xxx
```

The `html=callsign` method is unique in this interface in that it **does not return XML**, but rather regular HTML, including embedded CSS, just as the information would normally appear on the callsign's QRZ page.

## DXCC / Prefix Lookups

Support for prefix matching and DXCC entity lookups is also provided by this interface. Access to this method is via the `dxcc=` url parameter.

The dxcc lookup provides three functions:

1. Look up a DXCC entity by its code,
2. Determine the matching DXCC entity for a given callsign
3. Return the entire list of QRZ DXCC entities.

Functions are accessed using variations of the `dxcc=` input paramter. The allowed parameters are:

- `dxcc=123` (a number) - return up the given DXCC entity.
- `dxcc=xx1xx` (a callsign) - return the computed DXCC entity.
- `dxcc=all` (keyword 'all') - return all the QRZ DXCC entities.

In the first form, `dxcc=123`, the server returns the QRZ DXCC record for the given entity. An entity lookup is inferred whenever the input argument is all numeric. This entity corelates to the key given in both the `dxcc` and the `ccode` fields returned in callsign listings.

In the second form, `dxcc=xx1xx`, the server reduces the callsign to a 4, then a 3, then a 2-letter prefix and returns the first DXCC entity that matches. This can be useful to determine the country of a random callsign.

In the third form, `dxcc=all`, the server returns its entire list of 380+ DXCC records. Please use this option sparingly so as not to overburden the server.

A typical dxcc fetch for entity 291 (USA):

```
http://xmldata.qrz.com/xml/current/?s=d0cf9d7b3b937ed5f5de28ddf5a0122d;dxcc=291
```

Resulting in:

```xml
<?xml version="1.0" ?> 
<QRZDatabase version="1.33">
<DXCC>
<dxcc>291</dxcc>
<cc>US</cc>
<ccc>USA</ccc>
<name>United States</name>
<continent>NA</continent>
<ituzone>6</ituzone>
<cqzone>3</cqzone>
<timezone>-5</timezone>
<lat>37.788081</lat>
<lon>-97.470703</lon>
</DXCC>
<Session>
<Key>d0cf9d7b3b937ed5f5de28ddf5a0122d</Key>
<Count>12</Count>
<SubExp>Wed Jan 13 13:59:00 2013</SubExp>
<GMTime>Mon Oct 12 22:33:56 2012</GMTime>
</Session>
</QRZDatabase>
```

### DXCC node fields

| Field | Description |
| - | - |
| dxcc | DXCC entity number for this record |
| cc | 2-letter country code (ISO-3166) |
| cc | 3-letter country code (ISO-3166) |
| name | long name |
| continent | 2-letter continent designator |
| ituzone | ITU Zone |
| cqzone | CQ Zone |
| timezone | UTC timezone offset +/- |
| lat | Latitude (approx.) |
| lon | Longitude (approx.) |
| notes | Special notes and/or exceptions |

Notes:

- QRZ's DXCC list includes some entities that are unique to QRZ. Such entities have identifiers that are greater than 900.

- Timezone values are notated in hours, +/- UTC. Odd timezones, such as `0545` mean "5 hours, 45 minutes". The plus (+) sign is implied.

- Lat and Lon values are usually located at the approximate geographic center of the entity.

- Lookups other than **'all'** return `No DXCC information for: xxxx` on non-match failures.

## Error Conditions

There are two general types of errors. Data errors, which are typically of the form "item not found", and Session errors, which deal with the user's session key. If the <Session> node contains an <Error> sub-node, then the message should be examined and/or presented to the user.

**Here's an example of a typical "callsign not found" error:**

```xml
<?xml version="1.0" ?> 
<QRZDatabase version="1.33">
    <Session>
    <Error>Not found: g1srdd</Error> 
    <Key>1232u4eaf13b8336d61982c1fd1099c9a38ac</Key> 
    <GMTime>Sun Nov 16 05:07:14 2003</GMTime> 
    </Session>
</QRZDatabase>
```

Should a session expire or become invalidated, the **`<Key>`** field will not be sent.

Here's an example of a "Session" error:

```xml
<?xml version="1.0" ?> 
<QRZDatabase version="1.33">
    <Session>
    <Error>Session Timeout</Error> 
    <GMTime>Sun Nov 16 05:11:58 2003</GMTime> 
    </Session>
</QRZDatabase>
```

A special `Error` message, `Connection refused` is significant in that it indicates that service is refused for the given user. No further explanation is given, however, it does indicate that successful login will not be possible for at least 24 hours.

## Further Information

For questions concerning this document or the XML interface, please contact the author at flloyd@qrz.com .

## Revision History:

| Date | Revision |
| - | - |
| 1.0 Sat Nov 15, 2003 | original draft |
| 1.1 Mon Feb 28, 2005 | update to reflect correct image path |
| 1.2 Wed Mar 2, 2005 | image tag now contains URL |
| 1.3 Thu Jun 22, 2006 | reworded and added Alert tag and extensions policy |
| 1.4 Thu Jun 23, 2006 | documented the "agent=" parameter |
| 1.5 Wed Dec 3, 2008 | some new data fields, field description section |
| 1.6 Wed Jan 15, 2009 | rename site from online.qrz.com to xml.qrz.com |
| 1.7 Wed Jan 15, 2009 | new bio fetch procedure |
| 1.8 Fri Apr 10, 2009 | field list (grid), session key description, Expires removed |
| 1.9 Thu Sep 10, 2009 | numerous incl. QRZDatabase version attribute |
| 1.10 Sat Sep 12, 2009 | some typographical corrections |
| 1.11 Sun Sep 13, 2009 | add emphasis to the Extensions paragraph |
| 1.12 Thu Sep 17, 2009 | rewrite - new Message node, new explanations. |
| 1.13 Sat Oct 3, 2009 | revised support for eqsl,mqsl,lotw and iota fields |
| 1.14 Fri Oct 9, 2009 | update for dxcc entity field |
| 1.15 Mon Oct 12, 2009 | added new dxcc / prefix search, added ccode field |
| 1.16 Tue Jun 8, 2010 | corrected documentation to reflect returned xmlns behavior |
| 1.17 Tue Jan 26, 2011 | corrected documentation to reflect new bio/image results |
| 1.18 Thu Mar 24, 2011 | corrected documentation to reflect new bio/image results |
| 1.19 Wed Nov 9, 2011 | documentation update for the 'agent' parameter |
| 1.20 Wed Nov 14, 2011 | internal update, added CSS data to HTML request, no interface changes |
| 1.25 Wed Jun 20, 2012 | Versioned interface, cross references, aliases, https documentation, new bio fetch procedure |
| 1.26 Mon Jul 23, 2012 | Revised lat/long and grid calculations to match www.qrz.com |
| 1.27 Wed Jul 25, 2012 | Replaced the 'locref' field with new 'geoloc' field to describe source of lat/long coordinates. |
| 1.28 date not specified | Internal update |
| 1.29 date not specified | Internal update |
| 1.30 Tue Feb 11, 2014 | Bug Fix to restore biography data and image results. |
| 1.31 Tue Feb 25, 2014 | Bug Fix to biography image URL's |
| 1.32 Tue June 24, 2014 | Bug Fix to fetch latest biography version |
| 1.33 Thu Sep 18, 2014 | Bug Fix to fetch latest biography primary image (picture) plus Doc Update |

<hr>

Copyright © 2003-2014 by QRZ.COM
