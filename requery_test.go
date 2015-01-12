// Copyright (c) 2015 Liu Dong <ddliuhb@gmail.com>
// Licensed under the MIT license

package requery

import (
    "testing"
    "io/ioutil"
)

func getDoc() *Context {
    content, _ := ioutil.ReadFile("test/test.html")
    return NewDoc(content)
}

func TestCommon(t *testing.T) {
    doc := getDoc()

    // find title
    if title := doc.Find("<title>(.+)</title>").Sub(1).String(); title != "Hacker News" {
        t.Error("title: " + title)
    }

    // find list number
    if list := doc.FindAll(`<td align="right" valign="top" class="title">\d+\.</td>`); len(list) != 30 {
        t.Error()
    }
}