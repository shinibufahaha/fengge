package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetComponents(apkPath string) string {
	configmap := LoadConfig("../config.yaml")
	if configmap == nil {
		log.Println("Failed to load config.yaml")
		return ""
	}
	outPath := ""
	apkName := filepath.Base(apkPath)
	apkNameWithoutExt := strings.TrimSuffix(apkName, filepath.Ext(apkName))
	outDir := filepath.Join(configmap.Path.ProjectPath, "files", "decompile", apkNameWithoutExt)

	// Create output directory if it doesn't exist
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err := os.MkdirAll(outDir, os.ModePerm)
		if err != nil {
			log.Printf("Error creating output directory: %v\n", err)
			return ""
		}
	}
	cmd := fmt.Sprintf("%s -d %s --show-bad-code %s", configmap.Path.JadxPath, outDir, apkPath)
	log.Printf(cmd)
	err := exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		log.Printf("Error decompiling APK: %v\n", err)
		return ""
	}
	log.Printf("Decompiled APK successfully at %s\n", outDir)
	
	// Scan all files in outDir to find AndroidManifest.xml
	err = filepath.Walk(outDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Base(path) == "AndroidManifest.xml" {
			log.Printf("Found AndroidManifest.xml at %s\n", path)
			
			// Replace AndroidManifest.xml with components.json as output path
			outPath = strings.Replace(path, "AndroidManifest.xml", "components.json", -1)
			
			// Parse AndroidManifest.xml to find components
			pyScript := filepath.Join(configmap.Path.ProjectPath, "python", "parseManifest.py")
			pyCmd := fmt.Sprintf("python3 %s %s %s", pyScript, path, outPath)
			fmt.Printf("Executing command: %s\n", pyCmd)
			err := exec.Command("sh", "-c", pyCmd).Run()
			if err != nil {
				log.Printf("Error parsing AndroidManifest.xml: %v\n", err)
			}
			return nil
		}
		return nil
	})
	if err != nil {
		log.Printf("Error finding AndroidManifest.xml: %v\n", err)
		return ""
	}

	// Scan all files in outDir to find classes.dex
	// (Implementation for this part is missing in the provided code)

	return outPath
}