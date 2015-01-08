# go-requery

Query text with the power of regexp.

## Usage

```go
package main

import (
    "github.com/ddliu/go-requery"
)

func main() {
    doc := requery.NewDoc(`hello _xxxx_ world _xx_`)
    doc.Find(`_x+_`).Extract()
    doc.Find(`_(x+)_).Use(1).Extract()
    doc.FindAll(`_x+_`).Extract()
    doc.FindAll(`_x+_`).FindAll(`_`).ExtractSubmatch()
}
```

## API

Context

```go
Context.Find(`regexp`)
Context.FindAll(`regexp`)
Collection.Find(`regexp`)
Collection.FindAll(`regexp`)
Collection.Extract()
Collection.Extract
Extract()
ExtractSubmatch(3)
```