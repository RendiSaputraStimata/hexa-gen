/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"embed"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var templateFS embed.FS

const (
	contractRepository = "templates/contract_repository.go.tpl"
	contractService    = "templates/contract_service.go.tpl"
	impl_service       = "templates/service_impl.go.tpl"
	imple_repository   = "templates/repository_impl.go.tpl"
	entity             = "templates/domain.go.tpl"
	request            = "templates/request.go.tpl"
	response           = "templates/response.go.tpl"
)

var domainName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hexa.gen",
	Short: "Generate a new domain",
	Long:  `Generate repository, request, response structs, contracts for a new domain.`,
	Run: func(cmd *cobra.Command, args []string) {
		if domainName == "" {
			fmt.Println("Please provide a domain name")
			return
		}

		nameSpace, err := GetModuleName()
		if err != nil {
			fmt.Printf("Error getting module name: %v\n", err)
			return
		}

		generateDomain(domainName, nameSpace)
	},
}

type FileArgument struct {
	Domain    string
	NameSpace string
}

// generateDomain generates a domain
func generateDomain(domain, module string) {
	files := map[string]string{
		"./internal/domain/repository/" + strings.ToLower(domain) + "_repository.go":         contractRepository,
		"./internal/application/port/" + strings.ToLower(domain) + "_service.go":             contractService,
		"./internal/domain/entity/" + strings.ToLower(domain) + ".go":                        entity,
		"./internal/application/port/in/" + strings.ToLower(domain) + "_request.go":          request,
		"./internal/application/port/out/" + strings.ToLower(domain) + "_response.go":        response,
		"./internal/adapters/repository/mongo/" + strings.ToLower(domain) + "_repository.go": imple_repository,
		"./internal/application/service/" + strings.ToLower(domain) + "_service.go":          impl_service,
	}

	for fileName, templatePath := range files {
		createFileFromTemplate(domain, fileName, module, templatePath)
	}
}

// createFileFromTemplate creates a file from a template
func createFileFromTemplate(domain, fileName, module, templatePath string) {

	tmplContent, err := templateFS.ReadFile(templatePath)
	if err != nil {
		fmt.Printf("Error reading template: %v\n", err)
		return
	}

	// Open the template file
	tmpl, err := template.New("template").Parse(string(tmplContent))
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file
	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Data to pass to the template
	data := FileArgument{
		Domain:    domain,
		NameSpace: module,
	}

	// Execute the template and write to the file
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// getModuleName reads the module name from the go.mod file
func GetModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}
	}

	return "", fmt.Errorf("module name not found in go.mod")
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hexa.gen.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&domainName, "domain", "d", "", "Domain name (required)")
	err := rootCmd.MarkFlagRequired("domain")

	if err != nil {
		panic(err)
	}
}
