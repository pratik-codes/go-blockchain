# Blockchain Interaction: Central Server & Miner Nodes

## **Step-by-Step Interaction When a User Adds a Transaction Request:**

```text
  User                             Central Server                          Miners (M1, M2, ...)
    |                                    |                                      |
    |------ Add Transaction -------------|                                      |
    |                                    |                                      |
    |                                    |                                      |
    |                               1. Broadcast Transaction -------------------| 
    |                                    |--------- Receive Transaction --------|
    |                                    |                                      |
    |                                    |                                      |
    |                              2. Add transaction to txn pool               |
    |                                    |                                      |
    |                                    |--------- Broadcast new txns -------->|
    |                                    |                                      |
    |                              3. Wait for block mining                    3. Start mining
    |                                    |                                      |---|
    |                                    |                                      |   |--> Do proof of work
    |                                    |                                      |---|
    |                                    |                                      |
    |                                    |                                      |---|
    |                                    |                                      |   |--> Successfully mine a block
    |                                    |                                      |---|
    |                                    |                                      |
    |                                    | <----- Broadcast new block --------- |
    |                              4. Verify Block                             4. Verify Block
    |                              5. Add Block to chain                       5. Add Block to chain
    |                              6. Adjust difficulty                        6. Adjust difficulty
    |                                    |                                      |
    |                                    |                                      |
    |                                    | ----- Broadcast new block ---------> |
    |                                    |            to                        |
    |                                    |        other miners                  |
    |                                    |                                      |
    |                                    |                                      |
    |------- Notify users of success ----|                                      |
```

## Central Server Responsibilities:
1. Broadcast Transaction:
When a user adds a new transaction, the central server broadcasts it to all connected miners.

2. Manage Transaction Pool:
The server adds the transaction to its own transaction pool and removes transactions that are already included in a mined block.

3. Broadcast New Transactions:
The server keeps miners updated by broadcasting new transactions regularly.

4. Verify and Add Block:
Once a block is mined by any miner, the central server verifies it and broadcasts the new block to the rest of the miners.

5. Adjust Difficulty:
After receiving each new block, the server adjusts the difficulty level to maintain a steady block time.

## Miner Server Responsibilities:
1. Receive Transactions:
Each miner receives new transactions from the central server and maintains its own transaction pool.

2. Mine Blocks:
Miners mine blocks at a fixed interval by performing Proof of Work (PoW). This involves finding a valid hash based on the difficulty level.

3. Broadcast Mined Block:
After successfully mining a block, the miner broadcasts it to the central server, which further distributes it to all other miners.

4. Verify Block:
Each miner verifies the received block and checks if the chain is valid. Smaller or erroneous chains are rejected.

5. Update Blockchain:
The miner updates its local copy of the blockchain and adjusts the mining difficulty based on the central server’s broadcast.
