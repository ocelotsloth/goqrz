package goqrz

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetSessionKey takes user/pass and returns a valid session
func GetSessionKey(user string, pass string, agent string) (string, error) {

	xmlBytes, err := getXML(
		fmt.Sprintf("http://xmldata.qrz.com/xml/1.34/?username=%s;password=%s;agent=%s",
			user, pass, agent))
	if err != nil {
		return "", err
	}

	var database QRZDatabase
	xml.Unmarshal(xmlBytes, &database)

	newSession := database.Session

	if newSession.Error != "" {
		return "", errors.New(newSession.Error)
	}

	return newSession.Key, nil
}

// GetSession creates a session object from just a key
func GetSession(key string, agent string) (Session, error) {

	xmlBytes, err := getXML(fmt.Sprintf("http://xmldata.qrz.com/xml/1.34/?s=%s;agent=%s", key, agent))
	if err != nil {
		return Session{}, err
	}

	var database QRZDatabase
	xml.Unmarshal(xmlBytes, &database)

	if database.Session.Error != "" {
		return database.Session, errors.New(database.Session.Error)
	}

	return database.Session, nil
}

// GetCallsign takes a callsign and returns the QRZ information on that
// callsign.
func GetCallsign(key string, callsign string, agent string) (Callsign, error) {

	xmlBytes, err := getXML(fmt.Sprintf("http://xmldata.qrz.com/xml/1.34/?s=%s;callsign=%s;agent=%s", key, callsign, agent))
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
func GetDXCC(key string, dxcc string, agent string) (DXCC, error) {

	if dxcc == "all" {
		return DXCC{}, errors.New("get all DXCC not implemented in this function")
	}

	xmlBytes, err := getXML(fmt.Sprintf("http://xmldata.qrz.com/xml/1.34/?s=%s;dxcc=%s;agent=%s", key, dxcc, agent))
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
