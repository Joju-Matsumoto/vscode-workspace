package vscodeworkspace

import (
	"errors"
	"fmt"
	"path/filepath"
)

var (
	ErrWorkspaceAlreadyExist = fmt.Errorf("workspace already exist")
)

type (
	Usecase interface {
		GetWorkspaceUsecase
		ListWorkspaceUsecase
		AddWorkspaceUsecase
		RenameWorkspaceUsecase
		DeleteWorkspaceUsecase
		SearchWorkspaceUsecase
		SearchWorkspaceFromDirectoryUsecase
		InitWorkspaceUsecase
		OpenWorkspaceUsecase
	}

	GetWorkspaceUsecase interface {
		GetWorkspace(name string) (Workspace, error)
	}

	ListWorkspaceUsecaseOption struct {
		SortBy string // default: name
	}

	ListWorkspaceUsecase interface {
		ListWorkspace(opt ListWorkspaceUsecaseOption) ([]Workspace, error)
	}

	AddWorkspaceUsecase interface {
		AddWorkspace(name string, path string) error
	}

	RenameWorkspaceUsecase interface {
		RenameWorkspace(oldName string, newName string) error
	}

	DeleteWorkspaceUsecase interface {
		DeleteWorkspace(name string) error
	}

	SearchWorkspaceUsecase interface {
		SearchWorkspace(query string, opt ListWorkspaceUsecaseOption) ([]Workspace, error)
	}

	SearchWorkspaceFromDirectoryUsecase interface {
		SearchWorkspaceFromDirectory(baseDir string, query string) ([]Workspace, error)
	}

	InitWorkspaceUsecase interface {
		InitWorkspace(name string, dir string) (Workspace, error)
	}

	OpenWorkspaceUsecase interface {
		OpenWorkspace(name string) error
	}
)

func NewUsecase(repository WorkspaceRepository, executer Executer) *usecase {
	return &usecase{
		repository: repository,
		executer:   executer,
	}
}

type usecase struct {
	repository WorkspaceRepository
	executer   Executer
}

// GetWorkspace implements Usecase.
func (u *usecase) GetWorkspace(name string) (Workspace, error) {
	return u.repository.Get(name)
}

// AddWorkspace implements Usecase.
func (u *usecase) AddWorkspace(name string, path string) error {
	ws, err := NewWorkspace(name, path)
	if err != nil {
		return err
	}

	if ws, err := u.repository.Get(ws.Name); !errors.Is(err, ErrWorkspaceNotFound) {
		if err == nil {
			// 見つかった場合はエラー
			return fmt.Errorf("%w: name='%s', path='%s'", ErrWorkspaceAlreadyExist, ws.Name, ws.Path)
		}
		return err
	}

	if err := u.repository.Save(ws); err != nil {
		return err
	}

	return nil
}

// ListWorkspace implements Usecase.
func (u *usecase) ListWorkspace(opt ListWorkspaceUsecaseOption) ([]Workspace, error) {
	wss, err := u.repository.List(ListWorkspaceRepositoryOption{
		SortBy: opt.GetSortBy(),
	})
	if err != nil {
		return nil, err
	}

	return wss, nil
}

// RenameWorkspace implements Usecase.
func (u *usecase) RenameWorkspace(oldName string, newName string) error {
	ws, err := u.repository.Get(oldName)
	if err != nil {
		return err
	}

	if _, err := u.repository.Get(newName); !errors.Is(err, ErrWorkspaceNotFound) {
		// newNameが既に存在する場合はエラー
		return err
	}

	if err := u.repository.Delete(ws.Name); err != nil {
		return err
	}

	ws.Name = newName

	if err := ws.Validate(); err != nil {
		return err
	}

	if err := u.repository.Save(ws); err != nil {
		return err
	}

	return nil
}

// DeleteWorkspace implements Usecase.
func (u *usecase) DeleteWorkspace(name string) error {
	if _, err := u.repository.Get(name); err != nil {
		return err
	}

	if err := u.repository.Delete(name); err != nil {
		return err
	}
	return nil
}

// SearchWorkspace implements Usecase.
func (u *usecase) SearchWorkspace(query string, opt ListWorkspaceUsecaseOption) ([]Workspace, error) {
	return u.repository.Search(query, ListWorkspaceRepositoryOption{
		SortBy: opt.GetSortBy(),
	})
}

// SearchWorkspaceFromDirectory implements Usecase.
func (u *usecase) SearchWorkspaceFromDirectory(baseDir string, query string) ([]Workspace, error) {
	return SearchWorkspacesFromBaseDirectory(baseDir, query)
}

// InitWorkspace implements Usecase.
func (u *usecase) InitWorkspace(name string, dir string) (Workspace, error) {
	var err error
	dir, err = filepath.Abs(dir)
	if err != nil {
		return Workspace{}, err
	}
	if len(name) == 0 {
		name = filepath.Base(dir)
	}
	path, err := u.executer.Init(name, dir)
	if err != nil {
		return Workspace{}, err
	}

	ws, err := NewWorkspace(name, path)
	if err != nil {
		return Workspace{}, err
	}
	return ws, nil
}

// OpenWorkspace implements Usecase.
func (u *usecase) OpenWorkspace(name string) error {
	ws, err := u.repository.Get(name)
	if err != nil {
		return err
	}

	if err := u.executer.Open(ws.Path); err != nil {
		return err
	}

	ws.Open()

	if err := u.repository.Save(ws); err != nil {
		return err
	}
	return nil
}

var _ Usecase = (*usecase)(nil)

func (opt *ListWorkspaceUsecaseOption) GetSortBy() SortBy {
	switch opt.SortBy {
	case "opened_at":
		return SORTBY_OPENEDAT
	case "count":
		return SORTBY_COUNT
	}
	return SORTBY_NAME
}
