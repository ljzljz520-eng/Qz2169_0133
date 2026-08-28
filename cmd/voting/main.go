package main

import (
	"classsong/internal/query"
	"classsong/internal/service"
	"classsong/internal/store"
	"classsong/internal/workflow"
	"fmt"
	"os"
)

func main() {
	path := "classsong.db"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	s, err := store.Open(path)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	p := workflow.NewPipeline(service.NewRecordService(s))
	r, err := p.Receive("31", "班级歌会", "cli")
	if err != nil {
		panic(err)
	}
	_, _ = p.Process(r.ID, "cli")
	q := query.New(s)
	latest, _ := q.Latest("31")
	fmt.Printf("%s %s\n", latest.Song, latest.Status)
}
