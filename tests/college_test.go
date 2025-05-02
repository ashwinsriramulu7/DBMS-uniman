package tests

import (
	"testing"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
)
func TestAddCollege(t *testing.T) {
	testData := map[string]any{
		"name":     "Test College",
		"location": "Test City",
		"estd":     2025,
	}
	err := modules.AddCollege(testData)
	if err != nil {
		t.Errorf("AddCollege failed: %v", err)
	}
}

