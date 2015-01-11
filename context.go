package requery

func NewContext() *Context {
    return &Context {}
}

func NewDoc(content []byte) *Context {
    c := &Context {
        List: [][]byte {content},
        Names: make(map[string]int),
    }

    return c
}

type Context struct {
    Parent *Context
    List [][]byte
    Names: map[string]int
    // NamedMatches map[string][]byte
}

// Get string content of the context
func (this *Context) String() string {
    return string(this.Bytes())
}

// Get bytes content of the context
func (this *Context) Bytes() []byte {
    if len(this.List) {
        return this.List[0]
    }

    return make([]byte)
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
func (this *Context) FindAll(re interface{}) *Collection {
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
        Collection[i] = &Context {
            Parent: this,
            List: match,
            Names: names,
        }
    }

    return result
}

func (this *Context) Sub(s interface{}) *Context {
    switch t := s.(Type) {
    case string:
        i, ok := this.Names[s]
        if !ok {
            panic("submatch not exist: " + s)
        }

    }
}