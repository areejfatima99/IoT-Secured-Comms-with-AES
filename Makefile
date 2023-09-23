.PHONY: build

build:
	@echo "Building..."
	@solc --optimize --abi ./contracts/protocal.sol -o build  --overwrite
	@echo "Generating bin file..."
	@solc --optimize --bin ./contracts/protocal.sol -o build  --overwrite
	@echo "Generating go file..."
	@abigen --abi=./build/IoTMessageProtocol.abi --bin=./build/IoTMessageProtocol.bin --pkg=pkg --out=./pkg/iot-message-protocal.go
	@echo "Compiling..."
	@mkdir -p bin
	@go build -o ./bin/iot-message-protocal main.go



run:
	@echo "Running..."
	@chmod +x ./bin/iot-message-protocal
	@./bin/iot-message-protocal