package main

import (
	"time"
)

type withList struct {
	head *Item
	tail *Item
	len  int
}

func (w *withList) Clear() {
	w.head, w.tail, w.len = nil, nil, 0
}

func (w *withList) All() []Item {
	all := make([]Item, 0, w.len)
	c := w.tail
	for c != nil {
		all = append(all, *c)
		c = c.next
	}
	return all
}

func (w *withList) Length() int {
	return w.len
}

func (w *withList) Delete(link string, p ...DeleteOption) (int, error) {
	prms := Params{}

	for _, prm := range p {
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

			if link == currItem.Link {
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

func (w *withList) Get(link string, p ...GetOption) ([]Item, error) {
	prms := Params{}

	for _, prm := range p {
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

func (w *withList) Push(val string, opts ...PushOption) {
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
		return
	}

	w.head.next = &i
	w.head = w.head.next
	w.len++
}
