
# Golang Microservice: URL Shortener

This is a microservice that provides a URL shortening functionality. Given a long URL, the service generates a short URL that redirects to the original URL.

The microservice is built using the Go programming language and the Gorilla Mux package for routing.

## Usage

### Creating a short URL

To create a short URL, send a POST request to the `/shorten` endpoint with a JSON payload containing the long URL:

```
curl -X POST -H "Content-Type: application/json" -d '{"url": "http://example.com/very-long-url"}' http://localhost:8080/shorten
``` 

The response will be a JSON object containing the short URL:

```
{"short_url":"http://localhost:8080/7VZv2vm9"}
``` 
### Accessing a short URL

To access a short URL, simply navigate to the short URL in your web browser. The service will redirect you to the original URL.

### Running the service

To run the service, use the following rules that described in Makefile:

```
make start
```

## Dependencies

The microservice has the following dependencies:

-   Go (version 1.16 or later)
-   Gorilla Mux package (version 1.8 or later)
-   Teris shortid package (version 2.1 or later)

You can install the dependencies using the following command:

```
make setup
```