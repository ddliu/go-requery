// Copyright (c) 2015 Liu Dong <ddliuhb@gmail.com>
// Licensed under the MIT license

package requery

import (
    "regexp"
)

func getRegexp(re interface{}) *regexp.Regexp {
    switch t := re.(type) {
    case string:
        return regexp.MustCompile(t)
    case *regexp.Regexp:
        return t
    default:
        panic("not a regexp")
    }
}

func getRegexpString(re interface{}) string {
    switch t := re.(type) {
    case string:
        return t
    case *regexp.Regexp:
        return t.String()
    default:
        return ""
    }
}