package main_test

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task number 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-task", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := task + "\n"
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := ""
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})

	t.Run("NoFlags_PrintsCustomErrorAndExits1", func(t *testing.T) {
		cmd := exec.Command(cmdPath) // No args --> Default branch in the switch
		out, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatalf("Expected non-zero exit, got nit error; output=%q", string(out))
		}
		var ee *exec.ExitError
		if !errors.As(err, &ee) {
			t.Fatalf("expected *exec.ExitError, got %T instead", err)
		}
		if code := ee.ExitCode(); code != 1 {
			t.Fatalf("expected exit code 1, got %d; output=%q", code, string(out))
		}
		if !strings.Contains(string(out), "invalid option") {
			t.Fatalf("expected output to contain %q, got %q instead", "invalid option", string(out))
		}
	})

	t.Run("UnknownFlag_Exit2FromFlagParse", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-invalidFlag", "x")
		out, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatalf("expected non-zero exit, got nil error; output=%q", string(out))
		}
		var ee *exec.ExitError
		if !errors.As(err, &ee) {
			t.Fatalf("expected *exec.ExitError, got %T instead.", err)
		}
		if code := ee.ExitCode(); code != 2 {
			t.Fatalf(
				"expected exit code 2 from flag.Parse, got %d instead; output=%q",
				code,
				string(out),
			)
		}
		if !strings.Contains(string(out), "flag provided but not defined") {
			t.Fatalf("expected flag.Parse error text, got %q instead", string(out))
		}
	})
}
