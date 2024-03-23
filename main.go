package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	. "github.com/brianw0924/go_assignment_2/quorum"
)

func main() {

	// How many members
	numMember, err := strconv.Atoi(os.Args[1])
	fmt.Println(numMember)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize members
	MemberList := []Member{}
	for i := 0; i < numMember; i += 1 {
		MemberList = append(MemberList, NewMember(i, numMember))
	}

	// Register each others' communication channel, could be replace to other, like rpc
	for i, m1 := range MemberList {
		for j, m2 := range MemberList {
			if i != j {
				m1.MemberEntries[m2.Id] = NewMemberEntry(m2)
			}
		}
	}

	// Start the members
	for _, m := range MemberList {
		go m.Start()
	}

	// Kill member
	var operation, id string
	for {
		_, err := fmt.Scan(&operation, &id)
		if err != nil {
			fmt.Println(err)
		}
		killId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Invalid command")
		}
		MemberList[killId].Kill()
	}
}
