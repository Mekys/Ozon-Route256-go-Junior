package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	service "homework-3/internal/api"
	"homework-3/internal/app/answer"
	"homework-3/internal/app/command"
	"homework-3/internal/cli"
	"homework-3/internal/kafka"
	"homework-3/internal/module"
	"homework-3/internal/storage"
	__order "homework-3/pkg/api/proto/order/v1/order/v1"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

const (
	dbURL = "postgres://postgres:password@localhost:5433/postgres?sslmode=disable"
)

var (
	commands      cli.CLI
	readFromKafka = false
)

func main() {
	const (
		numJobs    = 5
		numWorkers = 5
	)
	var (
		jobs      = make(chan string, numJobs) // буферизированный канал для задач (размер numJobs)
		consumers = make(chan struct{})        // буферизированный канал для задач (размер numJobs)
		wg        = sync.WaitGroup{}
	)
	fs := flag.NewFlagSet("kafka", flag.ContinueOnError)
	fs.BoolVar(&readFromKafka, "kafka", false, "a string")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if readFromKafka {
		go consumer([]string{"127.0.0.1:9091"}, &wg, consumers)
		wg.Add(1)
	}

	if err != nil {
		os.Exit(0)
	}
	storageJSON := storage.NewStorage(pool)
	phoneBookService := module.NewModule(module.Deps{
		Storage: storageJSON,
	})
	orderService := &service.TelephoneService{
		Module: phoneBookService,
	}
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc := grpc.NewServer()

	__order.RegisterOrderServer(grpc, orderService)

	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}

	go func() {
		fmt.Println("Start")
		if err = grpc.Serve(lis); err != nil {
			fmt.Println("Error")
		}
		chSignal <- syscall.SIGTERM
	}()

	<-chSignal
	fmt.Println("gracefull shutdown")
	close(jobs)
	close(consumers)
	wg.Wait()
	fmt.Println("Можем завершаться")

}

func worker(id int, wg *sync.WaitGroup, jobs <-chan string) {
	defer wg.Done()
	fmt.Printf("Воркер %d начал свою работу\n", id)

	for j := range jobs {
		if !readFromKafka {
			fmt.Printf("Воркер %d выполняет работу %s\n", id, j)
		}
		if err := commands.Run(j); err != nil {
			fmt.Println(err)
		}
		if !readFromKafka {
			fmt.Printf("Воркер %d закончил выполнять работу %s\n", id, j)
		}
	}
}

func consumer(brokers []string, wg *sync.WaitGroup, chanel <-chan struct{}) {
	defer wg.Done()
	kafkaConsumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		fmt.Println(err)
	}

	handlers := map[string]command.HandleFunc{
		"logs": func(message *sarama.ConsumerMessage) {
			pm := answer.CommandMessage{}
			err = json.Unmarshal(message.Value, &pm)
			if err != nil {
				fmt.Println("Consumer error", err)
			}
			if pm.IsStart {
				fmt.Printf("[%s] %s - CommandName: %s, Command: %s, Status: started\n", pm.Id, pm.CreatedAt.String(), pm.CommandName, pm.Command)
			} else {
				fmt.Printf("[%s] %s - CommandName: %s, Command: %s, Status: finished\n", pm.Id, pm.CreatedAt.String(), pm.CommandName, pm.Command)
			}
		},
	}
	recv := command.NewReceiver(kafkaConsumer, handlers)
	recv.Subscribe("logs")

	for range chanel {

	}
}
