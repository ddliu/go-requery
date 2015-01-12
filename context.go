// Copyright (c) 2015 Liu Dong <ddliuhb@gmail.com>
// Licensed under the MIT license

package requery

import (
    "fmt"
)

func NewDoc(content interface{}) *Context {
    c := &Context {
        Names: make(map[string]int),
    }

    switch t := content.(type) {
    case string:
        if t == "" {
            c.List = nil
        } else {
            c.List = [][]byte{[]byte(t)}
        }
    case []byte:
        if t == nil {
            c.List = nil
        } else {
            c.List = [][]byte{t}
        }
    default: 
        // panic?
        // panic("invalid doc")
    }

    return c
}

type Context struct {
    Parent *Context
    List [][]byte
    Names map[string]int
    // NamedMatches map[string][]byte
}

// Get string content of the context
func (this *Context) String() string {
    return string(this.Bytes())
}

// Get bytes content of the context
func (this *Context) Bytes() []byte {
    if len(this.List) != 0 {
        return this.List[0]
    }

    return nil
}

// Find the first match
func (this *Context) Find(re interface{}) *Context {
    r := getRegexp(re)
    matches := r.FindSubmatch(this.Bytes())
    names := make(map[string]int)

    for k, v := range r.SubexpNames() {
        if v != "" {
            names[v] = k
        }
    }

    result := &Context {
        Parent: this,
        List: matches,
        Names: names,
    }

    return result
}

// Find all matches
func (this *Context) FindAll(re interface{}) Collection {
    r := getRegexp(re)

    matches := r.FindAllSubmatch(this.Bytes(), -1)

    names := make(map[string]int)

    for k, v := range r.SubexpNames() {
        if v != "" {
            names[v] = k
        }
    }

    result := make(Collection, len(matches))
    for i, match := range matches {
        result[i] = &Context {
            Parent: this,
            List: match,
            Names: names,
        }
    }

    return result
}

func (this *Context) Empty() bool {
    return this.List == nil
}

func (this *Context) Sub(s interface{}) *Context {
    isEmpty := this.Empty()

    switch t := s.(type) {
    case string:
        i, ok := this.Names[t]
        if !ok {
            panic("submatch not exist: " + t)
        }

        var result *Context
        if isEmpty {
            result = NewDoc(nil)
        } else {
            result = NewDoc(this.List[i])
        }
        result.Parent = this

        return result
    case int:
        if len(this.List) < t + 1 && !isEmpty {
            panic(fmt.Sprintf("submatch index not exist: %d", t))
        }

        var result *Context
        if isEmpty {
            result = NewDoc(nil)
        } else {
            result = NewDoc(this.List[t])
        }

        result.Parent = this

        return result
    default:
        panic(fmt.Sprintf("invalid sub index: %v", s))
    }
}