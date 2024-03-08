package cutedb

import "fmt"

type kvStore struct {
	blockMap map[int64][]byte
	maxId    int64
}

func newKvStore() blockStore {
	return &kvStore{blockMap: make(map[int64][]byte, 0)}
}

func (s *kvStore) getBlockFromDisk(index int64) ([]byte, error) {
	if val, ok := s.blockMap[index]; ok {
		return val, nil
	}
	return nil, nil
}

func (s *kvStore) writeBlockToDisk(block *diskBlock) error {
	fmt.Printf("writeBlockToDisk: %d\n", block.id)
	if block.id > uint64(s.maxId) {
		s.maxId = int64(block.id)
	}
	blockBuffer := block.getBufferFromBlock()
	s.blockMap[int64(block.id)] = blockBuffer
	return nil
}

func (s *kvStore) getLatestBlockID() (int64, error) {
	if len(s.blockMap) == 0 {
		return -1, nil
	}
	return s.maxId, nil
}
