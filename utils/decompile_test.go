package utils

import (
	"testing"
	"fmt"
)

func TestGetComponents(t *testing.T) {
	apkPath := "/mnt/git/fengge/uploads/20250407110208_rotating_cube-release.apk"
	
	componentPath := GetComponents(apkPath)
	fmt.Printf("Component path: %s\n", componentPath)
	// if len(components) != len(expectedComponents) {
	// 	t.Errorf("Expected %d components, got %d", len(expectedComponents), len(components))
	// }

	// for i, component := range components {
	// 	if component != expectedComponents[i] {
	// 		t.Errorf("Expected component %s, got %s", expectedComponents[i], component)
	// 	}
	// }
}