/** This package creates a new package to your golang project */
package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// CreatePackage creates a new package to your golang project
func CreatePackage(packageName, folder string) {
	if packageName == "" {
		fmt.Println("please provide a package name")
		os.Exit(1)
	}
	packageName = strings.ToLower(packageName)

	folder = filepath.Join(folder, packageName)
	os.MkdirAll(folder, os.ModePerm)

	// Create the main.go file
	mainFile := filepath.Join(folder, packageName+".go")
	f, err := os.Create(mainFile)
	if err != nil {
		fmt.Println("error creating main.go file", err)
		os.Exit(1)
	}

	fmt.Println("creating package ", packageName, " in ", folder)

	// Write the main.go file
	f.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	f.WriteString("import (\n")
	f.WriteString("\t\"fmt\"\n")
	f.WriteString(")\n\n")

	f.WriteString("func main() {\n")
	f.WriteString("\tfmt.Println(\"Hello, World!\")\n")
	f.WriteString("}\n")

	f.Close()

	// create the main_test.go file
	testFile := filepath.Join(folder, packageName+"_test.go")
	f, err = os.Create(testFile)
	if err != nil {
		fmt.Println("error creating main_test.go file", err)
		os.Exit(1)
	}

	// get module name from go.mod
	cmd := exec.Command("go", "list", "-m")
	cmd.Dir = folder
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("error getting module name", err)
		os.Exit(1)
	}

	modulePath := path.Join(strings.TrimSpace(string(out)), folder)
	fmt.Println("creating tests for ", modulePath)

	// Write the main_test.go file
	f.WriteString(fmt.Sprintf("package %s_test\n\n", packageName))
	f.WriteString("import (\n")
	f.WriteString(fmt.Sprintf("\t\"%s\"\n", modulePath))
	f.WriteString("\t\"testing\"\n")
	f.WriteString(")\n\n")

	f.WriteString("func TestMain(t *testing.T) {\n")
	f.WriteString("\t// Add your test here\n")
	f.WriteString("}\n")

	f.Close()
}
