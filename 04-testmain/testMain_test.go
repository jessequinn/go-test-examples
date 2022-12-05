package testMain

import (
	"os"
	"testing"
)

func TestC(t *testing.T) {
	t.Log("Test C run")
}

func TestA(t *testing.T) {
	t.Log("Test A run")
}

func TestB(t *testing.T) {
	t.Log("Test B run")
}

func TestMain(m *testing.M) {
	// setup()
	exitVal := m.Run()
	if exitVal == 0 {
		// teardown()
	}
	os.Exit(exitVal)
}
