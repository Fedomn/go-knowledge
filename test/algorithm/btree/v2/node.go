package v2

type node struct {
	isLeaf   bool
	degree   int     // minimum degree (defines the range for number of keys)
	keys     []int   // inserted keys
	count    int     // current number of keys
	children []*node // child pointers
}
