<h1 id="mocha-top" align="center">Mocha</h1>

<div align="center">
    <a href="#"><img src="docs/logo.png" width="120px" alt="Mocha Logo"></a>
    <p align="center">
        HTTP mocking in GO
        <br />
        <a href="https://github.com/vitorsalgado/mocha/actions/workflows/ci.yml"><strong>CI</strong></a> 
    </p>
    <div>
      <a href="https://github.com/vitorsalgado/mocha/actions/workflows/ci.yml">
        <img src="https://github.com/vitorsalgado/mocha/actions/workflows/ci.yml/badge.svg" alt="CI Status" />
      </a>
      <a href="https://github.com/vitorsalgado/mocha/actions/workflows/codeql-analysis.yml">
        <img src="https://github.com/vitorsalgado/mocha/actions/workflows/codeql-analysis.yml/badge.svg" alt="CodeQL Status" />
      </a>
      <a href="https://codecov.io/gh/vitorsalgado/mocha">
        <img src="https://codecov.io/gh/vitorsalgado/mocha/branch/main/graph/badge.svg?token=XOFUV52P31" alt="Coverage"/>
      </a>
      <a href="#">
        <img src="https://img.shields.io/badge/go-1.18-blue" alt="Go 1.18" />
      </a>
      <a href="https://conventionalcommits.org">
        <img src="https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg" alt="Conventional Commits"/>
      </a>
    </div>
</div>

## Overview

HTTP server mocking tool for Go.  
**Mocha** creates an real HTTP server and lets you configure response stubs for specific requests based on a set of
criterias. It provides a functional like API that allows you to match any part of a request against a set of matching
functions.

Inspired by [WireMock](https://github.com/wiremock/wiremock) and [Nock](https://github.com/nock/nock).

> Work In Progress

## Installation

```bash
go get -u github.com/vitorsalgado/mocha
```

## Features

- Configure HTTP response stubs for specific requests based on a criteria set.
- Matches request URL, headers, queries, body.
- Stateful matches to create scenarios, mocks for a specific number of calls.
- Response body template.
- Response delays.
- Run in your automated tests.

## Getting Started

```go
package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorsalgado/mocha"
	"github.com/vitorsalgado/mocha/matcher"
	"github.com/vitorsalgado/mocha/reply"
)

func Test(t *testing.T) {
	m := mocha.ForTest(t)
	m.Start()

	scoped := m.Mock(mocha.Get(matcher.URLPath("/test")).
		Header("test", matcher.EqualTo("hello")).
		Query("filter", matcher.EqualTo("all")).
		Reply(reply.Created().BodyString("hello world")))

	req, _ := http.NewRequest(http.MethodGet, m.Server.URL+"/test?filter=all", nil)
	req.Header.Add("test", "hello")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	assert.Nil(t, err)
	assert.True(t, scoped.IsDone())
	assert.Equal(t, 201, res.StatusCode)
	assert.Equal(t, string(body), "hello world")
}

```

## Todo

- [ ] CLI
- [ ] Proxy and Record
- [ ] Configure mocks with JSON/YAML files

## License

This project is [MIT Licensed](LICENSE).

<p align="center"><a href="#mocha-top">back to the top</a></p>
