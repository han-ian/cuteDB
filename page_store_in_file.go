package cutedb

import (
	"fmt"
	"os"
)

type blockFileStore struct {
	file *os.File
}

func newFileStore(file *os.File) blockStore {
	return &blockFileStore{file: file}
}

func (s *blockFileStore) getBlockFromDisk(index int64) ([]byte, error) {
	offset := index * blockSize
	_, err := s.file.Seek(offset, 0)
	if err != nil {
		return nil, err
	}

	blockBuffer := make([]byte, blockSize)
	_, err = s.file.Read(blockBuffer)
	if err != nil {
		return nil, err
	}
	return blockBuffer, nil
}

func (s *blockFileStore) writeBlockToDisk(block *diskBlock) error {
	fmt.Printf("writeBlockToDisk: %d\n", block.id)
	seekOffset := blockSize * block.id
	blockBuffer := block.getBufferFromBlock()
	_, err := s.file.Seek(int64(seekOffset), 0)
	if err != nil {
		return err
	}
	_, err = s.file.Write(blockBuffer)
	if err != nil {
		return err
	}
	return nil
}

func (s *blockFileStore) getLatestBlockID() (int64, error) {
	fi, err := s.file.Stat()
	if err != nil {
		return -1, err
	}

	length := fi.Size()
	if length == 0 {
		return -1, nil
	}
	// Calculate page number required to be fetched from disk
	return (int64(fi.Size()) / int64(blockSize)) - 1, nil
}
