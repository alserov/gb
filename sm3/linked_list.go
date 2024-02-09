package main

import (
	"slices"
	"time"
)

type withList struct {
	head *Item
	tail *Item
	len  int
}

func (w *withList) Clear() error {
	w.head, w.tail, w.len = nil, nil, 0
	return nil
}

func (w *withList) All() ([]Item, error) {
	all := make([]Item, 0, w.len)
	c := w.tail
	for c != nil {
		all = append(all, *c)
		c = c.next
	}
	return all, nil
}

func (w *withList) Length() int {
	return w.len
}

func (w *withList) Delete(link string, opts ...Option) (int, error) {
	prms := Params{}

	for _, prm := range opts {
		prm(&prms)
	}

	var (
		currItem       = w.tail
		prev           *Item
		shouldContinue bool
	)

	for {
		if !shouldContinue && prev != nil || currItem == nil {
			return prms.currentLimit, nil
		}
		shouldContinue = func() bool {
			defer func() {
				currItem = currItem.next
			}()

			if link == currItem.Link || prms.Tag != nil && slices.Contains(currItem.Tags, *prms.Tag) {
				if prms.Limit != nil {
					if prms.currentLimit >= *prms.Limit {
						return false
					}
				}

				if prms.Offset != nil {
					if prms.currentOffset < *prms.Offset {
						prev = currItem
						prms.currentOffset++
						return true
					}
				}

				if prev == nil {
					w.tail = w.tail.next
				}

				w.len--
				prms.currentLimit++
			}

			if prev == nil {
				prev = currItem
			} else {
				prev.next = currItem.next
			}
			return true
		}()
	}
}

func (w *withList) Get(link string, opts ...Option) ([]Item, error) {
	prms := Params{}

	for _, prm := range opts {
		prm(&prms)
	}

	var (
		v        []Item
		currItem = w.tail
	)
	for {
		if currItem == nil {
			return v, nil
		}
		func() {
			defer func() {
				currItem = currItem.next
			}()
			if prms.Limit != nil {
				if prms.currentLimit >= *prms.Limit {
					return
				}
			}
			if link == currItem.Link {
				if prms.Offset != nil {
					if prms.currentOffset >= *prms.Offset {
						v = append(v, *currItem)
						prms.currentLimit++
						return
					}
				}
				v = append(v, *currItem)
			}
		}()
	}
}

func (w *withList) Push(val string, opts ...PushOption) error {
	i := Item{
		Link: val,
		Date: time.Now(),
	}

	for _, opt := range opts {
		opt(&i)
	}

	if w.tail == nil {
		w.head = &i
		w.tail = &i
		w.len++
		return nil
	}

	w.head.next = &i
	w.head = w.head.next
	w.len++

	return nil
}
