// Copyright 2026 Jimmy Ma
// SPDX-License-Identifier: MIT

package cli

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gosusnp/whackamole/internal/db"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMCPHandlers(t *testing.T) {
	// Setup temp DB
	tmpFile, err := os.CreateTemp("", "whack_mcp_test_*.db")
	require.NoError(t, err)
	testDbPath := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(testDbPath)

	database, err := db.Open(testDbPath)
	require.NoError(t, err)
	defer database.Close()

	history := db.NewHistoryStore(database)
	taskStore := db.NewTaskStore(database, history)
	projectStore := db.NewProjectStore(database, history)

	// Create a project first
	p, err := projectStore.Create("Test Project", "p1")
	require.NoError(t, err)

	t.Run("list_tasks_empty", func(t *testing.T) {
		handler := listTasksHandler(taskStore, projectStore)
		req := mcp.CallToolRequest{}
		req.Params.Arguments = map[string]any{
			"projectKey": "p1",
		}

		result, err := handler(context.Background(), req)
		assert.NoError(t, err)
		assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "No tasks found")
	})

	t.Run("add_task", func(t *testing.T) {
		handler := addTaskHandler(taskStore, projectStore)
		req := mcp.CallToolRequest{}
		req.Params.Arguments = map[string]any{
			"projectKey": "p1",
			"name":       "New Task",
		}

		result, err := handler(context.Background(), req)
		assert.NoError(t, err)
		assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "Task created")

		// Verify task was actually created
		tasks, err := taskStore.ListByProject(p.ID, true)
		assert.NoError(t, err)
		assert.Len(t, tasks, 1)
		assert.Equal(t, "New Task", tasks[0].Name)
	})

	t.Run("list_tasks_with_content", func(t *testing.T) {
		handler := listTasksHandler(taskStore, projectStore)
		req := mcp.CallToolRequest{}
		req.Params.Arguments = map[string]any{
			"projectKey": "p1",
		}

		result, err := handler(context.Background(), req)
		assert.NoError(t, err)
		assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "New Task")
	})

	t.Run("remove_task", func(t *testing.T) {
		tasks, _ := taskStore.ListByProject(p.ID, true)
		taskID := tasks[0].ID

		handler := removeTaskHandler(taskStore)
		req := mcp.CallToolRequest{}
		req.Params.Arguments = map[string]any{
			"taskId": float64(taskID), // mcp-go uses float64 for numbers from JSON
		}

		result, err := handler(context.Background(), req)
		assert.NoError(t, err)
		assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "removed")

		// Verify task was actually removed
		tasks, err = taskStore.ListByProject(p.ID, true)
		assert.NoError(t, err)
		assert.Len(t, tasks, 0)
	})
}

func TestInstallClaudeCode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("fake claude shim is a shell script; not supported on windows")
	}

	t.Run("registers whack via the claude CLI", func(t *testing.T) {
		binDir := t.TempDir()
		callLog := filepath.Join(binDir, "calls.log")

		fakeClaude := filepath.Join(binDir, "claude")
		script := "#!/bin/sh\necho \"$@\" > " + callLog + "\necho added\n"
		require.NoError(t, os.WriteFile(fakeClaude, []byte(script), 0o755))

		t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))

		b := bytes.NewBufferString("")
		cmd := &cobra.Command{}
		cmd.SetOut(b)
		cmd.SetErr(b)

		err := installClaudeCode(cmd)
		require.NoError(t, err)
		assert.Contains(t, b.String(), "added")

		logged, err := os.ReadFile(callLog)
		require.NoError(t, err)
		assert.Contains(t, string(logged), "mcp add whackamole --")
		assert.Contains(t, string(logged), "mcp serve")
		assert.NotContains(t, string(logged), "--database")
	})

	t.Run("forwards --database when explicitly set", func(t *testing.T) {
		binDir := t.TempDir()
		callLog := filepath.Join(binDir, "calls.log")

		fakeClaude := filepath.Join(binDir, "claude")
		script := "#!/bin/sh\necho \"$@\" > " + callLog + "\necho added\n"
		require.NoError(t, os.WriteFile(fakeClaude, []byte(script), 0o755))

		t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))

		cmd := &cobra.Command{}
		cmd.Flags().String("database", "", "")
		require.NoError(t, cmd.Flags().Set("database", "/tmp/custom.db"))
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetErr(b)

		err := installClaudeCode(cmd)
		require.NoError(t, err)

		logged, err := os.ReadFile(callLog)
		require.NoError(t, err)
		assert.Contains(t, string(logged), "mcp serve --database /tmp/custom.db")
	})

	t.Run("errors when claude CLI is missing", func(t *testing.T) {
		emptyDir := t.TempDir()
		t.Setenv("PATH", emptyDir)

		cmd := &cobra.Command{}
		err := installClaudeCode(cmd)
		assert.ErrorContains(t, err, "claude CLI not found")
	})
}
