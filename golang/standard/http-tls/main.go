package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

// 2个文件pulic_key.pem和private_key可用于验证证书

/*
tr := &http.Transport {
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true
	},
}

client := &http.Client{Transport: tr}
*/

func main() {
	caCert, err := ioutil.ReadFile("rootCA.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := http.DefaultClient
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
		},
	}
	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{
	// 			RootCAs: caCertPool,
	// 		},
	// 	},
	// }

	_, err = client.Get("https://secure.domain.com")
	if err != nil {
		panic(err)
	}
}

/*
1. Gen CA
	openssl genrsa -out rootCA.key 4096
	openssl req -x509 -new -key rootCA.key -days 3650 -out rootCA.crt

2. Generate certificate for secure.domain.com signed with created CA
	openssl genrsa -out secure.domain.com.key 2048
	openssl req -new -key secure.domain.com.key -out secure.domain.com.csr

	#In answer to question `Common Name (e.g. server FQDN or YOUR name) []:
	#` you should set `secure.domain.com` (your real domain name)
	openssl x509 -req -in secure.domain.com.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -days 365 -out secure.domain.com.crt
*/
