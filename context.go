// Copyright (c) 2015 Liu Dong <ddliuhb@gmail.com>
// Licensed under the MIT license

package requery

import (
    "fmt"
)

// Create a doc to query.
// 
// Type of content can be string or []byte
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

func (this *Context) createEmptySub() *Context {
    return &Context {
        Parent: this,
    }
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

func (this *Context) MustFind(re interface{}) *Context {
    result := this.Find(re)
    if result.Empty() {
        panic("Not found")
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

func (this *Context) MustFindAll(re interface{}) Collection {
    result := this.FindAll(re)
    if len(result) == 0 {
        panic("Not found")
    }

    return result
}

// Is it an empty context
func (this *Context) Empty() bool {
    return this.List == nil
}

// Get the sub match by index.
// 
// If index is string, the coresponding named match is returned.
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
            result = this.createEmptySub()
        } else {
            result = NewDoc(this.List[i])
            result.Parent = this
        }

        return result
    case int:
        if len(this.List) < t + 1 && !isEmpty {
            panic(fmt.Sprintf("submatch index not exist: %d", t))
        }

        var result *Context
        if isEmpty {
            result = this.createEmptySub()
        } else {
            result = NewDoc(this.List[t])
            result.Parent = this
        }

        return result
    default:
        panic(fmt.Sprintf("invalid sub index: %v", s))
    }
}


// Return sub match string by index.
func (this *Context) SubString(s interface{}) string {
    return this.Sub(s).String()
}

// Return sub match bytes by index.
func (this *Context) SubBytes(s interface{}) []byte {
    return this.Sub(s).Bytes()
}