package vscodeworkspace

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	WorkspaceFileExt = ".code-workspace"
)

var (
	ErrWorkspaceConfigFileAlreadyExist = fmt.Errorf("workspace config file already exist")
)

type (
	Executer interface {
		Open(path string) error
		Init(name string, dir string) (path string, err error)
	}
)

func NewExecuter() *executer {
	return &executer{}
}

type executer struct{}

// Open implements Executer.
func (e *executer) Open(path string) error {
	return exec.Command("code", path).Run()
}

// Init implements Executer.
func (e *executer) Init(name string, dir string) (string, error) {
	savePath := filepath.Join(dir, ".vscode", fmt.Sprintf("%s%s", name, WorkspaceFileExt))
	// 既に存在する場合はエラー
	if _, err := os.Stat(savePath); err == nil {
		return "", fmt.Errorf("%w: path='%s'", ErrWorkspaceConfigFileAlreadyExist, savePath)
	}

	// Workspace設定ファイルの作成
	cfg := NewWorkspaceConfig("..")

	// JSONを保存
	if err := saveJson(savePath, cfg); err != nil {
		return "", err
	}
	return savePath, nil
}

var _ Executer = (*executer)(nil)

//
// workspace config
//

type workspaceConfig struct {
	Folders []folder `json:"folders"`
}

type folder struct {
	Path string `json:"path"`
}

func NewWorkspaceConfig(path string) workspaceConfig {
	return workspaceConfig{
		Folders: []folder{
			{
				Path: path,
			},
		},
	}
}

//
// JSON util
//

func saveJson(path string, data any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)

	return enc.Encode(data)
}
