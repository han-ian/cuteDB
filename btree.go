package cutedb

// btree - Our in memory btree struct
type btree struct {
	root node
}

// node - Interface for node
type node interface {
	insertPair(value *pairs, bt *btree) error
	getValue(key string) (string, error)
	printTree(level int)
}

func (bt *btree) isRootNode(n node) bool {
	return bt.root == n
}

// NewBtree - Create a new btree
func initializeBtree(path ...string) (*btree, error) {
	bs := initBlockService()
	dns := newDiskNodeService(bs)

	root, err := dns.getRootNodeFromDisk()
	if err != nil {
		panic(err)
	}
	return &btree{root: root}, nil
}

// insert - insert element in tree
func (bt *btree) insert(value *pairs) error {
	return bt.root.insertPair(value, bt)
}

func (bt *btree) get(key string) (string, bool, error) {
	value, err := bt.root.getValue(key)
	if err != nil {
		return "", false, err
	}
	if value == "" {
		return "", false, nil
	}
	return value, true, nil
}

func (bt *btree) setRootNode(n node) {
	bt.root = n
}
