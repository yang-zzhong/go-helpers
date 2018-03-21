package helpers

type eachCall func(key string, value interface{}) bool

type P struct {
	p map[string]interface{}
}

func NewP() *P {
	return &P{make(map[string]interface{})}
}

func (p *P) Set(key string, value interface{}) {
	p.p[key] = value
}

func (p *P) Get(key string) interface{} {
	return p.p[key]
}

func (p *P) Keys() []string {
	return Keys(p.p)
}

func (p *P) Del(key string) {
	if p.Exists(key) {
		delete(p.p, key)
	}
}

func (p *P) Exists(key string) bool {
	_, ok := p.p[key]
	return ok
}

func (p *P) Each(call eachCall) bool {
	for key, value := range p.p {
		if call(key, value) {
			continue
		}
		return false
	}
	return true
}
