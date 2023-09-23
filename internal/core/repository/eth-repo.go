package repository

// Importing necessary packages.
import (
	"custom-blockchain-iot-protocol/internal/core/dto"   // Custom blockchain IoT protocol DTO package.
	"custom-blockchain-iot-protocol/internal/core/ports" // Custom blockchain IoT protocol ports package.
	"custom-blockchain-iot-protocol/pkg"                 // Custom blockchain IoT protocol package.
	"github.com/ethereum/go-ethereum/accounts/abi/bind"  // Ethereum contract bindings package.
	"github.com/ethereum/go-ethereum/common"             // Ethereum common utilities package.
	"github.com/ethereum/go-ethereum/ethclient"          // Ethereum client package.
	"log"                                                // Logging package for error handling.
	"math/big"                                           // Package for arbitrary-precision arithmetic.
)

// Struct definition for individual messages.
type msg struct {
	from string // Sender's address.
	msg  string // Message content.
}

// Definition of an inbox map for storing messages.
type inbox map[string][]msg

// Variable to store incoming messages.
var messages inbox

// Struct definition for implementing an Ethereum repository.
type ethRepoImpl struct {
	client     *ethclient.Client // Ethereum client instance.
	conn       *pkg.Pkg          // Custom blockchain IoT protocol package instance.
	myAddr     common.Address    // Ethereum address of the repository owner.
	encryption *pkg.Encryption   // Encryption utility for message handling.
}

// SendMessage method to send an encrypted message.
func (e ethRepoImpl) SendMessage(message string, address string) (*dto.TransactionDto, error) {

	// Encrypt the message.
	encryptedMessage, err := e.encryption.Encrypt([]byte(message))
	if err != nil {
		log.Println("Error encrypting message") // Log an error if encryption fails.
		return nil, err
	}

	// Get account authentication.
	auth, err := pkg.GetAccountAuth(e.client, address)
	if err == nil {
		// Send the encrypted message.
		reply, err := e.conn.SendMessage(auth, common.HexToAddress(address), encryptedMessage)
		if err != nil {
			log.Println("Error sending message") // Log an error if sending fails.
			return nil, err
		}

		// Return a DTO for the transaction.
		return dto.NewTransactionDto(reply), nil
	}

	// Store the message locally if authentication fails.
	messages[address] = append(messages[address], msg{e.myAddr.Hex(), string(encryptedMessage)})
	msgDto := &dto.TransactionDto{
		GasFee: 0,
		Value:  &big.Int{},
	}

	return msgDto, nil
}

// GetMessages method to retrieve messages.
func (e ethRepoImpl) GetMessages() ([]pkg.IoTMessageProtocolMessage, error) {
	reply, err := e.conn.GetMessages(&bind.CallOpts{}) // Retrieve messages from the blockchain.
	if reply != nil && err == nil {
		return reply, nil
	}

	// Process messages stored locally.
	msgs := make([]pkg.IoTMessageProtocolMessage, 0)
	for _, v := range messages[e.myAddr.Hex()] {
		fromAddr := common.HexToAddress(v.from)
		toAddr := e.myAddr

		// Decrypt the message.
		decryptedMessage, err := e.encryption.Decrypt([]byte(v.msg))
		if err != nil {
			return nil, err
		}

		// Add the message to the list of retrieved messages.
		msgs = append(msgs, pkg.IoTMessageProtocolMessage{
			Sender:           fromAddr,
			Recipient:        toAddr,
			EncryptedMessage: decryptedMessage,
		})
	}

	return msgs, nil
}

// NewEthRepo function to create a new Ethereum repository instance.
func NewEthRepo(client *ethclient.Client, address common.Address, encryption *pkg.Encryption) (ports.EtheriumRepository, error) {
	// Create a new package instance.
	conn, err := pkg.NewPkg(common.HexToAddress(address.Hex()), client)
	if err != nil {
		return nil, err
	}

	// Initialize the inbox for storing messages.
	messages = make(inbox)

	// Generate a random encryption key.
	encryption.Key = pkg.GenerateRandomKey()

	// Return the repository instance.
	return &ethRepoImpl{client: client, conn: conn, myAddr: address, encryption: encryption}, nil
}
