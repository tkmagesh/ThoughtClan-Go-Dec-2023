//go:generate mockery --name MessageService
package services

type MessageService interface {
	Send(msg string) bool
}
