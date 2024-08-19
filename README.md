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

