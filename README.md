* The implementation is based on Raft version
```
go run main.go <NUM_MEMBERS>
kill <TARGET_MEMBER>
```
* Example output
```
$ go run main.go 10
10
Member 1: Hi
Member 4: Hi
Member 2: Hi
Member 6: Hi
Member 5: Hi
Member 3: Hi
Member 7: Hi
Member 8: Hi
Member 0: Hi
Member 9: Hi
Member 1: I want to be leader
Member 4: Accept member 1 to be leader
Member 8: Accept member 1 to be leader
Member 3: Accept member 1 to be leader
Member 7: Accept member 1 to be leader
Member 9: Accept member 1 to be leader
Member 6: Accept member 1 to be leader
Member 1 voted to be leader (6 > 10/2)
Member 0: Accept member 1 to be leader
Member 2: Accept member 1 to be leader
Member 5: Accept member 1 to be leader
kill 1
Member 7: I want to be leader
Member 4: Accept member 7 to be leader
Member 9: Accept member 7 to be leader
Member 2: Accept member 7 to be leader
Member 3: Accept member 7 to be leader
Member 6: Accept member 7 to be leader
Member 0: Accept member 7 to be leader
Member 5: Accept member 7 to be leader
Member 8: Accept member 7 to be leader
Member 7 voted to be leader (6 > 10/2)
```