package cli

const (
	help                = "help"
	addOrder            = "add"
	returnToDeliverer   = "return"
	giveToAddressee     = "give"
	listOrder           = "list"
	returnFromAddressee = "refund"
	listRefund          = "refund_list"
	exit                = "exit"
)

type command struct {
	name        string
	description string
}
