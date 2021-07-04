package logger

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func TestFatal(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		Fatal("fatal testing")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")

	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestError(t *testing.T) {
	Error("error testing")
}

func TestWarning(t *testing.T) {
	Warning("warning testing")
}

func TestInfo(t *testing.T) {
	Info("info testing")
}

func TestSuccess(t *testing.T) {
	Success("success testing")
}

func TestLog(t *testing.T) {
	Log("log testing")
}

func TestWriteLog(t *testing.T) {
	tmpdir := t.TempDir()

	SetLogLocation(tmpdir)
	if logLocation != tmpdir {
		t.Fatalf("log location variable has been not updated")
	}

	write("test write %s", "log")

	f, err := os.Stat(filepath.Join(
		logLocation,
		fmt.Sprintf(
			"%s.log",
			time.Now().Format("20060102"),
			),
		))
	if err != nil {
		t.Fatal(err)
	}

	if f.Size() == 0 {
		t.Fatalf("log file size can't be null ?")
	}
}

func TestSetLogLocationNotDir(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		SetLogLocation("test/fake_file.txt")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestSetLogLocationNotDir")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")

	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("process ran with err %v, want exit status 1", err)
}

