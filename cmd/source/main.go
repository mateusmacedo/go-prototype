package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mateusmacedo/govibranium/prototype/internal/application/service"
	"github.com/mateusmacedo/govibranium/prototype/internal/core/err"
)

type OSValidator struct {
}

func (v *OSValidator) Validate(target any) error {
	switch target := target.(type) {
		case string:
			if target == "" {
				return err.ErrorFactory(service.InvalidOSFileSourcePathErrMsg, target)
			}
			if stat, _err := os.Stat(target); os.IsNotExist(_err) || stat.IsDir() {
				return err.ErrorFactory(service.InvalidOSFileSourcePathErrMsg, target)
			}
		default:
			return err.ErrorFactory(service.InvalidOSFileSourcePathErrMsg, target)
	}
	return nil
}

type OSSourceAdapter struct {
}

func (a *OSSourceAdapter) Adapt(target any) (interface{}, error) {
	switch target := target.(type) {
		case string:
			file, _err := os.Open(target)
			if _err != nil {
				return nil, err.ErrorFactory(service.InvalidOSFileSourcePathErrMsg, target)
			}
			return file, nil
	}

	return nil, err.ErrorFactory(service.InvalidOSFileSourcePathErrMsg, target)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório atual:", err)
		return
	}
	filePath := filepath.Join(dir, "data/dummy_file.txt")
	validator := &OSValidator{}
	adapter := &OSSourceAdapter{}

	source, err := service.NewOSFileSource(
		service.WithOSFileSourcePath(filePath),
		service.WithOSFileSourceValidators(validator),
		service.WithOSFileSourceAdapter(adapter),
	)

	if err != nil {
		panic(err)
	}

	sourceOpened, err := source.Open()
	if err != nil {
		panic(err)
	}

	if file, ok := sourceOpened.(*os.File); !ok {
		fmt.Println("Erro ao tentar converter o arquivo para *os.File")
	} else {
		defer file.Close()
		for {
			b := make([]byte, 1)
			_, err := file.Read(b)
			if err != nil {
				break
			}
			fmt.Print(string(b))
		}
	}

	fmt.Println("Arquivo aberto com sucesso!")

}
