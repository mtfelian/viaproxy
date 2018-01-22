[![Go Report Card](https://goreportcard.com/badge/github.com/mtfelian/viaproxy)](https://goreportcard.com/report/github.com/mtfelian/viaproxy)
[![GoDoc](https://godoc.org/github.com/mtfelian/viaproxy?status.png)](http://godoc.org/github.com/mtfelian/viaproxy)
[![Build status](https://travis-ci.org/mtfelian/viaproxy.svg?branch=master)](https://travis-ci.org/mtfelian/viaproxy)

# viaproxy

**viaproxy** is a proxy connection library.

It provides the interface `Connector` and two implementations:
`viaSOCKS5` and `viaNoProxy`.

See the usage example below.

```
	var proxyConnector viaproxy.Connector
	validProxyAddresses := validProxies()
	if len(validProxyAddresses) == 0 {
		proxyConnector = viaproxy.NewViaNoProxy()
	} else {
		proxyConnector = viaproxy.NewViaSOCKS5()
		proxyConnector.AddProxyAddr(validProxyAddresses...)
	}

	entity.Parse(specific.New(proxyConnector))

	...

	response, err := p.Connector().DoRequest(http.MethodGet, url, nil, 2*time.Second)
```