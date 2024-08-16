package vite

import (
	"encoding/json"
	"os"
	"regexp"
)

type ManifestEntry struct {
	File string `json:"file"`
	Src  string `json:"src"`
}

func LoadViteManifest() (map[string]*ManifestEntry, error) {
	entries := make(map[string]*ManifestEntry)

	manifest, err := os.Open("public/manifest.json")

	if err != nil {
		return nil, err
	}

	defer manifest.Close()

	decoder := json.NewDecoder(manifest)
	if err := decoder.Decode(&entries); err != nil {
		return nil, err
	}

	var mapEntries = make(map[string]*ManifestEntry)

	for key, entry := range entries {
		regexp := regexp.MustCompile(`\.[a-zA-Z0-9]+\.`)
		newKey := regexp.ReplaceAllString(key, ".")
		mapEntries[newKey] = entry
	}

	return mapEntries, nil
}
