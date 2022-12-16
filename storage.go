package storage

// Queue FIFO
type Queue[V any] struct {
	queue []*V
}

func (q *Queue[V]) Push(v V) {
	q.queue = append(q.queue, &v)
}

func (q *Queue[V]) Pop() V {
	l := q.Len()
	if l == 0 {
		return *new(V)
	}
	v := *q.queue[0]
	q.queue[0] = nil
	q.queue = q.queue[1:]
	return v
}

func (q Queue[V]) Len() int {
	return len(q.queue)
}

// Stack FILO
type Stack[V any] struct {
	stack []*V
}

func (s *Stack[V]) Push(v V) {
	s.stack = append(s.stack, &v)
}

func (s *Stack[V]) Pop() V {
	l := s.Len()
	if l == 0 {
		return *new(V)
	}
	v := *s.stack[l-1]
	s.stack[l-1] = nil
	s.stack = s.stack[:l-1]
	return v
}

func (s Stack[V]) Len() int {
	return len(s.stack)
}

// Dictionary
type dict[K comparable, V any] struct {
	m map[K]V
}

func NewDict[K comparable, V any]() *dict[K, V] {
	return &dict[K, V]{m: make(map[K]V)}
}

func (d *dict[K, V]) Add(key K, value V) {

	d.m[key] = value
}

func (d *dict[K, V]) Delete(key K) {

	delete(d.m, key)
}

// Get returns the value of the given key
func (d dict[K, V]) Get(key K) V {
	return d.m[key]
}

// Values returns all values in the container
func (d dict[K, V]) Values() []V {
	var values []V
	for k := range d.m {
		values = append(values, d.m[k])
	}
	return values
}

// Keys returns all keys in the container
func (d dict[K, V]) Keys() []K {
	var keys []K
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

func (d dict[K, V]) Len() int {
	return len(d.m)
}

// Trie
type Trie[T any] interface {
	Find(key []byte) (bool, T)
	Insert(key []byte, value T) bool
	Len() int
	Remove(key []byte) bool
}

type trie[T any] struct {
	klength int
	length  int
	root    *node[T]
}

type node[T any] struct {
	value *T
	child map[byte]*node[T]
}

func newNode[T any]() *node[T] {
	return &node[T]{child: make(map[byte]*node[T])}
}

func NewTrie[T any](klen int) *trie[T] {
	return &trie[T]{klength: klen, root: newNode[T]()}
}

func (t *trie[T]) Insert(key []byte, value T) bool {
	if len(key) != t.klength {
		return false
	}
	node := t.root
	for _, b := range key {
		if _, ok := node.child[b]; !ok {
			node.child[b] = newNode[T]()
		}
		node = node.child[b]
	}
	node.value = &value
	t.length++
	return true
}

func (t *trie[T]) Find(key []byte) (bool, T) {
	if len(key) != t.klength {
		return false, *new(T)
	}
	node := t.root
	for _, b := range key {
		if _, ok := node.child[b]; !ok {
			return false, *new(T)
		}
		node = node.child[b]
	}
	return true, *node.value
}

func (t *trie[T]) Remove(key []byte) bool {
	if len(key) != t.klength {
		return false
	}
	node := t.root
	branch, k := node, key[0]
	for _, b := range key {
		if _, ok := node.child[b]; !ok {
			return false
		}
		if len(node.child) > 1 {
			branch, k = node, b
		}
		node = node.child[b]
	}
	delete(branch.child, k)
	return true
}

func (t *trie[T]) Len() int {
	return t.length
}
