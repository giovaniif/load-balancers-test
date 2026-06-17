# Load Balancer Test

A hands-on lab for exploring NGINX load balancing strategies using Docker Compose and a minimal Go HTTP server.

## Architecture

```
Client → NGINX (port 80) → app-instance-alpha  (weight 3)
                         → app-instance-beta
                         → app-instance-gama
```

Three identical Go app instances sit behind an NGINX reverse proxy. Each instance responds with its hostname so you can observe which node handled each request.

## Features explored

- **Round robin** — default distribution across all nodes
- **Weighted round robin** — `app-node-1` carries 3× the traffic of the others
- **Health checks** — a node is taken out of rotation for 10 seconds after 2 consecutive failures (`max_fails=2 fail_timeout=10s`)
- **Rate limiting** — configuration is present in `nginx.conf` but commented out; uses the token bucket algorithm with a burst queue and `nodelay`

## Requirements

- Docker
- Docker Compose

## Running

```bash
docker compose up --build
```

NGINX listens on port `80`. Send requests to observe load balancing in action:

```bash
# Single request
curl http://localhost

# Loop to see distribution across instances
for i in $(seq 1 10); do curl -s http://localhost; done
```

## Testing health check failover

Modify one instance to return an error (e.g., call `panic` or return `http.StatusInternalServerError` in `main.go`), rebuild, and run the curl loop. After 2 failures NGINX stops routing to that node for 10 seconds.

## Enabling rate limiting

Uncomment the `limit_req_zone` and `limit_req` directives in `nginx.conf` to enforce a limit of 2 requests/second per IP, with a burst queue of 3.

## Project structure

```
.
├── main.go            # Go HTTP server — responds with its hostname
├── Dockerfile         # Builds the Go app into a minimal Alpine image
├── docker-compose.yml # Spins up NGINX + 3 app instances
└── nginx.conf         # NGINX upstream config (weights, health checks, rate limit)
```
