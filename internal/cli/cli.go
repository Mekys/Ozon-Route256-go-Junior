package cli

import (
	"errors"
	"flag"
	"fmt"
	"homework-3/internal/app/answer"
	"homework-3/internal/kafka"
	"homework-3/internal/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Module interface {
	AddOrder(models.Order) error
	ReturnToDeliverer(models.OrderId) error
	ReturnOrder(models.Order) error
	DispatchOrders([]models.OrderId) error
	ListRefund(int64, int64) error
	ListOrders(int64, int64) error
}

type Deps struct {
	Module Module
}

type CLI struct {
	Deps
	commandList []command
}

var (
	brokers = []string{"127.0.0.1:9091"}
)

// NewCLI creates a command line interface
func NewCLI(d Deps) CLI {
	return CLI{
		Deps: d,
		commandList: []command{
			{
				name:        help,
				description: "Справка",
			},
			{
				name:        addOrder,
				description: "Добавить заказ: использование add --orderId=142 --addresseeId=4323 --shelfLife=2024/06/31",
			},
			{
				name:        returnToDeliverer,
				description: "Вернуть заказ курьеру return --orderId=123",
			},
			{
				name:        giveToAddressee,
				description: "Выдать заказ клиенту: использование give 1 5 44 23 33",
			},
			{
				name:        listOrder,
				description: "Получить список заказов: использование list --clientid=1 --n=10(optional)",
			},
			{
				name:        returnFromAddressee,
				description: "Принять возврат от клиента: использование refund",
			},
			{
				name:        listRefund,
				description: "Получить список возвратов: использование refund_list --pagelen=10 (optional, default 10) --pagenumber=2(optional, default 1)",
			},
			{
				name:        exit,
				description: "Выйти из программы",
			},
		},
	}
}

// Run ..
func (c CLI) Run(text string) error {
	args := strings.Split(text, " ")
	if len(args) == 0 {
		fmt.Println("command isn't set")
		return fmt.Errorf("command isn't set help")
	}
	operationId := uuid.New().String()

	commandName := args[0]
	switch commandName {
	case help:
		Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: true})
		defer Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: false})
		c.help()
		return nil
	case addOrder:
		Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: true})
		defer Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: false})
		return c.addOrder(args[1:])
	case returnToDeliverer:
		Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: true})
		defer Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: false})
		return c.ReturnToDeliverer(args[1:])
	case returnFromAddressee:
		Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: true})
		defer Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: false})
		return c.ReturnFromAddressee(args[1:])
	case listRefund:
		Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: true})
		defer Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: false})
		return c.PrintRefund(args[1:])
	case listOrder:
		Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: true})
		defer Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: false})
		return c.PrintOrder(args[1:])
	case giveToAddressee:
		Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: true})
		defer Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: false})
		return c.DispatchOrders(args[1:])
	case exit:
		Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: true})
		defer Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: commandName, IsStart: false})
		os.Exit(1)
	}
	Produce(brokers, &answer.CommandMessage{Id: operationId, CreatedAt: time.Now(), Command: text, CommandName: "Undefined", IsStart: true})

	return fmt.Errorf("command isn't set")
}

func (c CLI) addOrder(args []string) error {
	var orderId, addresseeId, weight, price, wrapType int64
	var shelfLife string

	fs := flag.NewFlagSet(addOrder, flag.ContinueOnError)
	fs.Int64Var(&orderId, "orderId", -1, "use --orderId=123")
	fs.Int64Var(&addresseeId, "addresseeId", -1, "use --addresseeId=123")
	fs.StringVar(&shelfLife, "shelfLife", "", "use --shelfLife=2024/12/31")
	fs.Int64Var(&weight, "weight", 0, "use --weight=15")
	fs.Int64Var(&price, "price", 0, "use --price=20")
	fs.Int64Var(&wrapType, "wrap", 0, "use --wrap=1")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if orderId == -1 {
		return errors.New("orderId is empty")
	}
	if orderId < 0 {
		return errors.New("orderId must be greater or equal  zero")
	}
	if addresseeId == -1 {
		return errors.New("addresseeId is empty")
	}
	if addresseeId < 0 {
		return errors.New("addresseeId must be greater or equal  zero")
	}
	t, parseErr := time.Parse("2006/01/02 15:04:05", shelfLife+" 23:59:59")
	if parseErr != nil {
		return parseErr
	}
	if !t.After(time.Now()) {
		return errors.New("shelfLife cannot be early current date")
	}
	newOrder := models.Order{
		OrderId:     models.OrderId(orderId),
		AddresseeId: models.AddresseeId(addresseeId),
		ShelfLife:   t,
		Price:       int(price),
		Weight:      int(weight),
	}

	wrapper, err := models.GetWrapper(int(wrapType))

	if err != nil {
		return err
	}
	newOrder.Wrapper = wrapper

	return c.Module.AddOrder(newOrder)
}
func (c CLI) ReturnToDeliverer(args []string) error {
	var orderId int64

	fs := flag.NewFlagSet(returnToDeliverer, flag.ContinueOnError)
	fs.Int64Var(&orderId, "orderId", -1, "use --orderId=123")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if orderId == -1 {
		return errors.New("orderId is empty")
	}

	return c.Module.ReturnToDeliverer(models.OrderId(orderId))
}
func (c CLI) ReturnFromAddressee(args []string) error {
	var orderId, addresseeId int64

	fs := flag.NewFlagSet(returnToDeliverer, flag.ContinueOnError)
	fs.Int64Var(&orderId, "orderId", -1, "use --orderId=123")
	fs.Int64Var(&addresseeId, "addresseeId", -1, "use --addresseeId=123")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if orderId == -1 {
		return errors.New("orderId is empty")
	}
	if addresseeId == -1 {
		return errors.New("addresseeId is empty")
	}

	return c.Module.ReturnOrder(models.Order{
		OrderId:     models.OrderId(orderId),
		AddresseeId: models.AddresseeId(addresseeId)})
}
func (c CLI) PrintRefund(args []string) error {
	var pageLen, pageNumber int64

	fs := flag.NewFlagSet(returnFromAddressee, flag.ContinueOnError)
	fs.Int64Var(&pageLen, "pagelen", 10, "use --pagelen=10")
	fs.Int64Var(&pageNumber, "pagenumber", 1, "use --pagenumber=2")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if pageLen <= 0 {
		return errors.New("pageLen must be greater than zero")
	}
	if pageNumber <= 0 {
		return errors.New("pageNumber must be greater than zero")
	}

	return c.Module.ListRefund(pageLen, pageNumber)
}
func (c CLI) PrintOrder(args []string) error {
	var clientId, n int64

	fs := flag.NewFlagSet(returnFromAddressee, flag.ContinueOnError)
	fs.Int64Var(&clientId, "clientId", -1, "use --clientId=10")
	fs.Int64Var(&n, "n", 0, "use --n=2")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if clientId == -1 {
		return errors.New("clientId is Empty")
	}
	if clientId < 0 {
		return errors.New("clientId must be greater than zero")
	}
	if n < 0 {
		return errors.New("n must be greater or equal zero")
	}

	return c.Module.ListOrders(clientId, n)
}
func (c CLI) DispatchOrders(args []string) error {
	orderIds := make([]models.OrderId, 0)

	for _, idStr := range args {
		idInt, err := strconv.Atoi(idStr)

		if err != nil {
			return fmt.Errorf("Cannot parse id: %s", idStr)
		}

		if idInt < 0 {
			return errors.New("OrderId must be positive number")
		}

		orderIds = append(orderIds, models.OrderId(idInt))
	}

	return c.Module.DispatchOrders(orderIds)
}

func (c CLI) help() {
	fmt.Println("command list:")
	for _, cmd := range c.commandList {
		fmt.Println("", cmd.name, cmd.description)
	}
	return
}

func Produce(brokers []string, payment *answer.CommandMessage) error {
	kafkaProducer, err := kafka.NewProducer(brokers)
	if err != nil {
		fmt.Println(err)
		return err
	}

	producer := answer.NewKafkaSender(kafkaProducer, "logs")
	producer.SendAsyncMessage(payment)

	return nil
}
