# go-requery [![Build Status](https://travis-ci.org/ddliu/go-requery.svg)](https://travis-ci.org/ddliu/go-requery) [![GoDoc](https://godoc.org/github.com/ddliu/go-requery?status.svg)](https://godoc.org/github.com/ddliu/go-requery)

Query text with the power of regexp.

## Usage

```go
package main

import (
    "github.com/ddliu/go-requery"
)

func main() {
    // create doc
    doc := requery.NewDoc(`<html>...</html>`)

    // get the title tag
    doc.Find(`<title>.*</title>`).String()

    // get page title
    doc.Find(`<title>(.*)</title>`.Sub(1).String()

    // named match
    doc.Find(`<title>(?P<title>.*)</title>`.Sub("title").String()

    // a shortcut
    doc.Find(`<title>(?P<title>.*)</title>`.SubString("title")

    // collection
    doc.FindAll(`<table>.*?</table>`)[0].String()

    // column of string
    doc.FindAll(`<table>.*?</table>`).FindAll(`<a\s+href="(.*?)">`).SubStringAll(1)
}
```