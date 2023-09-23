package http

import (
	"custom-blockchain-iot-protocol/internal/core/dto"
	ports2 "custom-blockchain-iot-protocol/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type mainHandlerImpl struct {
	repo ports2.EtheriumRepository
}

func (m mainHandlerImpl) SendMessage() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get message from request body
		var messageDto dto.SendMessageRequestDto
		var defaultResponseDto dto.BaseResponseDto

		if err := context.ShouldBindJSON(&messageDto); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Send message to blockchain
		transactionDto, err := m.repo.SendMessage(messageDto.Message, messageDto.To)
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		defaultResponseDto.Success = true
		defaultResponseDto.Data = transactionDto

		transactionDto.Hash = nil

		// Return transaction hash
		context.JSON(200, defaultResponseDto)
	}
}

func (m mainHandlerImpl) GetMessages() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get messages from blockchain
		messages, err := m.repo.GetMessages()
		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Return messages
		context.JSON(200, gin.H{"messages": messages})
	}
}

func NewMainHandler(repo ports2.EtheriumRepository) ports2.MainHandler {
	return &mainHandlerImpl{repo}
}
