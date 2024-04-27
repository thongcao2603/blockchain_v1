package core

type Storage interface {
	Put(*Block) error
}

type MemoryStorage struct {
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (*MemoryStorage) Put(*Block) error {
	return nil
}
