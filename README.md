# OMS [Order Management System]

### Algorithm 
To handle the race condition, I am using ACID property of MySQL. Mysql Database will never make two transactions at the same time. 
It will make one transaction and will not allow the item to be affected again till after the transaction is complete so it will 
always choose one of the transactions to process (being the same time..) before it allows the item to have another transaction processed. 
Meaning the item would come back as out of stock to the second transaction. In this way, no DB lock is needed.

So ACID property saved me from race condition but what about the live traffic at DB which will certainly downgrade the database hence our whole System.
To handle traffic at the database, I am using Kafka. So whatever traffic is coming to site I am pushing it to Kafka and Kafka consumer is listening 
the traffic and calling database sequentially.

PFB the step by step execution [From Add to Cart To Place Order]

1. Receive addToCart request via HTTP API.
2. Call Inventory service to check the stock of Item from Cache.
3. If item is in stock, then push cart request to Kafka and set order status INIT in redis.
4. Receive placeOrder request via HTTP API.
5. Check for BLOCK and FAILED status for requested orderID for 3 minutes.(In background, this service is listening Kafka messages and calling 
Inventory service API to check and block the Inventory so if Item is in stock then I am blocking that Item and putting the order in BLOCK 
state in Redis and if Item is not is stock then I am marking order FAILED.)
6. If no order or failed order is found in Redis then I am failing the cart.
7. If blocked order is found then I am initiating payment.
8. Then setting payment fail/pass status in Redis.
9. If order status is BLOCK or PAYMENT_FAILED then I am calling Inventory service API to add the blocked Items back to Inventory 
and failing the cart.
10. If order status is PAYMENT_SUCCESS then I am calling inventory to check negative Inventory and if Inventory is in negative then adding items 
back to inventory and failing the cart, if not then your order is placed successfully.
