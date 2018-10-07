package test

import (
	"math/rand"
	"strconv"
)

func (s *suite) NewPort() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	port := rand.Intn(99999)
	if port < 200 {
		port = port + 200
	}

	p := strconv.FormatInt(int64(port), 10)
	if _, ok := s.ports[p]; ok {
		return s.NewPort()
	}
	s.ports[p] = struct{}{}
	return p
}
