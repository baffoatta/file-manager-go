package app

import (
    "github.com/spf13/cobra"
    
    "github.com/baffoatta/filemanager/internal/config"
    "github.com/baffoatta/filemanager/internal/service/fileservice"
)


// Logger interface for logging messages	
type Logger interface {
    Info(msg string, keysAndValues ...interface{})
    Error(msg string, keysAndValues ...interface{})
}

// App struct holds the configuration, logger, and file service
type App struct {
    cfg    *config.Config // Configuration for the application
    logger Logger // Logger for logging messages
    fs     *fileservice.FileService // File service for file operations
}



// New function initializes a new App instance with the provided configuration and logger
func New(cfg *config.Config, logger Logger) *App {
    return &App{
        cfg:    cfg, // Set the configuration
        logger: logger, // Set the logger
        fs:     fileservice.New(cfg.BaseDir, logger), // Initialize the file service with base directory and logger
    }
}

// Run method starts the command-line application
func (a *App) Run() error {
    // Create the root command for the CLI application
    rootCmd := &cobra.Command{
        Use:   "filemanager", // Command name
        Short: "A modern file manager CLI application", // Short description
    }

    // Add subcommands to the root command
    rootCmd.AddCommand(
        a.listCommand(), // Command to list files
        a.createCommand(), // Command to create a file
        a.deleteCommand(), // Command to delete a file
        a.copyCommand(), // Command to copy a file
        a.moveCommand(), // Command to move a file
    )

    // Execute the root command
    return rootCmd.Execute()
}

// listCommand method returns a command to list files in a directory
func (a *App) listCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "list [path]", // Command usage
        Short: "List files in a directory", // Short description
        Args:  cobra.MaximumNArgs(1), // Accepts 0 or 1 argument
        RunE: func(cmd *cobra.Command, args []string) error {
            path := "." // Default path is current directory
            if len(args) > 0 {
                path = args[0] // Use provided path if available
            }

            // List files in the specified path
            files, err := a.fs.List(path)
            if err != nil {
                return err // Return error if listing fails
            }

            // Print each file's name, size, and modification time
            for _, file := range files {
                cmd.Printf("%s\t%d\t%s\n", file.Name, file.Size, file.ModifiedAt)
            }
            return nil
        },
    }
}

// createCommand method returns a command to create a new file
func (a *App) createCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "create [path]", // Command usage
        Short: "Create a new file", // Short description
        Args:  cobra.ExactArgs(1), // Requires exactly 1 argument
        RunE: func(cmd *cobra.Command, args []string) error {
            path := args[0] // Get the file path from arguments
            
            // Create the file using the file service
            if err := a.fs.Create(path); err != nil {
                return err // Return error if creation fails
            }
            
            cmd.Printf("Successfully created file: %s\n", path) // Print success message
            return nil
        },
    }
}

// deleteCommand method returns a command to delete a file
func (a *App) deleteCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "delete [path]", // Command usage
        Short: "Delete a file", // Short description
        Args:  cobra.ExactArgs(1), // Requires exactly 1 argument
        RunE: func(cmd *cobra.Command, args []string) error {
            path := args[0] // Get the file path from arguments
            
            // Delete the file using the file service
            if err := a.fs.Delete(path); err != nil {
                return err // Return error if deletion fails
            }
            
            cmd.Printf("Successfully deleted file: %s\n", path) // Print success message
            return nil
        },
    }
}

// copyCommand method returns a command to copy a file
func (a *App) copyCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "copy [source] [destination]", // Command usage
        Short: "Copy a file from source to destination", // Short description
        Args:  cobra.ExactArgs(2), // Requires exactly 2 arguments
        RunE: func(cmd *cobra.Command, args []string) error {
            src := args[0] // Get the source file path
            dst := args[1] // Get the destination file path
            
            // Copy the file using the file service
            if err := a.fs.Copy(src, dst); err != nil {
                return err // Return error if copying fails
            }
            
            cmd.Printf("Successfully copied file from %s to %s\n", src, dst) // Print success message
            return nil
        },
    }
}

// moveCommand method returns a command to move a file
func (a *App) moveCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "move [source] [destination]", // Command usage
        Short: "Move a file from source to destination", // Short description
        Args:  cobra.ExactArgs(2), // Requires exactly 2 arguments
        RunE: func(cmd *cobra.Command, args []string) error {
            src := args[0] // Get the source file path
            dst := args[1] // Get the destination file path
            
            // Move the file using the file service
            if err := a.fs.Move(src, dst); err != nil {
                return err // Return error if moving fails
            }
            
            cmd.Printf("Successfully moved file from %s to %s\n", src, dst) // Print success message
            return nil
        },
    }
}