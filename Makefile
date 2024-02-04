install_dependencies:
	rm -vrf lib
	mkdir lib
	git clone https://github.com/OpenZeppelin/openzeppelin-contracts.git lib/@openzeppelin --branch v5.0.1 --single-branch

gen:
	rm -vrf artifacts contracts
	mkdir -p artifacts contracts
	sh gen.sh
	# find src -name "*.sol" -exec sh -c 'mkdir -p artifacts/$$(dirname {} | sed "s/src\///"); solc --pretty-json --overwrite --base-path ./src --include-path ./lib --bin {} --abi {} -o artifacts/$$(dirname {} | sed "s/src\///"); mkdir -p contracts/$$(dirname {} | sed "s/src\///")/$$(basename {} .sol); abigen --bin artifacts/$$(dirname {} | sed "s/src\///")/$$(basename {} .sol).bin --abi artifacts/$$(dirname {} | sed "s/src\///")/$$(basename {} .sol).abi --pkg contracts --out contracts/$$(dirname {} | sed "s/src\///")/$$(basename {} .sol)/$$(basename {} .sol).go' \;
