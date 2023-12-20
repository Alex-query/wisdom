Design and implement “Word of Wisdom” tcp server.

• TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work),
the challenge-response protocol should be used.

• The choice of the POW algorithm should be explained.

• After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other
collection of the quotes.

• Docker file should be provided both for the server and for the client that solves the POW challenge

You can run it with docker-compose:

```bash
docker-compose build
docker-compose up
```

I used Hashcash POW algorithm,because it is simple and easy to implement.
Basicaly, we need to defense from DDOS attacks, so we need to make it hard to send a lot of requests to server.
We can measure the time of request processing and make it hard to send a lot of requests in a short time.

We try to defense request get wisdom , but this request is simple and fast, we just return string from array.
if we organise pow defense our server will work slower than without it so that any solution is ok.

Environment variables:

### Конфигурация сервера

| ENV                              | Description                | Default value |
|----------------------------------|----------------------------|---------------|
| WS_CHALLENGE_NUMBER_OF_ZERO_BITS | count numbers of zero bits | 20            |
| WS_CHALLENGE_SALT_LENGTH         | salt length                | 8             |
| WS_CHALLENGE_EXTRA               | extra parametr             | extra         |
| WS_CLIENT_PORT                   | client port                | 1444          |
| WS_CLIENT_HOST                   | client host                | 127.0.0.1     |
| WS_SERVER_PORT                   | server port                | 1444          |
| WS_SERVER_HOST                   | server host                | 127.0.0.1     |
