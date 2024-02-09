package main

import (
	"fmt"
	"time"
)

type List interface {
	Push(link string, opts ...PushOption) error
	Delete(link string, opts ...Option) (int, error)
	Get(link string, opts ...Option) ([]Item, error)

	Length() int
	All() ([]Item, error)
	Clear() error
}

type (
	PushOption func(item *Item)
	Option     func(p *Params)

	ListType int
)

const (
	LINKED_LIST ListType = iota
	FILE
	// also can be used a map
)

func NewCollection(t ListType) List {
	switch t {
	case LINKED_LIST:
		return &withList{}
	case FILE:
		return &withFile{}
	default:
		panic(fmt.Sprintf("invalid list type: %d", t))
	}
}

type Item struct {
	Name string
	Date time.Time
	Tags []string
	Link string

	next *Item
}

type Params struct {
	Limit  *int
	Offset *int
	Tag    *string

	currentLimit  int
	currentOffset int
}

func WithName(name string) PushOption {
	return func(item *Item) {
		item.Name = name
	}
}

func WithTags(t []string) PushOption {
	return func(item *Item) {
		item.Tags = t
	}
}

func WithLimit(l int) Option {
	n := l
	return func(p *Params) {
		p.Limit = &n
	}
}

func WithOffset(o int) Option {
	n := o
	return func(p *Params) {
		p.Offset = &n
	}
}

func WithTag(t string) Option {
	s := t
	return func(p *Params) {
		p.Tag = &s
	}
}
