package goqrz

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetSession takes user/pass and returns a valid session
func GetSession(user string, pass string, agent string) (Session, error) {

	xmlBytes, err := getXML(
		fmt.Sprintf("http://xmldata.qrz.com/xml/1.33/?username=%s;password=%s;agent=%s",
			user, pass, agent))
	if err != nil {
		return Session{}, err
	}

	var database QRZDatabase
	xml.Unmarshal(xmlBytes, &database)

	newSession := database.Session
	newSession.user = user
	newSession.pass = pass

	if newSession.Error != "" {
		return newSession, errors.New(newSession.Error)
	}

	return newSession, nil
}

// GetCallsign takes a callsign and returns the QRZ information on that
// callsign.
func (CurrentSession *Session) GetCallsign(callsign string) (Callsign, error) {

	xmlBytes, err := getXML(
		fmt.Sprintf("http://xmldata.qrz.com/xml/1.33/?s=%s;callsign=%s",
			CurrentSession.Key, callsign))
	if err != nil {
		return Callsign{}, err
	}

	var database QRZDatabase
	xml.Unmarshal(xmlBytes, &database)

	if database.Session.Error != "" {
		return Callsign{}, errors.New(database.Session.Error)
	}
	return database.Callsign, nil
}

// GetDXCC takes a dxcc id and returns the QRZ information on that
// region.
func (CurrentSession *Session) GetDXCC(dxcc string) (DXCC, error) {

	if dxcc == "all" {
		return DXCC{}, errors.New("get all DXCC not implemented in this function")
	}

	xmlBytes, err := getXML(
		fmt.Sprintf("http://xmldata.qrz.com/xml/1.33/?s=%s;dxcc=%s",
			CurrentSession.Key, dxcc))
	if err != nil {
		return DXCC{}, err
	}

	var database QRZDatabase
	xml.Unmarshal(xmlBytes, &database)

	if database.Session.Error != "" {
		return DXCC{}, errors.New(database.Session.Error)
	}

	return database.DXCC, nil
}

// from: https://gist.github.com/james2doyle/e2f05b5756e4ee46848a8d987405f152
func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
