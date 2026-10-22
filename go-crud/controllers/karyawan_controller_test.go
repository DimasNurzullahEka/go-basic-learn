package controllers

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetFormLoadsFromDifferentWorkingDirectory(t *testing.T) {
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	}()

	response := httptest.NewRecorder()
	GetForm(response, httptest.NewRequest("GET", "/karyawan/get-form", nil))

	if response.Code != 200 {
		t.Fatalf("expected status 200, got %d: %s", response.Code, response.Body.String())
	}
	if !strings.Contains(response.Body.String(), `name="nama_karyawan"`) {
		t.Fatalf("form response does not contain the employee name field: %s", response.Body.String())
	}
}
