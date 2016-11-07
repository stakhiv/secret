package main

type Secret struct {
	coder   Coder
	storage Storage
}

func NewSecret(c Coder, s Storage) *Secret {
	return &Secret{
		coder:   c,
		storage: s,
	}
}

func (s *Secret) Store(name string, val []byte) error {
	key, err := s.coder.Key([]byte(name))
	if err != nil {
		return err
	}
	c, err := s.coder.Encode(val)
	if err != nil {
		return err
	}
	err = s.storage.Set(key, c)
	if err != nil {
		return err
	}
	return nil
}

func (s *Secret) Get(name string) ([]byte, error) {
	key, err := s.coder.Key([]byte(name))
	if err != nil {
		return nil, err
	}
	b, err := s.storage.Get(key)
	if err != nil {
		return nil, err
	}

	b, err = s.coder.Decode(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
