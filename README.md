# 🗄️ Own Redis

Own Redis is a lightweight key-value store inspired by Redis.
It runs over UDP, supports SET, GET, and PING commands, stores data in RAM, and allows optional TTL (time-to-live) for keys.

### ✨ Features

✅ SET – store a key-value pair with an optional TTL
✅ GET – retrieve the value of a key
✅ PING – check if the server is responsive
✅ TTL (PX) – keys can expire after a given time
✅ UDP protocol – lightweight communication

### 🚀 How It Works

The server listens on a UDP socket for incoming commands.
SET key value [PX <ms>] – stores a key with an optional TTL (in milliseconds).
GET key – retrieves the value if the key exists.
PING – returns PONG to check server health.
📡 Supported Commands

🔄 PING
Checks if the server is alive.

PING
Response:

PONG
📝 SET
Stores a key-value pair.

SET foo bar
Response:

OK
With TTL (expires in 10 seconds):

SET foo bar PX 10000
📥 GET
Retrieves the value for a key.

GET foo
Response:

bar
If the key doesn’t exist:

(nil)
❌ Error Handling

Missing arguments:
SET foo
Response:

(error) ERR wrong number of arguments for 'SET' command
Invalid TTL:
SET foo bar PX abc
Response:

(error) ERR invalid milliseconds value
### 🛠 Getting Started

go build -o own-redis .

./own-redis --port 8080

Test with netcat

echo "PING" | nc -u -w1 127.0.0.1 8080
# Should return: PONG
📚 Example

echo "SET foo bar" | nc -u -w1 127.0.0.1 8080  # → OK   
echo "PING" | nc -u -w1 127.0.0.1 8080         # → PONG
