// Copyright (c) 2015 Liu Dong <ddliuhb@gmail.com>
// Licensed under the MIT license

package requery

func Find(content, re interface{}) *Context {
    return NewDoc(content).Find(re)
}

func FindAll(content, re interface{}) Collection {
    return NewDoc(content).FindAll(re)
}