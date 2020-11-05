package main

import (
	"fmt"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"path/filepath"
)

func t1() {
	ss := snap.New("/root/test/member/snap")
	var snapshot *raftpb.Snapshot
	var err error
	snapshot, err = ss.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println(snapshot.Metadata.Index)
	fmt.Println(snapshot.Metadata.Term)
	fmt.Println(filepath.Join("/root/test/member/snap", fmt.Sprintf("%016x.snap.db", snapshot.Metadata.Index)))
	snapPath, err := ss.DBFilePath(snapshot.Metadata.Index)
	if err != nil {
		panic(err)
	}
	fmt.Println(snapPath)
}

func main() {
	t1()
}