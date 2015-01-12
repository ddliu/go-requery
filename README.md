# go-requery
[![Build Status](https://travis-ci.org/ddliu/go-requery.svg)](https://travis-ci.org/ddliu/go-requery)
[![GoDoc](https://godoc.org/github.com/ddliu/go-requery?status.svg)](https://godoc.org/github.com/ddliu/go-requery)

Query text with the power of regexp.

## Usage

```go
package main

import (
    "github.com/ddliu/go-requery"
)

func main() {
    doc := requery.NewDoc(`hello _xxxx_ world _xx_`)
    doc.Find(`_x+_`).String()
    doc.Find(`_(x+)_).Sub(1).String()
    doc.FindAll(`_x+_`).String()
    doc.FindAll(`_x+_`).FindAll(`_`)
}
```