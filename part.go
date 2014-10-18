package hydrator

type Part interface {
	Render() []byte
}

type LiteralPart []byte

func (p LiteralPart) Render() []byte {
	return p
}

type ReferencePart struct {
	id string
	t  string
}

func (p *ReferencePart) Render() []byte {
	return Get(p.id, p.t)
}
