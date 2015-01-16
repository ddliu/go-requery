// Copyright (c) 2015 Liu Dong <ddliuhb@gmail.com>
// Licensed under the MIT license

package requery

type Collection []*Context

// Find first match in the collection
func (this Collection) Find(s interface{}) *Context {
    re := getRegexp(s)

    for _, c := range this {
        cc := c.Find(re)
        if !cc.Empty() {
            return cc
        }
    }

    return NewDoc(nil)
}

// Find with assertion
func (this Collection) MustFind(s interface{}) *Context {
    c := this.Find(s)
    if c.Empty() {
        panic("Not found")
    }

    return c
}

// Find all matches in the collection
func (this Collection) FindAll(s interface{}) Collection {
    re := getRegexp(s)
    var result Collection
    for _, c := range this {
        result = append(result, c.FindAll(re)...)
    }

    return result
}

// Find all with assertion
func (this Collection) MustFindAll(s interface{}) Collection {
    result := this.FindAll(s)
    if len(result) == 0 {
        panic("Not found")
    }

    return result
}

// Get a column of string
func (this Collection) SubStringAll(s interface{}) []string {
    result := make([]string, len(this))

    for k, c := range this {
        result[k] = c.SubString(s)
    }

    return result
}

// Get a column of bytes
func (this Collection) SubBytesAll(s interface{}) [][]byte {
    result := make([][]byte, len(this))

    for k, c := range this {
        result[k] = c.SubBytes(s)
    }

    return result
}

// Loop the collection
func (this Collection) Each(f func(*Context)) Collection {
    for _, c := range this {
        f(c)
    }

    return this
}