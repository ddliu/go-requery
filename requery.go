package requery

func NewDoc() 

type Context struct {

}

type Collection struct {

}

func (this *Context) Find(re string) *Context {

}

func (this *Context) FindAll(re string) *Collection {

}

func (this *Context) Extract() []byte {

}

func (this *Context) ExtractSubmatch() []byte {

}

func (this *Context) Use(index interface{}) *Context {

}

func (this *Context) Call(func())

func (this *Context) ExtractAll()