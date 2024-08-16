package server

import "fmt"

func (s *Server) LoadAsset(name string) (string, error) {
	entry, ok := s.manifest[name]

	if !ok {
		return "", fmt.Errorf("asset %s not found", name)
	}

	return entry.Src, nil
}
