# go_frameworks

Source files from the course: [Choose Your Go Framework: Chi Router, FastHTTP, Fiber, Echo, Gin Gonic, Go Kratos
](https://www.linkedin.com/learning/choose-your-go-framework-chi-router-fasthttp-fiber-echo-gin-gonic-go-kratos) taught by [Akhil Sharma](https://www.linkedin.com/learning/instructors/akhil-sharma?u=76737724). Every folder consists of the different frameworks + default package for http requests.

## Comparison
Framework selection will always depend on the project restrictions, business rules and so on. Here is a quick overview of the main differences between the 6 most popular golang web frameworks: <br>

| Framework  | Main characteristic            | Brief explanation                                                                                                                                                |
| ---------- | -----------------------------  | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Chi router | Lightweight and flexible       | No external dependencies, good for Proofs of Concepts and fast iterations. Short learning curve                                                                  |
| Echo       | Simplicity and ease to use     | Great for MVPs, Scalable, Robust error handling, high performance, scalable, extensible and minimalist                                                           |
| FastHTTP   | High performance               | Not a complete framework but designed for handling 1000s of small to medium requests/s with a consistently low response time. It's 10 times faster than net/http |
| Fiber      | Express.js style, fast & light | Built on top of FastHttp, implements the same way of working of Express but with golang. Perfect for MVPs and scalable systems                                   |
| Gin        | Loads of features              | Stable features, scalable, production-level apps, ease of maintainance, larger than Echo and Fiber and not suitable for small projects                           |
| Kratos     | Scalable microservices         | Suitable for large-production microservices, service discovery, monitoring, supports gRPC, HTTP and any other protocol, distributed systems, etc.                |
<br>

### FastHTTP vs net/http
| FastHTTP                                                                                  | net/http                                                                                                             |
| ----------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| Can reuse existing request objects    | Lifetime of request object isn't limited by request handler exec time, the server creates a new one per each request & needs to create a new response object per request |
| Can reuse response                                                                        | Requires creating a new response object per each request                                                             |
| Optimizes even little things like returnin a slice of bytes ([]byte) rather than a string | Doesn't support optimizations                                                                                        |

NOTE: conversion from byte to string isn't free and requires memory allocation and copy

# References
