package quorum

type Message struct {
	Id   int
	Term int
}

func NewMessage(id, term int) Message {
	return Message{
		Id:   id,
		Term: term,
	}
}

func (m *Message) IsValid(memberTerm int) bool {
	return m.Term >= memberTerm
}
