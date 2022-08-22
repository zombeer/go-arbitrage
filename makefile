build:
	solc --bin --abi contracts/Arb.sol -o build

abigen:
	abigen --bin=./build/Arb.bin --abi=./build/Arb.abi --pkg=arb --out=./gen/arb.go 

# abigen --bin=./build/IERC20.bin --abi=./build/IERC20.abi --pkg=arb --out=./gen/ierc20.go