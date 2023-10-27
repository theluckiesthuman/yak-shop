package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/theluckiesthuman/yakshop/internal/dto"
	query "github.com/theluckiesthuman/yakshop/internal/usecase/implementation"
)

func main() {
	rootCmd := setupRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}

func setupRootCommand() *cobra.Command {
	var filePath string
	var elapsedTime int

	rootCmd := &cobra.Command{
		Use:   "yakshop",
		Short: "A program to manage a yak shop",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if filePath == "" {
				return fmt.Errorf("file path must be provided")
			}
			if elapsedTime < 0 {
				return fmt.Errorf("elapsed time must be non-negative")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			calculateAndDisplayStockAndAge(filePath, elapsedTime)
		},
	}

	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the XML file")
	rootCmd.Flags().IntVarP(&elapsedTime, "time", "T", 0, "Elapsed time in days")

	return rootCmd
}

func calculateAndDisplayStockAndAge(filePath string, elapsedTime int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening XML file: %v", err)
	}
	defer file.Close()

	fq, err := query.NewFileQuery(file)
	if err != nil {
		log.Fatalf("Error creating file query: %v", err)
	}
	stock, err := fq.CalculateStock(elapsedTime)
	if err != nil {
		log.Fatalf("Error calculating stock: %v", err)
	}

	herd, err := fq.CalculateAge(elapsedTime)
	if err != nil {
		log.Fatalf("Error calculating age: %v", err)
	}

	displayOutput(stock, herd)
}

func displayOutput(stock *dto.Stock, herd *dto.Herd) {
	fmt.Printf("In Stock:\n\t%.3f liters of milk\n\t%d skins of wool\n", stock.Milk, stock.Skins)
	fmt.Println("Herd:")
	for _, yak := range herd.Yaks {
		if yak.Age >= 10 {
			fmt.Printf("\t%s is dead\n", yak.Name)
			continue
		}
		fmt.Printf("\t%s %.2f years old\n", yak.Name, yak.Age)
	}
}
