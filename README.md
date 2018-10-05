# GoQRZ Library

GoQRZ is a simple Go Library which implements the QRZ.com [specification](https://www.qrz.com/XML/current_spec.html).

## Usage

First, import the library:

```go
import "github.com/ocelotsloth/goqrz"
```

Next, create a new session, providing a QRZ.com username and password as well as a User Agent which identifies your program to the API service:

```go
qrzSession := goqrz.GetSession("user", "pass", "userAgent")
```

Finally, request data using the `GetCallsign` or `GetDXCC` functions:

```go
qrzSession.GetCallsign("KN4IJZ")
```

Documentation for the returned datatypes can be found via godoc or within the source code.

## Future Plan

I plan to implement a simple cli interface which will allow for simple scripting of QRZ lookups from the command line. Look out for updates to this library.
