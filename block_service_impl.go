package cutedb

import (
	"os"
)

func initblockServiceFileStore() blockService {
	path := "./db/test.db"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir("./db", os.ModePerm)
	}
	if _, err := os.Stat(path); err == nil {
		// path/to/whatever exists
		err := os.Remove(path)
		if err != nil {
			panic(err)
		}
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	return &blockServiceImpl{blockStore: newFileStore(file)}
}

func initblockServicKvStore() blockService {
	return &blockServiceImpl{blockStore: newKvStore()}
}

type blockServiceImpl struct {
	blockStore blockStore
}

func (bs *blockServiceImpl) getLatestBlockID() (int64, error) {
	return bs.blockStore.getLatestBlockID()
}

//@Todo:Store current root block data somewhere else
func (bs *blockServiceImpl) getRootBlock() (*diskBlock, error) {

	/*
		1. Check if root block exists
		2. If exisits, fetch it, else initialize a new block
	*/
	if !bs.rootBlockExists() {
		// Need to write a new block
		return bs.newBlock()

	}
	return bs.getBlockFromDiskByBlockNumber(0)

}

func (bs *blockServiceImpl) getBlockFromDiskByBlockNumber(index int64) (*diskBlock, error) {
	if index < 0 {
		panic("Index less than 0 asked")
	}

	blockBuffer, err := bs.blockStore.getBlockFromDisk(index)
	if err != nil {
		return nil, err
	}
	block := bs.getBlockFromBuffer(blockBuffer)
	return block, nil
}

func (bs *blockServiceImpl) getBlockFromBuffer(blockBuffer []byte) *diskBlock {
	return getBlockFromBuffer(blockBuffer)
}

func (bs *blockServiceImpl) newBlock() (*diskBlock, error) {

	latestBlockID, err := bs.getLatestBlockID()
	block := &diskBlock{}
	if err != nil {
		// This means that no file exists
		block.id = 0
	} else {
		block.id = uint64(latestBlockID) + 1
	}
	block.currentLeafSize = 0
	err = bs.writeBlockToDisk(block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (bs *blockServiceImpl) writeBlockToDisk(block *diskBlock) error {
	err := bs.blockStore.writeBlockToDisk(block)
	if err != nil {
		return err
	}
	return nil
}

func (bs *blockServiceImpl) getNodeAtBlockID(blockID uint64) (*DiskNode, error) {
	block, err := bs.getBlockFromDiskByBlockNumber(int64(blockID))
	if err != nil {
		return nil, err
	}
	return bs.convertBlockToDiskNode(block), nil
}

func (bs *blockServiceImpl) convertBlockToDiskNode(block *diskBlock) *DiskNode {
	return block.convertBlockToDiskNode(bs)
}

// NewBlockFromNode - Save a new node to disk block
func (bs *blockServiceImpl) saveNewNodeToDisk(n *DiskNode) error {
	// Get block id to be assigned to this block
	latestBlockID, err := bs.getLatestBlockID()
	if err != nil {
		return err
	}
	n.blockID = uint64(latestBlockID) + 1
	block := n.convertDiskNodeToBlock()
	return bs.writeBlockToDisk(block)
}

func (bs *blockServiceImpl) updateNodeToDisk(n *DiskNode) error {
	block := n.convertDiskNodeToBlock()
	return bs.writeBlockToDisk(block)
}

func (bs *blockServiceImpl) updateRootNode(n *DiskNode) error {
	n.blockID = 0
	return bs.updateNodeToDisk(n)
}

func (bs *blockServiceImpl) rootBlockExists() bool {
	latestBlockID, err := bs.getLatestBlockID()
	// fmt.Println(latestBlockID)
	//@Todo:Validate the type of error here
	if err != nil {
		// Need to write a new block
		return false
	} else if latestBlockID == -1 {
		return false
	} else {
		return true
	}
}

/**
@Todo: Implement a function to :
1. Dynamicaly calculate blockSize
2. Then based on the blocksize, calculate the maxLeafSize
*/
func (bs *blockServiceImpl) getMaxLeafSize() int {
	return maxLeafSize
}
