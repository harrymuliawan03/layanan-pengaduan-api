package scaffoldcommand

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/cmd/cli/commands/stubs"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
)

var ScaffoldCommand = &cobra.Command{
	Use:   "make:scaffold <file_name>",
	Short: "Create new scaffold",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		if fileName == "" {
			cmd.Help()
			return
		}
		// Open the go.mod file
		file, err := os.ReadFile("go.mod")
		if err != nil {
			fmt.Println("Error reading go.mod:", err)
			return
		}

		// Parse the go.mod file
		modFile, err := modfile.Parse("go.mod", file, nil)
		if err != nil {
			fmt.Println("Error parsing go.mod:", err)
			return
		}

		// Use a regular expression to insert an underscore before each capital letter
		re := regexp.MustCompile(`([A-Z])`)
		// Replace all matches with an underscore followed by the lowercase letter
		outputFileName := re.ReplaceAllString(fileName, "_$1")
		// Convert the entire string to lowercase and remove the leading underscore if any
		outputFileName = strings.TrimPrefix(strings.ToLower(outputFileName), "_")

		fileName = strings.ToLower(outputFileName)

		modPath := modFile.Module.Mod.Path
		CreateRepository(fileName, modPath)
		CreateRepositoryImpl(fileName, modPath)
		CreateRequest(fileName, modPath)
		CreateDto(fileName, modPath)
		CreateService(fileName, modPath)
		CreateServiceImpl(fileName, modPath)
		// createService(fileName)

		titleMessage := `
Success!
Scaffold command was create:
`
		bodyMessage := fmt.Sprintf(`
	- app/repositories/%v_repo/%v_repository.go
	- app/repositories/%v_repo/%v_repository_impl.go
	- app/http/requests/%v_request.go
	- app/dto/%v_data.go
	- app/services/%v_service/%v_service.go
	- app/services/%v_service/%v_service_impl.go
`, fileName, fileName, fileName, fileName, fileName, fileName, fileName, fileName, fileName, fileName)
		color.Greenln(titleMessage)
		color.Yellowln(bodyMessage)
	},
}

func CreateRepository(fileName string, modPath string) {
	basePath := fmt.Sprintf("app/repositories/%v_repo", fileName)
	err := os.Mkdir(basePath, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName1 := fmt.Sprintf("%v/%v_repository.go", basePath, fileName)
	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(stubs.CreateRepository(fileName, modPath))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateRepositoryImpl(fileName string, modPath string) {
	basePath := fmt.Sprintf("app/repositories/%v_repo", fileName)
	fileName1 := fmt.Sprintf("%v/%v_repository_impl.go", basePath, fileName)
	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(stubs.CreateRepositoryImpl(fileName, modPath))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateRequest(fileName string, modPath string) {
	basePath := "app/http/requests"
	fileName1 := fmt.Sprintf("%v/%v_request.go", basePath, fileName)
	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		return
	}
	operations := []string{"create", "update"}

	_, err = f.WriteString(stubs.CreateRequest(fileName, operations))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateDto(fileName string, modPath string) {
	basePath := "app/dto"
	fileName1 := fmt.Sprintf("%v/%v_data.go", basePath, fileName)
	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(stubs.CreateDto(fileName))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateService(fileName string, modPath string) {
	basePath := fmt.Sprintf("app/services/%v_service", fileName)
	err := os.Mkdir(basePath, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName1 := fmt.Sprintf("%v/%v_service.go", basePath, fileName)
	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(stubs.CreateService(fileName, modPath))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateServiceImpl(fileName string, modPath string) {
	basePath := fmt.Sprintf("app/services/%v_service", fileName)
	fileName1 := fmt.Sprintf("%v/%v_service_impl.go", basePath, fileName)
	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(stubs.CreateServiceImpl(fileName, modPath))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
