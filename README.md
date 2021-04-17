# Berry Miner

Fork from Tellor miner(telliot 4.02) for Berry oracle price feeding.

This is the workhorse of the Miner system as it takes on solving the PoW challenge.  
It's built on Go and utilizes a split structure.  The database piece is a LevelDB that keeps track of all variables (challenges, difficulty, values to submit, etc.) and the miner simply solves the PoW challenge.  This enables parties to split the pieces for optimization.

**The Berry system is a way to push data on-chain.  What the pieces of data are are specificied in the psr.json file. Note that the data corresponds to a specific API.  The berry mining system is set up to pull api data to generate these values to submit on-chain once a correct nonce is mined. These specific apis are just suggestions.  The system is not guarunteed to work for everyone.  It is up to the consnesus of the Berry token holders to determine what a correct value is. As an example, request ID 1 is BTC/USD.  If the api's all go down, it is the responsibility of the miner to still submit a valid BTC/USD price.  If they do not, they risk being disputed and slashed.  For these reasons, please contribute openly to the official Berry miner (or an open source variant), as consensus here is key.  If you're miner gets a different value than the majority of the of the other miners, you risk being punished.**


<p align="center">
    <img src= './public/minerspecs.png' width="450" alt='MinerSpecs' />
</p>


## Berry Deployed Addresses

BSC Mainnet - [0xf859Bf77cBe8699013d6Dbc7C2b926Aaf307F830](https://bscscan.com/address/0xf859bf77cbe8699013d6dbc7c2b926aaf307f830)

## How to build
Now only support linux and MacOS system.

```bash
./release_build.sh
```

## How to mine

Firstly you should config your account in `config.json`

Please input your account address and private key in `config.json`, both are not with prefix `0x`.

### Stake 1000 BRY

Please confirm that your account above have 1000 BRY and some BNB token for fee.

You can use follow command to check your balance:

```bash
./BerryMiner balance
```

It should output as follows:

```bash
0xe010aC6e0248790e08F42d5F697160DEDf97E024
      0.91 BNB
   1034.73 BRY
```

Now you can start stake 1000 BRY for mining with this command:

```bash
./BerryMiner stake deposit
```

Then you can use this command to check your stake status:

```bash
./BerryMiner stake status
```

It should output as follows:

```bash
Staked in good standing since 2021-03-09 00:00:00 +0000 UTC
```

### Start Mining

You can use this command to start mining:

```
nohup ./BerryMiner mine &
```


### DISCLAIMER


    Mine at your own risk.  

    Mining requires you deposit 1000 Berry Tributes.  These are a security deposity.  If you are a malicious actor (e.g. submit a bad value), the community can vote to slash your 1000 tokens.  

    Mining also requires submitting on-chain transactions on Ethereum.  These transactions cost gas (ETH) and can sometimes be signifiant if the cost of gas on Ethereum is high (i.e. the network is clogged).  Please reach out to the community to find the best tips for keeping gas costs under control or at least being aware of the costs. 

    If you are building a competing client, please contact us.  The miner specifications are off-chain and the validity of the mining process hinges on the consensus of the Berry community to determine what proper values are.  Competing clients that change different pieces run the risk of being disputed by the commmunity.  

    There is no guaruntee of profit from mining. 

    There is no promise that Berry Tributes currently hold or will ever hold any value. 

Please join our Telegram for more information and community updates. 


#### Contributors

This repository is maintained by the Berry team - [www.berrydata.co](https://www.berrydata.co)
