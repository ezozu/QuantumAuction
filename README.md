# QuantumAuction - Decentralized NFT Marketplace with On-Chain Order Book Aggregation

QuantumAuction is a cutting-edge, decentralized NFT marketplace engineered for optimal liquidity and efficient price discovery. Leveraging the power of smart contract composability and on-chain order book aggregation, QuantumAuction provides a superior trading experience compared to traditional NFT marketplaces. Our platform integrates directly with Automated Market Makers (AMMs) to provide dynamically routed liquidity, ensuring that users receive the best possible prices for their NFTs.

The core aim of QuantumAuction is to address the fragmentation and illiquidity inherent in the current NFT market. By aggregating buy and sell orders on-chain and intelligently routing trades through compatible AMMs, we minimize slippage and maximize efficiency. This approach fosters a more robust and liquid market, benefiting both NFT creators and collectors. QuantumAuction utilizes a modular architecture written in Go, allowing for easy extension and integration with other DeFi protocols. The smart contracts, primarily written in Solidity, are designed to be gas-efficient and secure, adhering to industry best practices.

QuantumAuction goes beyond simply facilitating NFT trades. It provides advanced features such as limit orders, order cancellation, and real-time market data. The on-chain order book ensures transparency and eliminates the risks associated with centralized intermediaries. Furthermore, our composable architecture enables integration with lending protocols, fractionalization platforms, and other DeFi applications, opening up new possibilities for NFT utilization and value creation. This creates a vibrant ecosystem around NFTs, driving adoption and fostering innovation within the space.

## Key Features

*   **On-Chain Order Book Aggregation:** Orders are stored and matched directly on the blockchain, ensuring transparency and immutability. This order book utilizes a priority queue data structure implemented in Go to efficiently manage and match orders.
*   **AMM-Based Liquidity Routing:** Intelligent routing algorithms analyze available liquidity across multiple AMMs and execute trades at the optimal price, minimizing slippage. This leverages smart contracts that interact with existing DEXes like Uniswap and Sushiswap.
*   **Smart Contract Composability:** The modular design allows for seamless integration with other DeFi protocols, enabling features such as NFT-backed loans and fractional ownership. The smart contracts use interfaces extensively to ensure compatibility and upgradability.
*   **Limit Orders & Order Cancellation:** Users can place limit orders at desired prices and cancel orders at any time, providing more control over their trading activity. This functionality is implemented through custom smart contract logic that handles order management.
*   **Real-Time Market Data:** Comprehensive market data, including order book depth, trade history, and price charts, is available through a dedicated API. The API is built using Go's `net/http` package and provides data in JSON format.
*   **Gas-Efficient Design:** The smart contracts are optimized for gas efficiency, minimizing transaction costs for users. Careful attention was paid to storage patterns and loop optimization to achieve this.
*   **Secure Audited Smart Contracts:** Our smart contracts are rigorously audited by independent security experts to ensure the safety of user funds. The audit reports are available upon request.

## Technology Stack

*   **Go:** Backend server, API, and off-chain services are written in Go, chosen for its performance, concurrency, and robust standard library.
*   **Solidity:** Smart contracts that govern the NFT marketplace logic and order book are written in Solidity.
*   **Ethereum Blockchain:** QuantumAuction is deployed on the Ethereum blockchain, providing a secure and decentralized foundation.
*   **IPFS:** NFT metadata is stored on IPFS to ensure immutability and persistence.
*   **PostgreSQL:** Database for storing market data, user information, and other application-specific data.
*   **Web3.js/Ethers.js:** JavaScript libraries for interacting with the Ethereum blockchain from the frontend.

## Installation

1.  **Prerequisites:** Ensure you have Go (version 1.18 or later), Node.js (version 16 or later), and Docker installed.

2.  **Clone the repository:**
    `git clone https://github.com/ezozu/QuantumAuction.git`
    `cd QuantumAuction`

3.  **Install Go dependencies:**
    `go mod download`
    `go mod verify`

4.  **Install Node.js dependencies:**
    `cd frontend`
    `npm install`
    `cd ..`

5.  **Set up the database:**
    Create a PostgreSQL database named `quantumauction`. Update the database connection details in the `.env` file (see Configuration section).

6.  **Compile and deploy smart contracts:**
    Follow the instructions in the `contracts/README.md` file to compile and deploy the smart contracts to a local Ethereum development environment (e.g., Ganache) or a testnet.

7.  **Start the backend server:**
    `go run cmd/quantumauction/main.go`

8.  **Start the frontend:**
    `cd frontend`
    `npm run start`
    `cd ..`

## Configuration

The application utilizes environment variables for configuration. Create a `.env` file in the root directory and populate it with the following variables:



*   `DATABASE_URL`: PostgreSQL database connection string.
*   `ETHEREUM_NETWORK`: Ethereum network to connect to (e.g., mainnet, goerli, sepolia).
*   `ETHEREUM_NODE_URL`: URL of an Ethereum node provider (e.g., Infura, Alchemy).
*   `CONTRACT_ADDRESS`: Address of the deployed QuantumAuction smart contract.
*   `IPFS_GATEWAY`: IPFS gateway URL for retrieving NFT metadata.
*   `PORT`: Port on which the backend server will listen.

## Usage

The frontend provides a user-friendly interface for interacting with the marketplace. Users can browse NFTs, place buy and sell orders, and manage their accounts.

The backend API provides programmatic access to the marketplace functionality. The API is documented using Swagger/OpenAPI. The Swagger UI can be accessed at `/swagger/index.html` after the backend server is started.

Example API endpoint for fetching all NFTs:



Example API endpoint for placing a buy order:



Detailed API documentation, including request/response formats and authentication details, is available in the Swagger UI.

## Contributing

We welcome contributions to QuantumAuction! Please follow these guidelines:

*   Fork the repository and create a branch for your changes.
*   Write clear and concise commit messages.
*   Follow the Go coding style guidelines (e.g., using `go fmt`).
*   Write unit tests for your code.
*   Submit a pull request with a detailed description of your changes.
*   Adhere to the project's code of conduct.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/ezozu/QuantumAuction/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the following individuals and organizations for their contributions to this project:

*   The Ethereum Foundation for providing the underlying blockchain infrastructure.
*   The open-source community for developing the tools and libraries we use.
*   Our team of developers, researchers, and advisors for their hard work and dedication.