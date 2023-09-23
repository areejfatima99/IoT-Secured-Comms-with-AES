// Solidity Smart Contract
pragma solidity >0.7.0 < 0.9.0;

contract IoTMessageProtocol {
    struct Message {
        address sender;
        address recipient;
        bytes encryptedMessage;
    }

    Message[] public messages;

    function sendMessage(address _recipient, bytes memory _encryptedMessage) public {
        Message memory newMessage = Message({
            sender: msg.sender,
            recipient: _recipient,
            encryptedMessage: _encryptedMessage
        });
        messages.push(newMessage);
    }

    function getMessages() public view returns (Message[] memory) {
        return messages;
    }
}
