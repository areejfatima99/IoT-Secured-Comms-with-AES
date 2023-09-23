package ports

import (
	"custom-blockchain-iot-protocol/internal/core/dto"
	"custom-blockchain-iot-protocol/pkg"
)

type EtheriumRepository interface {
	SendMessage(message string, address string) (*dto.TransactionDto, error)
	GetMessages() ([]pkg.IoTMessageProtocolMessage, error)
}
