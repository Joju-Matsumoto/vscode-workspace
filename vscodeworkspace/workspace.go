package vscodeworkspace

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

var (
	ErrWorkspaceValidation = fmt.Errorf("workspace validation error")
)

type Workspace struct {
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	OpenedAt time.Time `json:"opened_at"`
	Count    int       `json:"count"`
}

func (w *Workspace) Open() {
	w.OpenedAt = time.Now()
	w.Count += 1
}

func NewWorkspace(name string, path string) (Workspace, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return Workspace{}, err
	}
	if name == "" {
		name = strings.TrimSuffix(filepath.Base(path), WorkspaceFileExt)
	}
	ws := Workspace{
		Name:  name,
		Path:  abs,
		Count: 0,
	}
	if err := ws.Validate(); err != nil {
		return Workspace{}, err
	}
	return ws, nil
}

func (ws *Workspace) Validate() error {
	if len(ws.Name) == 0 {
		return fmt.Errorf("%w: empty name", ErrWorkspaceValidation)
	}
	if len(ws.Path) == 0 {
		return fmt.Errorf("%w: empty path", ErrWorkspaceValidation)
	}
	return nil
}
