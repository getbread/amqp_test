package shared

type Payload struct {
	Foo string
	Bar string
	Baz string
}

func NewPayload() Payload {
	return Payload{"foo", "bar", "baz"}
}
