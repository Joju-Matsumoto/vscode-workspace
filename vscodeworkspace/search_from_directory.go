package vscodeworkspace

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Option func(opt *opt)

type opt struct {
	ext string
}

func newOption() *opt {
	return &opt{
		ext: ".code-workspace",
	}
}

func WithExt(ext string) Option {
	return func(opt *opt) {
		opt.ext = ext
	}
}

func SearchWorkspacesFromBaseDirectory(baseDir string, query string, options ...Option) ([]Workspace, error) {
	// TODO: クエリを無視しない

	opt := newOption()

	for _, f := range options {
		f(opt)
	}

	wss := make([]Workspace, 0)

	if err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() || filepath.Ext(d.Name()) != opt.ext {
			return nil
		}
		name := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		ws, err := NewWorkspace(name, path)
		if err != nil {
			return err
		}
		wss = append(wss, ws)
		return nil
	}); err != nil {
		return nil, err
	}

	return wss, nil
}
