package vscodeworkspace

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	ErrWorkspaceNotFound = fmt.Errorf("workspace not found")
)

//
// Interface
//

type (
	WorkspaceRepository interface {
		GetWorkspaceRepository
		GetByPathWorkspaceRepository
		ListWorkspaceRepository
		SaveWorkspaceRepository
		DeleteWorkspaceRepository
		SearchWorkspaceRepository
	}

	GetWorkspaceRepository interface {
		Get(name string) (Workspace, error)
	}

	GetByPathWorkspaceRepository interface {
		GetByPath(path string) (Workspace, error)
	}

	ListWorkspaceRepository interface {
		List() ([]Workspace, error)
	}

	SaveWorkspaceRepository interface {
		Save(ws Workspace) error
	}

	DeleteWorkspaceRepository interface {
		Delete(name string) error
	}

	SearchWorkspaceRepository interface {
		Search(query string) ([]Workspace, error)
	}
)

//
// Implementation
//

func NewWorkspaceRepositoryFile(path string) (*workspaceRepositoryFile, error) {
	wsr := &workspaceRepositoryFile{
		path:       path,
		workspaces: make(map[string]Workspace),
	}
	if _, err := os.Stat(path); err != nil {
		// ファイルが存在しない場合は空データを返す
		return wsr, nil
	}
	// ファイルが存在する場合はJSONを読み込んで返す
	if err := wsr.loadFile(); err != nil {
		return nil, err
	}

	return wsr, nil
}

type workspaceRepositoryFile struct {
	path       string
	workspaces map[string]Workspace
}

// Get implements WorkspaceRepository.
func (wsr *workspaceRepositoryFile) Get(name string) (Workspace, error) {
	if ws, ok := wsr.workspaces[name]; ok {
		return ws, nil
	}
	return Workspace{}, fmt.Errorf("%w: name='%s'", ErrWorkspaceNotFound, name)
}

// Get implements WorkspaceRepository.
func (wsr *workspaceRepositoryFile) GetByPath(path string) (Workspace, error) {
	for _, ws := range wsr.workspaces {
		if ws.Path == path {
			return ws, nil
		}
	}
	return Workspace{}, fmt.Errorf("%w: path='%s'", ErrWorkspaceNotFound, path)
}

// List implements WorkspaceRepository.
func (wsr *workspaceRepositoryFile) List() ([]Workspace, error) {
	wss := make([]Workspace, 0, len(wsr.workspaces))
	for _, ws := range wsr.workspaces {
		wss = append(wss, ws)
	}
	return wss, nil
}

// Save implements WorkspaceRepository.
func (wsr *workspaceRepositoryFile) Save(ws Workspace) error {
	// 追加（or上書き）
	wsr.workspaces[ws.Name] = ws

	// 保存
	if err := wsr.saveFile(); err != nil {
		return err
	}
	return nil
}

// Delete implements WorkspaceRepository.
func (wsr *workspaceRepositoryFile) Delete(name string) error {
	delete(wsr.workspaces, name)

	return nil
}

// Search implements WorkspaceRepository.
func (wsr *workspaceRepositoryFile) Search(query string) ([]Workspace, error) {
	// TODO: 検索機能
	return wsr.List()
}

var _ WorkspaceRepository = (*workspaceRepositoryFile)(nil)

type jsonData struct {
	Workspaces []Workspace `json:"workspaces"`
}

func (wsr *workspaceRepositoryFile) loadFile() error {
	f, err := os.OpenFile(wsr.path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	var data jsonData
	if err := dec.Decode(&data); err != nil {
		return nil
	}

	wss := make(map[string]Workspace)

	for _, ws := range data.Workspaces {
		wss[ws.Name] = ws
	}

	wsr.workspaces = wss

	return nil
}

func (wsr *workspaceRepositoryFile) saveFile() error {
	wss, err := wsr.List()
	if err != nil {
		return err
	}
	data := jsonData{
		Workspaces: wss,
	}

	// ファイルを開く
	f, err := os.OpenFile(wsr.path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("%w: here", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)

	if err := enc.Encode(&data); err != nil {
		return err
	}
	return nil
}
