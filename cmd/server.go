package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/felipeweb/clean-arch/handlers"
	"github.com/felipeweb/clean-arch/repository"
	"github.com/felipeweb/clean-arch/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	Long: `Start the HTTP server For example:
	ports server --port 8080 --repo memory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := validateFlags(cmd)
		if err != nil {
			return err
		}
		port := cmd.Flag("port").Value.String()
		repo := cmd.Flag("repo").Value.String()
		return runServer(port, repo, cmd.OutOrStdout(), cmd.OutOrStderr())
	},
}

func runServer(port, repo string, out, errOut io.Writer) error {
	fmt.Fprintf(out, "Starting server on port %s with %s repository\n", port, repo)
	ctx := context.Background()
	var r usecase.PortRepository
	switch repo {
	case "memory":
		r = repository.NewInMemory()
	case "postgres":
		//TODO: implement postgres repository
	}
	svc := usecase.NewPortService(r)
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	handlers.MakePortsHandler(router, svc)
	logger := log.New(errOut, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + port,
		Handler:      router,
		ErrorLog:     logger,
	}
	cerr := make(chan error, 1)
	done := make(chan struct{}, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func(ctx context.Context) {
		<-shutdown
		fmt.Fprintln(out, "Shutting down server...")
		if err := srv.Shutdown(ctx); err != nil {
			cerr <- err
		}
		done <- struct{}{}
	}(ctx)
	go func(ctx context.Context) {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			cerr <- err
			return
		}
		done <- struct{}{}
	}(ctx)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-cerr:
		return err
	case <-done:
		return nil
	}
}

func validateFlags(cmd *cobra.Command) error {
	repo := cmd.Flag("repo").Value.String()
	if repo != "memory" && repo != "postgres" {
		return fmt.Errorf("invalid repository type: %s", repo)
	}
	return nil
}

func init() {
	serverCmd.PersistentFlags().IntP("port", "p", 8080, "Port to listen on, default: 8080")
	serverCmd.PersistentFlags().StringP("repo", "r", "memory", "ports repository values: [memory, postgres] default: memory")
	rootCmd.AddCommand(serverCmd)
}
