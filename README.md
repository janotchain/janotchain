# Janote

## About Janote

JanotChain is an innovative solution to the scalability problem with the Ethereum blockchain.
Our mission is to be a leading force in building the Internet of Value, and its infrastructure.
We are working to create an alternative, scalable financial system which is more secure, transparent, efficient, inclusive, and equitable for everyone.

JanotChain relies on a system of 150 Masternodes with a Proof of Stake Voting consensus that can support near-zero fee, and 2-second transaction confirmation times.
Security, stability, and chain finality are guaranteed via novel techniques such as double validation, staking via smart-contracts, and "true" randomization processes.

Janotchain supports all EVM-compatible smart-contracts, protocols, and atomic cross-chain token transfers.
New scaling techniques such as sharding, private-chain generation, and hardware integration will be continuously researched and incorporated into Janotchain's masternode architecture. This architecture will be an ideal scalable smart-contract public blockchain for decentralized apps, token issuances, and token integrations for small and big businesses.





## Building the source

Janotchain provides a client binary called `janot` for both running a masternode and running a full-node.
Building `janot` requires both a Go (1.7+) and C compiler; install both of these.

Once the dependencies are installed, just run the below commands:

```bash
$ git clone https://github.com/janotchain/janotchain 
$ cd janotchain
$ make janot
```

Alternatively, you could quickly download our pre-complied binary from our [github release page](https://github.com/janotchain/janotchain/releases)

## Running `Janot`

### Running a Janot masternode



### Attaching to the Janotchain test network



### Running `Janot` locally

#### Download genesis block
$GENESIS_PATH : location of genesis file you would like to put
```bash
export GENESIS_PATH=path/to/genesis.json
```
- Testnet
```bash
curl -L https://raw.githubusercontent.com/janotchain/janotchain/master/genesis/testnet.json -o $GENESIS_PATH
```

- Mainnet
```bash
curl -L https://raw.githubusercontent.com/janotchain/janotchain/master/genesis/mainnet.json -o $GENESIS_PATH
```

#### Create datadir
- create a folder to store janotchain data on your machine

```bash
export DATA_DIR=/path/to/your/data/folder 
mkdir -p $DATA_DIR/janot
```
#### Initialize the chain from genesis

```bash
janot init $GENESIS_PATH --datadir $DATA_DIR
```

#### Initialize / Import accounts for the nodes's keystore
If you already had an existing account, import it. Otherwise, please initialize new accounts 

```bash
export KEYSTORE_DIR=path/to/keystore
```

##### Initialize new accounts
```bash
janot account new \
  --password [YOUR_PASSWORD_FILE_TO_LOCK_YOUR_ACCOUNT] \
  --keystore $KEYSTORE_DIR
```
    
##### Import accounts
```bash
janot  account import [PRIVATE_KEY_FILE_OF_YOUR_ACCOUNT] \
     --keystore $KEYSTORE_DIR \
     --password [YOUR_PASSWORD_FILE_TO_LOCK_YOUR_ACCOUNT]
```

##### List all available accounts in keystore folder

```bash
janot account list --datadir ./  --keystore $KEYSTORE_DIR
```

#### Start a node
##### Environment variables
   - $IDENTITY: the name of your node
   - $PASSWORD: the password file to unlock your account
   - $YOUR_COINBASE_ADDRESS: address of your account which generated in the previous step
   - $NETWORK_ID: the networkId. Mainnet: 88. Testnet: 89
   - $BOOTNODES: The comma separated list of bootnodes. Find them [here](https://docs.janotchain.com/general/networks/)
   - $WS_SECRET: The password to send data to the stats website. Find them [here](https://docs.janotchain.com/general/networks/)
   - $NETSTATS_HOST: The stats website to report to, regarding to your environment. Find them [here](https://docs.janotchain.com/general/networks/)
   - $NETSTATS_PORT: The port used by the stats website (usually 443)
    
##### Let's start a node
```bash
janot  --syncmode "full" \    
    --datadir $DATA_DIR --networkid $NETWORK_ID --port 30303 \   
    --keystore $KEYSTORE_DIR --password $PASSWORD \    
    --rpc --rpccorsdomain "*" --rpcaddr 0.0.0.0 --rpcport 8545 --rpcvhosts "*" \   
    --rpcapi "db,eth,net,web3,personal,debug" \    
    --gcmode "archive" \   
    --ws --wsaddr 0.0.0.0 --wsport 8546 --wsorigins "*" --unlock "$YOUR_COINBASE_ADDRESS" \   
    --identity $IDENTITY \  
    --mine --gasprice 2500 \  
    --bootnodes $BOOTNODES \   
    --ethstats $IDENTITY:$WS_SECRET@$NETSTATS_HOST:$NETSTATS_PORT 
    console
```


##### Some explanations on the flags   
```
--verbosity: log level from 1 to 5. Here we're using 4 for debug messages
--datadir: path to your data directory created above.
--keystore: path to your account's keystore created above.
--identity: your full-node's name.
--password: your account's password.
--networkid: our network ID.
--janot-testnet: required when the networkid is testnet(89).
--port: your full-node's listening port (default to 30303)
--rpc, --rpccorsdomain, --rpcaddr, --rpcport, --rpcvhosts: your full-node will accept RPC requests at 8545 TCP.
--ws, --wsaddr, --wsport, --wsorigins: your full-node will accept Websocket requests at 8546 TCP.
--mine: your full-node wants to register to be a candidate for masternode selection.
--gasprice: Minimal gas price to accept for mining a transaction.
--targetgaslimit: Target gas limit sets the artificial target gas floor for the blocks to mine (default: 4712388)
--bootnode: bootnode information to help to discover other nodes in the network
--gcmode: blockchain garbage collection mode ("full", "archive")
--synmode: blockchain sync mode ("fast", "full", or "light". More detail: https://github.com/janotchain/janotchain/blob/master/eth/downloader/modes.go#L24)           
--ethstats: send data to stats website
```
To see all flags usage
   
```bash 
janot --help
```

#### See your node on stats page
   - Testnet: https://stats.testnet.janotchain.com
   - Mainnet: http://stats.janotchain.com


## Contributing and technical discussion

Thank you for considering to try out our network and/or help out with the source code.
We would love to get your help; feel free to lend a hand.
Even the smallest bit of code, bug reporting, or just discussing ideas are highly appreciated.

If you would like to contribute to the janotchain source code, please refer to our Developer Guide for details on configuring development environment, managing dependencies, compiling, testing and submitting your code changes to our repo.

Please also make sure your contributions adhere to the base coding guidelines:

- Code must adhere to official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e uses [gofmt](https://golang.org/cmd/gofmt/)).
- Code comments must adhere to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
- Pull requests need to be based on and opened against the `master` branch.
- Any code you are trying to contribute must be well-explained as an issue on our [github issue page](https://github.com/janotchain/janotchain/issues)
- Commit messages should be short but clear enough and should refer to the corresponding pre-logged issue mentioned above.

For technical discussion, feel free to join our chat at [Gitter](https://gitter.im/janotchain/janotchain).
