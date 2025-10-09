package tests

import (
	"bytes"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var update = flag.Bool("update", false, "update golden files")

// TestCodeGeneration tests that the protoc plugin generates the expected code
func TestCodeGeneration(t *testing.T) {
	// Build the plugin first
	cmd := exec.Command("go", "install", ".")
	cmd.Dir = filepath.Join("..")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build plugin: %v\n%s", err, output)
	}

	// Create output directory
	if err := os.MkdirAll("out", 0755); err != nil {
		t.Fatalf("Failed to create output directory: %v", err)
	}

	// Run protoc to generate code
	cmd = exec.Command("protoc",
		"--proto_path", ".",
		"-I=.",
		"-I=../graphql",
		"-I=../example",
		"--go_out=paths=source_relative:out",
		"--go-grpc_out=paths=source_relative:out",
		"--graphql_out=lang=go,paths=source_relative:out",
		"cases/messages.proto")
	cmd.Dir = filepath.Join(".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run protoc: %v\n%s", err, output)
	}
	t.Logf("protoc output: %s", output)

	// Read generated file
	generatedPath := filepath.Join("out", "cases", "messages.graphql.go")
	generated, err := os.ReadFile(generatedPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Path to golden file
	goldenPath := filepath.Join("golden", "messages.graphql.go")

	if *update {
		// Update golden file
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
			t.Fatalf("Failed to create golden directory: %v", err)
		}
		if err := os.WriteFile(goldenPath, generated, 0644); err != nil {
			t.Fatalf("Failed to write golden file: %v", err)
		}
		t.Log("Updated golden file")
		return
	}

	// Read golden file
	golden, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file (run with -update to create): %v", err)
	}

	// Compare
	if !bytes.Equal(golden, generated) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(string(golden), string(generated), false)
		t.Errorf("Generated code does not match golden file:\n%s", dmp.DiffPrettyText(diffs))
	}
}

// TestMain ensures cleanup after tests
func TestMain(m *testing.M) {
	flag.Parse()

	// Run tests
	code := m.Run()

	// Cleanup output directory (but not if updating golden files)
	if !*update {
		os.RemoveAll("out")
	}

	os.Exit(code)
}
