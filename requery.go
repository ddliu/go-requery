// Copyright (c) 2015 Liu Dong <ddliuhb@gmail.com>
// Licensed under the MIT license

package requery

func Find(content, re interface{}) *Context {
    return NewDoc(content).Find(re)
}

func MustFind(content, re interface{}) *Context {
    return NewDoc(content).MustFind(re)
}

func FindAll(content, re interface{}) Collection {
    return NewDoc(content).FindAll(re)
}

func MustFindAll(content, re interface{}) Collection {
    return NewDoc(content).MustFindAll(re)
}