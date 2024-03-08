package cutedb

type blockStore interface {
	getBlockFromDisk(index int64) ([]byte, error)
	writeBlockToDisk(block *diskBlock) error
	getLatestBlockID() (int64, error)
}
