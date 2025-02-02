package raft

import (
	"log"
	"math/rand"
	"time"
)

const Debug = true

func (rf *Raft) DPrintf(debug bool, format string, a ...interface{}) (n int, err error) {
	if Debug && debug {
		log.Printf(format, a...)
	}
	return
}

func GetElectionTimeout() int {
	rand.Seed(time.Now().UnixNano())
	return ELECTION_TIMEOUT_BASE + int(rand.Int31n(ELECTION_TIMEOUT_RANGE))
}

func (rf *Raft) GetLogEntry(index int) LogEntry {
	// 要考虑被日志压缩的条目
	if index == 0 {
		return LogEntry{Term: -1, Index: 0}
	} else if index == rf.LastIncludedIndex {
		return LogEntry{Term: rf.LastIncludedTerm, Index: rf.LastIncludedIndex}
	} else {
		return rf.Log[index-rf.LastIncludedIndex-1]
	}
}

func (rf *Raft) GetLastLogEntry() LogEntry {
	// 如果日志为空，返回快照的最后一个日志条目
	if len(rf.Log) == 0 {
		return LogEntry{Term: rf.LastIncludedTerm, Index: rf.LastIncludedIndex}
	}
	return rf.Log[len(rf.Log)-1]
}

func (rf *Raft) GetFirstLogEntry() LogEntry {
	if len(rf.Log) == 0 {
		return LogEntry{Term: rf.LastIncludedTerm, Index: rf.LastIncludedIndex}
	}
	return rf.Log[0]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
