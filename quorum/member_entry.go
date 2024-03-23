package quorum

type MemberEntry struct {
	HeartBeatChan   chan Message
	VoteRequestChan chan Message
	VoteChan        chan Message
}

func NewMemberEntry(member Member) *MemberEntry {
	return &MemberEntry{
		HeartBeatChan:   member.HeartBeatChan,
		VoteRequestChan: member.VoteRequestChan,
		VoteChan:        member.VoteChan,
	}
}
