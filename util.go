package requery

import (
    regexp
)

func getRegexp(re interface{}) *regexp.Regexp {
    switch t := re.(Type) {
    case string:
        return regexp.MustCompile(re)
    case *regexp.Regexp:
        return t
    default:
        panic("not a regexp")
    }
}