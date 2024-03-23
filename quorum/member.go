package quorum

import (
	"fmt"
)

type State int

const (
	FOLLOWER  = iota
	CANDIDATE = iota
	LEADER    = iota
)

type Member struct {
	Id        int
	Term      int
	VoteCount int
	State
	HeartBeatChan   chan Message
	VoteRequestChan chan Message
	VoteChan        chan Message
	ShutDown        chan struct{}
	MemberEntries   map[int]*MemberEntry
	MaxMember       int
}

func NewMember(id, maxMember int) Member {
	return Member{
		Id:              id,
		Term:            0,
		VoteCount:       1,
		State:           FOLLOWER,
		HeartBeatChan:   make(chan Message, 1),
		VoteRequestChan: make(chan Message, 1),
		VoteChan:        make(chan Message, maxMember),
		ShutDown:        make(chan struct{}),
		MemberEntries:   make(map[int]*MemberEntry),
		MaxMember:       maxMember,
	}
}

func (m *Member) Start() {
	fmt.Printf("Member %d: Hi\n", m.Id)
	for {
		switch m.State {
		case FOLLOWER:
			select {

			// Receive heart beat
			case heartBeat := <-m.HeartBeatChan:
				// fmt.Printf("Member %d: Accept heart beat\n", m.Id, m.Term)
				if heartBeat.Term > m.Term {
					m.Term = heartBeat.Term
				}

			// Timeout
			case <-ElectionTimeout():
				m.StartElection()

			// Receive vote request
			case voteRequest := <-m.VoteRequestChan:
				if voteRequest.Term > m.Term {
					m.Term = voteRequest.Term
					m.VoteLeader(voteRequest)
				}

			case <-m.ShutDown:
				return
			}
		case CANDIDATE:
			select {

			// Receive heart beat
			case heartBeat := <-m.HeartBeatChan:
				// fmt.Printf("Member %d: Accept heart beat\n", m.Id, m.Term)
				if heartBeat.Term >= m.Term {
					m.State = FOLLOWER
				}

			// Timeout
			case <-ElectionTimeout():
				// fmt.Printf("Member %d failed to be leader (%d <= %d/2)\n", m.Id, m.Term, m.VoteCount, m.MaxMember)
				m.StartElection()

			// Receive vote
			case vote := <-m.VoteChan:
				// fmt.Printf("Member %d: Got vote\n", m.Id, m.Term)
				if vote.Term == m.Term {
					m.VoteCount += 1
					if m.VoteCount > m.MaxMember/2 {
						fmt.Printf("Member %d voted to be leader (%d > %d/2)\n", m.Id, m.VoteCount, m.MaxMember)
						m.State = LEADER
					}
				}
			case <-m.ShutDown:
				return
			}
		case LEADER:
			select {

			// Send heart beat
			case <-HeartBeatInterval():
				m.SendHeartBeat()

			// Receive heart beat
			case heartBeat := <-m.HeartBeatChan:
				if heartBeat.Term > m.Term {
					m.State = FOLLOWER
				}
			// Receive vote request
			case voteRequest := <-m.VoteRequestChan:
				if voteRequest.Term > m.Term {
					m.VoteLeader(voteRequest)
				}
			case <-m.ShutDown:
				return
			}
		}
	}
}

func (m *Member) StartElection() {
	m.Term += 1
	m.VoteCount = 1
	m.State = CANDIDATE
	fmt.Printf("Member %d: I want to be leader\n", m.Id)
	m.SendVoteRequest()
}

func (m *Member) SendHeartBeat() {
	for _, e := range m.MemberEntries {
		select {
		case e.HeartBeatChan <- NewMessage(m.Id, m.Term):
			// fmt.Printf("Member %d: Send heart beat\n", m.Id, m.Term)
		default:
		}
	}
}

func (m *Member) SendVoteRequest() {
	for _, e := range m.MemberEntries {
		select {
		case e.VoteRequestChan <- NewMessage(m.Id, m.Term):
			// fmt.Printf("Member %d: Send vote request\n", m.Id, m.Term)
		default:
		}
	}
}

func (m *Member) VoteLeader(voteRequest Message) {
	select {
	case m.MemberEntries[voteRequest.Id].VoteChan <- NewMessage(m.Id, voteRequest.Term):
		fmt.Printf("Member %d: Accept member %d to be leader\n", m.Id, voteRequest.Id)
	default:
	}
}

func (m *Member) Kill() {
	close(m.ShutDown)
}
