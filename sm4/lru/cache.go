package lru

type Cache interface {
	Get(k string) ([]byte, bool)
	Set(k string, v []byte)
}

func NewLRU(lim int) Cache {
	return &lru{
		limit: lim,
	}
}

type lru struct {
	len  int
	tail *node
	head *node

	limit int
}

func (l *lru) Get(k string) ([]byte, bool) {
	c := l.tail

	for c != nil {
		if c.key == k {
			return c.value, true
		}
		c = c.next
	}

	return nil, false
}

func (l *lru) Set(k string, v []byte) {
	l.len++

	if l.len == l.limit {
		l.tail = l.tail.next
		l.len--
	}

	if l.head == nil {
		l.head = &node{
			key:   k,
			value: v,
			prev:  l.head,
		}
		l.tail = &node{
			key:   k,
			value: v,
			next:  l.head,
		}
	} else {
		l.head.next = &node{
			key:   k,
			value: v,
			prev:  l.head,
		}
		l.head = l.head.next
	}
}

type node struct {
	key   string
	value []byte
	prev  *node
	next  *node
}
