package storage

type mockStorage struct {
	Data map[string] interface{}
	Err error
	Spy bool
}

func NewStorageMock(data map[string]interface{}) *mockStorage {
	return &mockStorage{
		Data: data,
	}
}

func (storage *mockStorage) GetValue(key string) interface{} {
	storage.Spy = true
	if val := storage.Data[key]; val != nil{
		return val
	}
	return nil
}


