package cutedb

type blockService interface {
	getLatestBlockID() (int64, error)
	getRootBlock() (*diskBlock, error)
	getBlockFromDiskByBlockNumber(index int64) (*diskBlock, error)
	getBlockFromBuffer(blockBuffer []byte) *diskBlock
	getBufferFromBlock(block *diskBlock) []byte
	newBlock() (*diskBlock, error)
	writeBlockToDisk(block *diskBlock) error
	convertDiskNodeToBlock(node *DiskNode) *diskBlock
	getNodeAtBlockID(blockID uint64) (*DiskNode, error)
	convertBlockToDiskNode(block *diskBlock) *DiskNode
	saveNewNodeToDisk(n *DiskNode) error
	updateNodeToDisk(n *DiskNode) error
	updateRootNode(n *DiskNode) error
	rootBlockExists() bool
	getMaxLeafSize() int
}
