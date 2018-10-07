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

## GoQRZ Command Line Interface

Included in this library is a CLI tool which can be used to query QRZ for information on callsigns and DXCC entities.

### Install GoQRZ CLI

There are two ways to install and use the CLI; compiling from source and downloading precompiled binaries.

#### From Source

The CLI can be compiled from source by running the following:

```shell
go get github.com/ocelotsloth/goqrz/goqrz
```

Provided your `GOPATH` is configured correctly you should now be able to run `goqrz` from the command line.

#### Precompiled Binaries

Look at the [releases](https://github.com/ocelotsloth/goqrz/releases) on this GitHub page for downloadable binaries. This utility can be cross-compiled to most any popular archetecture in use. Open an issue if you need an additional archetecture added to the list of precompiled binaries.

### CLI Usage

Before pulling data from the XML API, it is important to log in. **Ensure 2 factor authentication is enabled on your account before using this utility!** QRZ.com does not have encryption enabled on their XML endpoint, so it is imparative that this password be unique to QRZ.com. Be sure to change it frequently as well. This library does its best to remain secure by not storing your username or password between calls. Instead it provides two methods to store the session key.

#### Get Session Key

First, run `goqrz login -u <username> -p <password>` to receive your session key. The key is printed to stdout, so you can store it as an environment variable (store it to `GOQRZ_KEY`) or by passing the key to each subsequent call via the `--key` flag.

#### Query Data

There are two queries currently implemented: callsigns and dxcc zones.

##### Callsigns

To retrieve callsign data, use the command `goqrz callsign [--key=<sessionKey>] Callsign [AdditionalCallsigns...]`. To be efficient, please consider batching your callsign requests to one single call to the `goqrz` CLI. This will reduce the overhead with setting up the connection to QRZ.com.

Data is returned as JSON. I personally recommend using `jq` to further deal with the data.

##### DXCC Zones

DXCC data can be retrieved in the exact same fasion as Callsigns, except with the following syntax: `goqrz dxcc [--key=<sessionKey>] DXCCID [AdditionalDXCCIDs...]`
