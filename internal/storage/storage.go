package storage

type Storage struct {
	Key       string
	ListSlang []string
}

func (s *Storage) setKeyVal(key string, vals []string) {
	s.Key = key
	s.ListSlang = vals
}

func (s *Storage) getKeyVal() (string, []string) {
	return s.Key, s.ListSlang
}
