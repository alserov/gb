package main

import (
	"fmt"
	"time"
)

func main() {

}

type List interface {
	Push(link string, opts ...PushOption)
	Delete(link string, opts ...DeleteOption) (int, error)
	Get(link string, opts ...GetOption) ([]Item, error)

	Length() int
	All() []Item
	Clear()
}

type (
	PushOption   func(item *Item)
	GetOption    func(p *Params)
	DeleteOption func(p *Params)

	ListType int
)

const (
	LINKED_LIST ListType = iota
	MAP
)

func NewCollection(t ListType) List {
	switch t {
	case LINKED_LIST:
		return &withList{}
	//case MAP:
	//	return &withMap{}
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

	currentLimit  int
	currentOffset int
}

func WithName(name string) func(item *Item) {
	return func(item *Item) {
		item.Name = name
	}
}

func WithTags(t []string) func(item *Item) {
	return func(item *Item) {
		item.Tags = t
	}
}

func WithLimit(l int) func(p *Params) {
	n := l
	return func(p *Params) {
		p.Limit = &n
	}
}

func WithOffset(o int) func(p *Params) {
	n := o
	return func(p *Params) {
		p.Offset = &n
	}
}
