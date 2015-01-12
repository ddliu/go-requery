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
    doc.Find(`_x+_`).String()
    doc.Find(`_(x+)_).Sub(1).String()
    doc.FindAll(`_x+_`).String()
    doc.FindAll(`_x+_`).FindAll(`_`).ExtractSubmatch()
}
```

## API

### Basic

```go
NewDoc
Find
FindAll
```

### Context

```go
Context.Find
Context.FindAll
Context.String
Context.Bytes
Context.Sub
Context.SubString
Context.SubBytes
```

### Collection

```go
Collection.Find
Collection.FindAll
Collection.Each
```