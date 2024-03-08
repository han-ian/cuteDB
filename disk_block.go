package cutedb

import (
	"encoding/binary"
)

const blockSize = 4096

// Based on the below calc
const maxLeafSize = 30

// diskBlock -- Make sure that it is accomodated in blockSize = 4096
type diskBlock struct {
	id                  uint64   // 4096 - 8 = 4088
	currentLeafSize     uint64   // 4088 - 8 = 4080
	currentChildrenSize uint64   // 4080 - 8 = 4072
	childrenBlockIds    []uint64 // 262 - (8 * 30) =  22
	dataSet             []*pairs // 4072 - (127 * 30) = 262
}

// 22 bytes are still wasted

func (b *diskBlock) setData(data []*pairs) {
	b.dataSet = data
	b.currentLeafSize = uint64(len(data))
}

func (b *diskBlock) setChildren(childrenBlockIds []uint64) {
	b.childrenBlockIds = childrenBlockIds
	b.currentChildrenSize = uint64(len(childrenBlockIds))
}

func (block *diskBlock) convertBlockToDiskNode(bs blockService) *DiskNode {
	node := &DiskNode{
		blockID:      block.id,
		blockService: bs,
		keys:         make([]*pairs, block.currentLeafSize),
	}
	for index := range node.keys {
		node.keys[index] = block.dataSet[index]
	}
	node.childrenBlockIDs = make([]uint64, block.currentChildrenSize)
	for index := range node.childrenBlockIDs {
		node.childrenBlockIDs[index] = block.childrenBlockIds[index]
	}
	return node
}

func (block *diskBlock) getBufferFromBlock() []byte {
	blockBuffer := make([]byte, blockSize)
	blockOffset := 0

	//Write Block index
	copy(blockBuffer[blockOffset:], uint64ToBytes(block.id))
	blockOffset += 8
	copy(blockBuffer[blockOffset:], uint64ToBytes(block.currentLeafSize))
	blockOffset += 8
	copy(blockBuffer[blockOffset:], uint64ToBytes(block.currentChildrenSize))
	blockOffset += 8

	//Write actual pairs now
	for i := 0; i < int(block.currentLeafSize); i++ {
		copy(blockBuffer[blockOffset:], convertPairsToBytes(block.dataSet[i]))
		blockOffset += pairSize
	}
	// Read children block indexes
	for i := 0; i < int(block.currentChildrenSize); i++ {
		copy(blockBuffer[blockOffset:], uint64ToBytes(block.childrenBlockIds[i]))
		blockOffset += 8
	}
	return blockBuffer
}

func getBlockFromBuffer(blockBuffer []byte) *diskBlock {
	blockOffset := 0
	block := &diskBlock{}

	//Read Block index
	block.id = uint64FromBytes(blockBuffer[blockOffset:])
	blockOffset += 8
	block.currentLeafSize = uint64FromBytes(blockBuffer[blockOffset:])
	blockOffset += 8
	block.currentChildrenSize = uint64FromBytes(blockBuffer[blockOffset:])
	blockOffset += 8
	//Read actual pairs now
	block.dataSet = make([]*pairs, block.currentLeafSize)
	for i := 0; i < int(block.currentLeafSize); i++ {
		block.dataSet[i] = convertBytesToPair(blockBuffer[blockOffset:])
		blockOffset += pairSize
	}
	// Read children block indexes
	block.childrenBlockIds = make([]uint64, block.currentChildrenSize)
	for i := 0; i < int(block.currentChildrenSize); i++ {
		block.childrenBlockIds[i] = uint64FromBytes(blockBuffer[blockOffset:])
		blockOffset += 8
	}
	return block
}

func uint64ToBytes(index uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(index))
	return b
}

func uint64FromBytes(b []byte) uint64 {
	return uint64(binary.LittleEndian.Uint64(b))
}
