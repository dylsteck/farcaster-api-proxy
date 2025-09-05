# farcaster-api-proxy

A lightweight Go REST API based on [farcaster-api-proxy](https://github.com/pugson/farcaster-api-proxy) by [pugson](https://github.com/pugson) that proxies requests to the Farcaster client API to avoid CORS issues when fetching cast data.

## Usage

### Development

```bash
# Download dependencies
go mod tidy
# Run the server
go run main.go
```

Server starts on port 8080 (or PORT environment variable).

### API Endpoints

All routes proxy to corresponding `client.farcaster.xyz/v2/` endpoints:

- `GET /{username}` → `user-by-username?username=${username}`
- `GET /{username}/{hash}` → `user-thread-casts?castHashPrefix=${hash}&username=${username}&limit=5`
- `GET /fids/{fid}` → `user?fid=${fid}`
- `GET /search/casts` → `search-casts?q=${q}&limit=${limit}`
- `GET /search/users` → `search-users?q=${q}&excludeSelf=${excludeSelf}&limit=${limit}&includeDirectCastAbility=${includeDirectCastAbility}`
- `GET /search/summary` → `search-summary?q=${q}&maxChannels=${maxChannels}&maxUsers=${maxUsers}&maxMiniApps=${maxMiniApps}&maxTokens=${maxTokens}&addFollowersYouKnowContext=${addFollowersYouKnowContext}`
- `GET /health` - Health check

### Examples

```bash
# Get user by username
curl http://localhost:8080/dwr

# Get cast thread data
curl http://localhost:8080/dwr/0x123abc

# Get user by FID
curl http://localhost:8080/fids/3

# Search casts
curl "http://localhost:8080/search/casts?q=hello&limit=10"

# Search users
curl "http://localhost:8080/search/users?q=vitalik&limit=5"

# Search summary
curl "http://localhost:8080/search/summary?q=ethereum"
```

Returns JSON data from Farcaster API with CORS headers and 1-hour cache.

### Deploy

#### Railway (Recommended)

Railway provides an easy one-click deployment for Go applications:

1. Fork this repository
2. Connect your GitHub account to [Railway](https://railway.app)
3. Create a new project and select your forked repository
4. Railway will automatically detect the Go application and deploy it
5. Your API will be available at the generated Railway URL

Railway automatically handles:
- Go build process
- Environment variables (PORT is set automatically)
- HTTPS certificates
- Custom domains (if needed)

#### Manual Deployment

For other platforms, set the `PORT` environment variable for production deployment.

```bash
# Build the application
go build -o farcaster-api-proxy

# Run with custom port
PORT=3000 ./farcaster-api-proxy
```
