package server

import (
	"fmt"
	"strings"
)

func (s *Server) LoadAsset(name string) (string, error) {
	name = strings.TrimSpace(name)
	entry, ok := s.manifest[name]

	if !ok {
		return "", fmt.Errorf("asset %s not found", name)
	}

	return entry.Src, nil
}
