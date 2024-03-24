# Proxy Service - Reverse Proxy with API Key
This project implements a reverse proxy server in Go that forwards requests to a target URL. It injects an x-api-key header with a value retrieved from the environment variable API_KEY.

### Features
Forwards requests to a designated target URL.
Injects an x-api-key header with a value from the API_KEY environment variable.
Removes the X-Powered-By header from the response.

### Usage
Set the API_KEY environment variable:
Set the API_KEY environment variable with your Shyft.to API key before running the program.  Instructions for setting environment variables can be found online and vary depending on your operating system.

### Run the program:

```bash
go run main.go
```

### How it Works
The program utilizes the net/http and net/http/httputil packages to create a reverse proxy. The init function parses the target URL and attempts to load environment variables using the translator.shyft.to/proxy/utils package. The createReverseProxy function constructs the reverse proxy with a custom director function that sets the x-api-key header and modifies the request URL. Finally, the main function retrieves the API_KEY from the environment and starts the reverse proxy server on port 8080.

### Dependencies
This code depends on the example.com/proxy/utils package, which presumably contains the LoadEnv function for loading environment variables.
Note: This readme assumes you have the example.com/proxy/utils package available within your project.