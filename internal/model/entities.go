package model

import "time"

type Record struct {
	ID, Class, Song, Status, Note string
	Version                       int
	UpdatedAt                     time.Time
}
type User struct {
	ID, Name, Role string
	Active         bool
}
type Event struct {
	ID, RecordID, Kind string
	At                 time.Time
	Payload            string
}
type Audit struct {
	ID, RecordID, Actor, Action, Detail string
	At                                  time.Time
}

func NewRecord(id, class, song string) Record {
	return Record{ID: id, Class: class, Song: song, Status: "draft", Version: 1, UpdatedAt: time.Now().UTC()}
}
func (r Record) IsForClass31() bool { return r.Class == "31" }
func (r Record) IsFinal() bool      { return r.Status == "approved" || r.Status == "archived" }
func (r Record) Clone() Record      { return r }
