package app

import (
    "github.com/spf13/cobra"
    
    "github.com/yourusername/filemanager/internal/config"
    "github.com/yourusername/filemanager/internal/service/fileservice"
)

type App struct {
    cfg    *config.Config
    logger Logger
    fs     *fileservice.FileService
}

func New(cfg *config.Config, logger Logger) *App {
    return &App{
        cfg:    cfg,
        logger: logger,
        fs:     fileservice.New(cfg.BaseDir, logger),
    }
}

func (a *App) Run() error {
    rootCmd := &cobra.Command{
        Use:   "filemanager",
        Short: "A modern file manager CLI application",
    }

    rootCmd.AddCommand(
        a.listCommand(),
        a.createCommand(),
        a.deleteCommand(),
        a.copyCommand(),
        a.moveCommand(),
    )

    return rootCmd.Execute()
}

func (a *App) listCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "list [path]",
        Short: "List files in a directory",
        Args:  cobra.MaximumNArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            path := "."
            if len(args) > 0 {
                path = args[0]
            }

            files, err := a.fs.List(path)
            if err != nil {
                return err
            }

            for _, file := range files {
                cmd.Printf("%s\t%d\t%s\n", file.Name, file.Size, file.ModTime)
            }
            return nil
        },
    }
}

func (a *App) createCommand() *cobra.Command {
	return &cobra.Command{
			Use:   "create [path]",
			Short: "Create a new file",
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
					path := args[0]
					
					if err := a.fs.Create(path); err != nil {
							return err
					}
					
					cmd.Printf("Successfully created file: %s\n", path)
					return nil
			},
	}
}

func (a *App) deleteCommand() *cobra.Command {
	return &cobra.Command{
			Use:   "delete [path]",
			Short: "Delete a file",
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
					path := args[0]
					
					if err := a.fs.Delete(path); err != nil {
							return err
					}
					
					cmd.Printf("Successfully deleted file: %s\n", path)
					return nil
			},
	}
}

func (a *App) copyCommand() *cobra.Command {
	return &cobra.Command{
			Use:   "copy [source] [destination]",
			Short: "Copy a file from source to destination",
			Args:  cobra.ExactArgs(2),
			RunE: func(cmd *cobra.Command, args []string) error {
					src := args[0]
					dst := args[1]
					
					if err := a.fs.Copy(src, dst); err != nil {
							return err
					}
					
					cmd.Printf("Successfully copied file from %s to %s\n", src, dst)
					return nil
			},
	}
}

func (a *App) moveCommand() *cobra.Command {
	return &cobra.Command{
			Use:   "move [source] [destination]",
			Short: "Move a file from source to destination",
			Args:  cobra.ExactArgs(2),
			RunE: func(cmd *cobra.Command, args []string) error {
					src := args[0]
					dst := args[1]
					
					if err := a.fs.Move(src, dst); err != nil {
							return err
					}
					
					cmd.Printf("Successfully moved file from %s to %s\n", src, dst)
					return nil
			},
	}
}