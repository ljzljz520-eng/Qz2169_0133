package workflow

import (
	"classsong/internal/model"
	"classsong/internal/service"
	"fmt"
)

type Pipeline struct{ Records *service.RecordService }

func NewPipeline(r *service.RecordService) *Pipeline { return &Pipeline{Records: r} }
func (p *Pipeline) Receive(class, song, actor string) (model.Record, error) {
	return p.Records.Register(class, song, actor)
}
func (p *Pipeline) Process(id, actor string) (model.Record, error) {
	r, err := p.Records.Submit(id, actor)
	if err != nil {
		return r, err
	}
	r, err = p.Records.Process(id, actor)
	if err != nil {
		return r, err
	}
	return p.Records.Approve(id, actor)
}
func (p *Pipeline) Archive(id, actor string) (model.Record, error) {
	r, err := p.Records.Archive(id, actor)
	if err != nil {
		return r, fmt.Errorf("archive: %w", err)
	}
	return r, nil
}
func (p *Pipeline) Complete(id, actor string) (model.Record, error) {
	r, err := p.Process(id, actor)
	if err != nil {
		return r, err
	}
	return p.Archive(id, actor)
}
