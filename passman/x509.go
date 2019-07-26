// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net/http"
)

func GetCertFromURL(url string) (*x509.Certificate, error) {
	/*
		resp, err := http.Get(loginURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	*/
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	var client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	peerCert := resp.TLS.PeerCertificates[0]
	if peerCert.PublicKey == nil {
		return nil, errors.New("remote tls cert has nil public key")
	}
	return peerCert, nil
}
