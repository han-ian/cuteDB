package cutedb

type blockService interface {
	getLatestBlockID() (int64, error)
	getRootBlock() (*diskBlock, error)
	getBlockFromDiskByBlockNumber(index int64) (*diskBlock, error)
	getBlockFromBuffer(blockBuffer []byte) *diskBlock
	newBlock() (*diskBlock, error)
	writeBlockToDisk(block *diskBlock) error
	getNodeAtBlockID(blockID uint64) (*DiskNode, error)
	convertBlockToDiskNode(block *diskBlock) *DiskNode
	saveNewNodeToDisk(n *DiskNode) error
	updateNodeToDisk(n *DiskNode) error
	updateRootNode(n *DiskNode) error
	rootBlockExists() bool
	getMaxLeafSize() int
}

func initBlockService() blockService {
	return initblockServiceFileStore()
	// return initblockServicKvStore()
}
