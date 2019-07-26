// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

import (
	"crypto/x509"
	"net/url"

	"golang.org/x/oauth2"
)

type Domain struct {
	Name         string // human readable representation of the domain/website
	LoginURL     string
	Certificate  *x509.Certificate
	OauthEnabled bool // can we use oauth2 to log in?
}

// Crerdentials is the global map that stores all logins
var Credentials = make(CredMap)

// CredMap is a map that ties user credentials to thier respective website.
type CredMap map[Domain]Credential

// NewDomainFromURL creates a Domain key from a webpages main url (such as www.google.com).
// It does its best to strip sub-domains if present, and extrenious info such as query strings,
// If it fails in this process it will return error of type UrlParseError.
func NewDomainFromURL(urlStr string) (*Domain, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	cert, err := GetCertFromURL(urlStr)
	if err != nil {
		return nil, err
	}
	d := &Domain{
		Name:        u.Host,
		Certificate: cert,
	}
	return d, nil

}

func CreateOAuth2(username string) error {
	config := new(oauth2.Config)
}
