// Package config reads and writes the CLI's profiles, in
// $XDG_CONFIG_HOME/yesly/config.json (by default ~/.config/yesly/config.json),
// readable only by the user.
package config

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

// Profile is one saved API key.
type Profile struct {
	APIKey  string `json:"api_key,omitempty"`
	BaseURL string `json:"base_url,omitempty"`
	// ListenSecret signs the deliveries `yesly listen` forwards.
	ListenSecret string `json:"listen_secret,omitempty"`
}

// File is the whole config file.
type File struct {
	Profiles map[string]*Profile `json:"profiles"`
}

// Path is where the file lives. YESLY_CONFIG overrides it.
func Path() (string, error) {
	if p := os.Getenv("YESLY_CONFIG"); p != "" {
		return p, nil
	}
	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dir = filepath.Join(home, ".config")
	}
	return filepath.Join(dir, "yesly", "config.json"), nil
}

// Load reads the file; a missing file is an empty config.
func Load() (*File, error) {
	f := &File{Profiles: map[string]*Profile{}}
	p, err := Path()
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(p)
	if errors.Is(err, fs.ErrNotExist) {
		return f, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, f); err != nil {
		return nil, err
	}
	if f.Profiles == nil {
		f.Profiles = map[string]*Profile{}
	}
	return f, nil
}

// Save writes the file with mode 0600, in a directory with mode 0700.
func (f *File) Save() error {
	p, err := Path()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}
	tmp := p + ".tmp"
	if err := os.WriteFile(tmp, append(b, '\n'), 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, p)
}

// Profile returns the named profile, or nil.
func (f *File) Profile(name string) *Profile {
	return f.Profiles[name]
}

// Names returns the profile names, sorted.
func (f *File) Names() []string {
	names := make([]string, 0, len(f.Profiles))
	for n := range f.Profiles {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}
