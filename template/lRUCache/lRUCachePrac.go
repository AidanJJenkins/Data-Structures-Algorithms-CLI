package lrucache

type Node struct {
	value any
	next  *Node
	prev  *Node
}

func newNode(value any) *Node {
	return &Node{value, nil, nil}
}

type Key interface{}

type Cache struct {
	length        int
	capacity      int
	head          *Node
	tail          *Node
	lookup        map[Key]*Node
	reverseLookup map[*Node]Key
}

func NewLRU(capacity int) *Cache {
	return &Cache{}
}

func (c *Cache) Update(key string, value any) {
}

func (c *Cache) Get(key string) any {

}

func (c *Cache) prepend(node *Node) {
}

func (c *Cache) detach(node *Node) {
}

func (c *Cache) trimCache() {
}
