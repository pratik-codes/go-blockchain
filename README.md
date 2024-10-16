Central Server:
    Implements the central WebSocket server responsible for broadcasting new transactions to all connected miners.
    Manages connections
    Broadcasts newly minted blocks to rest of the miners 
    Adjusts difficulty accordingly to match the desired average blocktime
    Manages txn pool for the chain and keeps it in sync (removes txns which are included in the block from pool)

Miner Server:
    Implements the miner server code where each miner maintains its own transaction pool.
    Handles block mining and interacts with the central server to receive new transactions.
    Mines blocks according to a fixed interval (e.g., every 30 seconds)

----

TODO: 

[IN PROGRESS] - Central server - A central websocket server that all miners connect to to exchange messages

Miner server -
Code that miners can run to be able to create blocks, do proof of work, broadcast the block via the central server.
Code that verifies the signature, balances and creates / adds a block
Code should reject smaller blockchains/erronours blocks
Should be able to catch up to the blockchain when the server starts

Frontend -
Lets the user create a BTC wallet
Lets the user sign a txn, send it over to one of the miner servers
