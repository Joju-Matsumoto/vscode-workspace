package cmd

import "github.com/Joju-Matsumoto/vscode-workspace/vscodeworkspace"

func NewUsecase(path string) (vscodeworkspace.Usecase, error) {
	repository, err := vscodeworkspace.NewWorkspaceRepositoryFile(path)
	if err != nil {
		return nil, err
	}
	executer := vscodeworkspace.NewExecuter()
	usecase := vscodeworkspace.NewUsecase(repository, executer)

	return usecase, nil
}
