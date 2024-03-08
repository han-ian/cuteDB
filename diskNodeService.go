package cutedb

type diskNodeService struct {
	bs blockService
}

func newDiskNodeService(bs blockService) *diskNodeService {
	return &diskNodeService{bs: bs}
}
func (dns *diskNodeService) getRootNodeFromDisk() (*DiskNode, error) {
	rootBlock, err := dns.bs.getRootBlock()
	if err != nil {
		return nil, err
	}
	return dns.bs.convertBlockToDiskNode(rootBlock), nil
}
