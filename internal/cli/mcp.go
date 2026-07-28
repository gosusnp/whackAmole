// Copyright 2026 Jimmy Ma
// SPDX-License-Identifier: MIT

package cli

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/gosusnp/whackamole/internal"
	"github.com/gosusnp/whackamole/internal/db"
	"github.com/gosusnp/whackamole/internal/types"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start the MCP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startMCPServer(cmd)
	},
}

var mcpServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the MCP server (alias for 'whack mcp')",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startMCPServer(cmd)
	},
}

var mcpInstallClaudeCodeCmd = &cobra.Command{
	Use:   "install-claude-code",
	Short: "Register whack's MCP server with Claude Code",
	RunE: func(cmd *cobra.Command, args []string) error {
		return installClaudeCode(cmd)
	},
}

func init() {
	mcpCmd.AddCommand(mcpServeCmd)
	mcpCmd.AddCommand(mcpInstallClaudeCodeCmd)
	rootCmd.AddCommand(mcpCmd)
}

func installClaudeCode(cmd *cobra.Command) error {
	whackPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to resolve whack executable path: %w", err)
	}

	claudePath, err := exec.LookPath("claude")
	if err != nil {
		return fmt.Errorf("claude CLI not found in PATH: %w", err)
	}

	serverArgs := []string{"mcp", "serve"}
	if cmd.Flags().Changed("database") {
		serverArgs = append(serverArgs, "--database", getDBPath(cmd))
	}
	addArgs := append([]string{"mcp", "add", "whackamole", "--", whackPath}, serverArgs...)
	install := exec.Command(claudePath, addArgs...)
	install.Stdout = cmd.OutOrStdout()
	install.Stderr = cmd.ErrOrStderr()
	if err := install.Run(); err != nil {
		return fmt.Errorf("failed to register whack with Claude Code: %w", err)
	}

	return nil
}

func startMCPServer(cmd *cobra.Command) error {
	s, err := createMCPServer(getDBPath(cmd))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Starting whackAmole MCP server version %s\n", internal.Version)
	return server.ServeStdio(s)
}

func createMCPServer(path string) (*server.MCPServer, error) {
	database, err := db.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	configStore := db.NewConfigStore(database)
	config, _ := configStore.GetConfig(types.ConfigKeyMCPInstructions)

	var options []server.ServerOption
	options = append(options, server.WithLogging())
	if config != nil && config.Value != "" {
		options = append(options, server.WithInstructions(config.Value))
	}

	s := server.NewMCPServer(
		"whackAmole",
		internal.Version,
		options...,
	)

	history := db.NewHistoryStore(database)
	taskStore := db.NewTaskStore(database, history)
	projectStore := db.NewProjectStore(database, history)

	// List Tasks
	s.AddTool(mcp.NewTool("list_tasks",
		mcp.WithDescription("List tasks for a project"),
		mcp.WithString("projectKey", mcp.Required(), mcp.Description("The Key of the project")),
	), listTasksHandler(taskStore, projectStore))

	// Show Project
	s.AddTool(mcp.NewTool("show_project",
		mcp.WithDescription("Show detailed information for a project"),
		mcp.WithString("projectKey", mcp.Required(), mcp.Description("The Key of the project")),
	), showProjectHandler(projectStore))

	// Show Task
	s.AddTool(mcp.NewTool("show_task",
		mcp.WithDescription("Show detailed information for a task"),
		mcp.WithNumber("taskId", mcp.Required(), mcp.Description("The ID of the task")),
	), showTaskHandler(taskStore))

	// Add Task
	s.AddTool(mcp.NewTool("add_task",
		mcp.WithDescription("Add a new task to a project"),
		mcp.WithString("projectKey", mcp.Required(), mcp.Description("The Key of the project")),
		mcp.WithString("name", mcp.Required(), mcp.Description("The name of the task")),
		mcp.WithString("description", mcp.Description("The description of the task (supports Markdown)")),
		mcp.WithString("type", mcp.Description("The type of the task (feat, bug, docs, refactor, chore)")),
		mcp.WithString("status", mcp.Description("The status of the task (notStarted, inProgress, review, blocked, completed, closed)")),
	), addTaskHandler(taskStore, projectStore))

	// Update Task
	s.AddTool(mcp.NewTool("update_task",
		mcp.WithDescription("Update an existing task"),
		mcp.WithNumber("taskId", mcp.Required(), mcp.Description("The ID of the task")),
		mcp.WithString("name", mcp.Description("The new name of the task")),
		mcp.WithString("description", mcp.Description("The new description of the task (supports Markdown)")),
		mcp.WithString("type", mcp.Description("The new type of the task")),
		mcp.WithString("status", mcp.Description("The new status of the task")),
	), updateTaskHandler(taskStore))

	// Remove Task
	s.AddTool(mcp.NewTool("remove_task",
		mcp.WithDescription("Remove a task"),
		mcp.WithNumber("taskId", mcp.Required(), mcp.Description("The ID of the task to remove")),
	), removeTaskHandler(taskStore))

	return s, nil
}

func listTasksHandler(taskStore *db.TaskStore, projectStore *db.ProjectStore) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.Params.Arguments.(map[string]any)
		pKey := types.ProjectKey(args["projectKey"].(string))
		p, err := projectStore.GetByKey(pKey)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		tasks, err := taskStore.ListByProject(p.ID, true)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		if len(tasks) == 0 {
			return mcp.NewToolResultText("No tasks found for this project."), nil
		}

		var result string
		for _, t := range tasks {
			result += fmt.Sprintf("ID: %d, Name: %s, Status: %s, Type: %s\n", t.ID, t.Name, t.Status, t.Type)
		}

		return mcp.NewToolResultText(result), nil
	}
}

func showProjectHandler(projectStore *db.ProjectStore) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.Params.Arguments.(map[string]any)
		pKey := types.ProjectKey(args["projectKey"].(string))
		p, err := projectStore.GetByKey(pKey)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		result := fmt.Sprintf("ID: %d\nKey: %s\nName: %s\nCreated At: %s\nUpdated At: %s\n",
			p.ID, p.Key, p.Name, p.CreatedAt, p.UpdatedAt)

		return mcp.NewToolResultText(result), nil
	}
}

func showTaskHandler(taskStore *db.TaskStore) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.Params.Arguments.(map[string]any)
		tID := int64(args["taskId"].(float64))
		t, err := taskStore.Get(types.TaskID(tID))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		result := fmt.Sprintf("ID: %d\nProject ID: %d\nName: %s\nDescription: %s\nType: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n",
			t.ID, t.ProjectID, t.Name, t.Description, t.Type, t.Status, t.CreatedAt, t.UpdatedAt)

		return mcp.NewToolResultText(result), nil
	}
}

func addTaskHandler(taskStore *db.TaskStore, projectStore *db.ProjectStore) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.Params.Arguments.(map[string]any)
		pKey := types.ProjectKey(args["projectKey"].(string))
		p, err := projectStore.GetByKey(pKey)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		name := args["name"].(string)
		desc, _ := args["description"].(string)
		tTypeStr, _ := args["type"].(string)
		statusStr, _ := args["status"].(string)

		t, err := taskStore.Create(p.ID, name, desc, types.TaskType(tTypeStr), types.TaskStatus(statusStr))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Task created: %d - %s", t.ID, t.Name)), nil
	}
}

func updateTaskHandler(taskStore *db.TaskStore) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.Params.Arguments.(map[string]any)
		tID := int64(args["taskId"].(float64))
		current, err := taskStore.Get(types.TaskID(tID))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		name := current.Name
		if val, ok := args["name"].(string); ok {
			name = val
		}
		desc := current.Description
		if val, ok := args["description"].(string); ok {
			desc = val
		}
		tType := current.Type
		if val, ok := args["type"].(string); ok {
			tType = types.TaskType(val)
		}
		status := current.Status
		if val, ok := args["status"].(string); ok {
			status = types.TaskStatus(val)
		}

		_, err = taskStore.Update(types.TaskID(tID), name, desc, tType, status)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Task %d updated.", tID)), nil
	}
}

func removeTaskHandler(taskStore *db.TaskStore) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.Params.Arguments.(map[string]any)
		tID := int64(args["taskId"].(float64))
		err := taskStore.Delete(types.TaskID(tID))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Task %d removed.", tID)), nil
	}
}
