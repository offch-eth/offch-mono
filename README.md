# Offch - off-l1 governance approach

Taiko grant winner project:
1. [Application](https://github.com/taikoxyz/grants/issues/28)
2. [Taiko blog](https://taiko.mirror.xyz/D4dthrBfRz_IuxbrNQKg5cCTUWp6xK53axz4eitUmEo)


## Repository structure

- `src` - Solidity smart contracts for Offch showcase
- `scripts` - golang scripts to interact with the smart contracts (deploy, execute, etc.)
- `artifacts` - generated go bindings for the smart contracts
- `lib` - Solidity external libraries

## How to run

### Prerequisites

1. Create `.env` file from `.env.template` and fill it with the required values


### Execution

Execute test files from the `scripts` directory in the order:
1. `go test ./scripts/l1_governor/deploy_test.go` - deploy L1Governor contract
   1. Change `L1GovernorAddress` in `config.go` to the address of the deployed L1Governor contract
2. `go test ./scripts/l2_voting/deploy_test.go` - deploy L2Factory contract
   1. Change `L2FactoryAddress` in `config.go` to the address of the deployed L2Factory contract
3. `go test ./scripts/l1_token/deploy_test.go` - deploy L1Token contract
   1. Change `L1TokenAddress` in `config.go` to the address of the deployed L1Token contract
4. Run two tests in parallel
   1. `go test ./scripts/l2_voting/create_vote_from_signal_test.go`
   2. `go test ./scripts/l1_governor/create_proposal_test.go`
   3. And change `snapshotBlockNumber` and `proposalKey` in `vote_test.go` to the values from the previous tests
5. `go test ./scripts/l2_voting/vote_test.go` - perform vote at L2
6. Wait for the vote delay time and execute `go test ./scripts/l2_voting/send_results_to_l1_test.go`
   1. Change `l2BlockNumberSignalSent` in `result_from_signal_test.go` to the value from the previous test
7. `go test ./scripts/l1_governor/execute_proposal_test.go` - execute proposal at L1

## Development

1. `make install_dependencies` - download and install all libs required for the project
2. Compile abigen from this [branch](https://github.com/dkeysil/go-ethereum/tree/add-string-methods-for-abigen-structures)
   1. It adds `String` method to go bindings for the smart contracts
3. `make gen` - generate go bindings for the smart contracts


## TODO:

- [ ] One integration test for the whole process
- [ ] Remove hardcoded values from the contracts
- [ ] Get snapshot block number from the last verified block instead of user input
