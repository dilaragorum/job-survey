package questionary

type Select struct {
	Index int
	Value interface{}
}

type MessageForClient struct {
	Text   string
	Status bool
}
