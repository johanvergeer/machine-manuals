package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// configureInitManualCmd represents the configure command
var configureInitManualCmd = &cobra.Command{
	Use:   "init_manual",
	Short: "Set up a folder and file structure",
	Long: `This command sets up a predefined folder and file structure. 
You can specify the root directory where the structure should be created.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve flag values
		manufacturer, _ := cmd.Flags().GetString("manufacturer")
		machineName, _ := cmd.Flags().GetString("machine_name")
		manualName, _ := cmd.Flags().GetString("manual_name")

		// Call the function to create the structure
		err := setupStructure(manufacturer, machineName, manualName)
		if err != nil {
			fmt.Printf("Error setting up structure: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Folder and file structure configured successfully!")
	},
}

func init() {
	// Add the command to the root command
	rootCmd.AddCommand(configureInitManualCmd)

	// Define flags
	configureInitManualCmd.Flags().StringP("manufacturer", "m", "", "The machine manufacturer")
	configureInitManualCmd.MarkFlagRequired("manufacturer") // Mark as required

	configureInitManualCmd.Flags().StringP("machine_name", "n", "", "The machine name")
	configureInitManualCmd.MarkFlagRequired("machine_name") // Mark as required

	configureInitManualCmd.Flags().String("manual_name", "operator_manual", "The manual name (optional)")
}

// setupStructure creates the desired folder and file structure
func setupStructure(manufacturer string, machineName string, manualName string) error {
	// Define the folder and file structure
	manualsDir := getManualsDir()

	fmt.Println("Root: " + manualsDir)

	manualDir := filepath.Join(manualsDir, manufacturer, machineName, manualName)

	if folderExists(manualDir) {
		return fmt.Errorf("the manual directory already exists")
	}

	// Define the folder structure
	folderStructure := []string{
		filepath.Join(manualDir, "input", "images"),
		filepath.Join(manualDir, "input", "ocr_output"),
		filepath.Join(manualDir, "input", "raw_pdf"),
		filepath.Join(manualDir, "latex", "images"),
		filepath.Join(manualDir, "latex", "chapters"),
	}

	// Create each folder and add a .gitkeep if it's empty
	for _, folder := range folderStructure {
		// Create the folder
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder %s: %w", folder, err)
		}

		// Add a .gitkeep file to empty folders
		gitkeepPath := filepath.Join(folder, ".gitkeep")
		file, err := os.Create(gitkeepPath)
		if err != nil {
			return fmt.Errorf("failed to create .gitkeep in %s: %w", folder, err)
		}
		file.Close() // Close the file after creation
	}

	// Create main.tex file
	mainTexPath := filepath.Join(manualDir, "latex", "main.tex")
	_, _ = os.Create(mainTexPath)

	// Create chatgpt_rules.md file with default content
	chatgptRulesPath := filepath.Join(manualDir, "latex", "chatgpt_rules.md")
	err := createChatGPTRulesFile(chatgptRulesPath)
	if err != nil {
		return fmt.Errorf("failed to create chatgpt_rules.md: %w", err)
	}

	fmt.Printf("Directory structure created successfully at: %s\n", manualsDir)
	return nil
}

// folderExists checks if a folder exists at the specified path
func folderExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		fmt.Printf("Error checking folder: %v\n", err)
		return false
	}
	return info.IsDir()
}

// createChatGPTRulesFile writes the default rules into chatgpt_rules.md
func createChatGPTRulesFile(filePath string) error {
	defaultContent := `# ChatGPT Processing Rules for This Manual

This file defines specific **translation and formatting rules** for ChatGPT when processing this manual.

## **General Translation Rules**
- Translate all text **to English**.
- Keep technical terms consistent with machinist terminology.
- Do **not** add or invent text — only translate what is present.

## **Formatting Conventions**
- Convert **underlined text** to \textbf{} (bold) in LaTeX.

## **Page-Specific Handling**
- If a page contains only a table, convert it to a tabularx LaTeX table.
- If a page has hand-drawn annotations, **omit them** unless they contain relevant data.

---

This file can be modified for **manual-specific instructions**.
`

	// Write content to file
	return os.WriteFile(filePath, []byte(defaultContent), 0644)
}
