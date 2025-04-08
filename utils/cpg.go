package utils

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"path/filepath"
)

func GenerateCPG(filePath string) {
	fileName := filepath.Base(filePath)
	cpgName := fileName[:len(fileName)-len(filepath.Ext(fileName))] + ".cpg"
	currentDir, _ := os.Getwd()
	cpgPath := currentDir + "/cpgs/" + cpgName
	
	cmd := fmt.Sprintf("/mnt/git/joern/joern-cli/joern-parse --language java -o %s %s", cpgPath, filePath)
	log.Printf(cmd)
	err := exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		fmt.Printf("Error generating CPG: %v\n", err)
		return
	}
	fmt.Printf("CPG generated successfully at %s\n", cpgPath)
}