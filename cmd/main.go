package main

import (
	"bufio"
	"context"
	"fmt"
	"homework-3/internal/cli"
	"homework-3/internal/module"
	"homework-3/internal/storage"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbURL = "postgres://postgres:password@localhost:5433/postgres?sslmode=disable"
)

var (
	commands cli.CLI
)

func main() {
	pool, err := pgxpool.New(context.Background(), dbURL)

	if err != nil {
		os.Exit(0)
	}
	storageJSON := storage.NewStorage(pool)
	phoneBookService := module.NewModule(module.Deps{
		Storage: storageJSON,
	})
	commands = cli.NewCLI(cli.Deps{Module: phoneBookService})
	const (
		numJobs    = 5
		numWorkers = 5
	)
	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	var (
		jobs = make(chan string, numJobs) // буферизированный канал для задач (размер numJobs)
		wg   = sync.WaitGroup{}
	)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = strings.Join(strings.Fields(text), " ")
			if text == "exit" {
				break
			}
			if text != "" {
				jobs <- text
			}
		}
		chSignal <- syscall.SIGTERM
	}()

	<-chSignal
	fmt.Println("gracefull shutdown")
	close(jobs)
	wg.Wait()
	fmt.Println("Можем завершаться")

}

func worker(id int, wg *sync.WaitGroup, jobs <-chan string) {
	defer wg.Done()
	fmt.Printf("Воркер %d начал свою работу\n", id)

	for j := range jobs {
		fmt.Printf("Воркер %d выполняет работу %s\n", id, j)
		if err := commands.Run(j); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Воркер %d закончил выполнять работу %s\n", id, j)
	}
}
