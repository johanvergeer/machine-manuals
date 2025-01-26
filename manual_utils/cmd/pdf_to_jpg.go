package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/go-fitz"
	"github.com/spf13/cobra"
)

// configureConvertPdfToJpgCmd represents the configure command
var configureConvertPdfToJpgCmd = &cobra.Command{
	Use:   "pdf_to_jpg",
	Short: "Convert the manual pdf files to jpeg files.",
	Long:  `This command takes all the files from input/raw_pdf, converts them to jpg files and saves them in input/images.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve flag values
		manufacturer, _ := cmd.Flags().GetString("manufacturer")
		machineName, _ := cmd.Flags().GetString("machine_name")
		manualName, _ := cmd.Flags().GetString("manual_name")

		inputDir := filepath.Join(getManualsDir(), manufacturer, machineName, manualName, "input")
		pdfsDir := filepath.Join(inputDir, "raw_pdf")
		imagesDir := filepath.Join(inputDir, "images")

		// Call the function to create the structure
		err := processPDFs(pdfsDir, imagesDir)
		if err != nil {
			fmt.Printf("Error converting pdfs to jpeg: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Folder and file structure configured successfully!")
	},
}

func init() {
	// Add the command to the root command
	rootCmd.AddCommand(configureConvertPdfToJpgCmd)

	// Define flags
	configureConvertPdfToJpgCmd.Flags().StringP("manufacturer", "m", "", "The machine manufacturer")
	configureConvertPdfToJpgCmd.MarkFlagRequired("manufacturer") // Mark as required

	configureConvertPdfToJpgCmd.Flags().StringP("machine_name", "n", "", "The machine name")
	configureConvertPdfToJpgCmd.MarkFlagRequired("machine_name") // Mark as required

	configureConvertPdfToJpgCmd.Flags().String("manual_name", "operator_manual", "The manual name (optional)")
}

func processPDFs(pdfDir, outputDir string) error {
	// Ensure the input directory exists
	if _, err := os.Stat(pdfDir); os.IsNotExist(err) {
		return fmt.Errorf("input directory does not exist: %s", pdfDir)
	}

	// Get all PDF files in the input directory
	pdfFiles, err := filepath.Glob(filepath.Join(pdfDir, "*.pdf"))
	if err != nil {
		return fmt.Errorf("failed to list PDF files: %w", err)
	}

	// Check if there are PDF files
	if len(pdfFiles) == 0 {
		return fmt.Errorf("no PDF files found in directory: %s", pdfDir)
	}

	// Process each PDF file
	for _, pdfFile := range pdfFiles {
		fmt.Printf("Processing file: %s\n", pdfFile)

		// Extract the file name without extension
		pdfName := strings.TrimSuffix(filepath.Base(pdfFile), filepath.Ext(pdfFile))
		pdfOutputDir := filepath.Join(outputDir, pdfName)

		// Convert the PDF to images
		err := convertPDFToImages(pdfFile, pdfOutputDir)
		if err != nil {
			return fmt.Errorf("failed to process PDF %s: %w", pdfFile, err)
		}
	}

	return nil
}

func convertPDFToImages(pdfPath, outputDir string) error {
	// Open the PDF document
	doc, err := fitz.New(pdfPath)
	if err != nil {
		return fmt.Errorf("failed to open PDF: %w", err)
	}
	defer doc.Close()

	// Ensure the output directory exists
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Convert each page to a JPEG image
	for i := 0; i < doc.NumPage(); i++ {
		// Render the page
		img, err := doc.Image(i)
		if err != nil {
			return fmt.Errorf("failed to render page %d: %w", i, err)
		}

		// Define the output file path
		outputPath := filepath.Join(outputDir, fmt.Sprintf("%d.jpg", i+1))

		// Save the image as a JPEG
		file, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("failed to create JPEG file: %w", err)
		}

		err = writeJpeg(file, img)
		if err != nil {
			file.Close()
			return fmt.Errorf("failed to save JPEG: %w", err)
		}

		file.Close()
		fmt.Printf("Page %d converted to: %s\n", i+1, outputPath)
	}

	return nil
}

// writeJpeg encodes the image as a JPEG and writes it to the file.
func writeJpeg(file *os.File, img image.Image) error {
	// Set JPEG quality (optional, default is 75 if omitted)
	options := &jpeg.Options{Quality: 90}

	// Encode the image to the file
	err := jpeg.Encode(file, img, options)
	if err != nil {
		return err
	}

	return nil
}
