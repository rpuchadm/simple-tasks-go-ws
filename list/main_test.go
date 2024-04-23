package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestList(t *testing.T) {
	tasks := []*Task{
		{ID: "1", Name: "Task 1", Status: "Pending"},
		{ID: "2", Name: "Task 2", Status: "Completed"},
		{ID: "3", Name: "Task 3", Status: "Pending"},
	}

	// Redirect stdout to capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	List(tasks)

	// Reset stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Assert the expected output
	expected := "ID: 1, Name: Task 1, Status: Pending\n" +
		"ID: 2, Name: Task 2, Status: Completed\n" +
		"ID: 3, Name: Task 3, Status: Pending\n"
	if buf.String() != expected {
		t.Errorf("Unexpected output:\nExpected: %s\nActual: %s", expected, buf.String())
	}
}
