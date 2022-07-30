# Load-Balanced-Service

A simple implementation of a JWT secured service with 3 replicas which is load balanced (in a round robin manner) via HAProxy.

# How to use:
- `git clone` this repository.
- `cd` inside.
- run `docker compose up` (make sure Docker is up and running).
- in order to generate new JWT token an GET request is needed to `http:://127.0.0.1/get-token` with header: `Api_key : 1234`.
- copy the generated token and add it as an header `Token: <JWT>` when ever querying the API.
- access API via `http://127.0.0.1/`.
