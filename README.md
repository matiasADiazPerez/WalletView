# WalletView Backend

This repository contains the implementation of the Wallet View Challenge. The API offers the following features:

- **API KEY Authorization:** Secure endpoints with API-KEY based authentication.
- **Wallet Balance:** Display of the balance in USD of all the ethereum ERC20 tokens in an address.
- **Currency Convertion:** If a valid ERC20 token is given (including ETH) the balances are displayed in that token.

The API is built using Golang, employing go-chi as the router. The API is containerized using Docker.

## Usage
### Deploy
1. Create a `.env` file with the following configurations:
```
API_KEY=${Some secret api key}

```
2. To deploy the containers and serve the API at `localhost:8080`, use the command `make app-deploy`
3. To run unit tests, execute: `make test`
### Request
The endpoint to get the balances is:
```
localhost:8080/walletBalance?address={address}&currency={token}
```
Where the address is obligatory and the currency is optional.
a `x-api-key` header with the api key as value must be provided.

## Notes
* A client interface is provided to the Wallet module. AnkrClient satisfice this interface, the idea is that this project can support others web3 token API providers, and is easier to mock for unit test.
* For simplicity, an in memory map is used as Symbol -> Contarct address database to get the token price of the currency. This map is initialized when the server starts and is never refreshed.
* For simplicity, only the wallet service has unit test. The client interface was mock with [Mockery](https://github.com/vektra/mockery)

## Module Diagrams
![wallet](https://github.com/matiasADiazPerez/WalletView/assets/130945302/fac7383f-5f88-4550-a1bb-5ec114562bd7)
