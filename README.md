# Back-end Assignment - MoMoney

# Problem Statement

We have a third-party API from [typicode](https://jsonplaceholder.typicode.com/) that we want to use to serve data on our application. Let’s assume that this API is chargeable per invocation but always serves the same data for a particular ID as the data is immutable.

# Project Goal

To create an efficient server that acts as a proxy for APIs from [typicode](https://jsonplaceholder.typicode.com/) and keeps the cost in check by caching the responses and serving from the cache whenever possible.

# Requirements

- Write a server that
    1. Exposes the following GET APIs:
        1. `/posts/:id` that fetches the data from [https://jsonplaceholder.typicode.com/posts/:id](https://jsonplaceholder.typicode.com/posts/:id) and returns it to the user.
        2. `/todos/:id` that fetches the data from [https://jsonplaceholder.typicode.com/todos/:id](https://jsonplaceholder.typicode.com/todos/:id) and returns it to the user.
    2. Write a generic middleware for caching of APIs that should:
        1. cache the data being served over API so that if the next time the same API is called with the same parameters, the response is served from the cache instead.
        2. be configurable - can be enabled/disabled for a particular route.
- Dockerize the project.
- The project should have proper code organization, comments, logging, validation, and error handling.

# Solution  Cache Server

The back-end server is implemented with the required features to cache responses.

In essence, the cache server acts as a proxy for APIs from [typicode](https://jsonplaceholder.typicode.com/) by caching the responses, thereby keeping the cost in check.

To get the expected data, you need to pass the `id` and `cache` parameters in the endpoints. The allowed endpoints are:

- [https://cacheserver.ayushr1.repl.co/posts/](https://cacheserver.ayushr1.repl.co/posts/)
- [https://cacheserver.ayushr1.repl.co/todos/](https://cacheserver.ayushr1.repl.co/todos/)

The server fetches the response from the API on the first call, and if the `cache` parameter is set to 1 and the cached data is available, subsequent calls are fetched from the cache.

The caching middleware is generic and configurable, allowing you to enable or disable it for a particular route. The project has proper code organization, comments, logging, validation, and error handling.

You can access the Cache Server following this link: [https://cacheserver.ayushr1.repl.co/](https://cacheserver.ayushr1.repl.co/).

Overall, the Cache Server provides an efficient solution to serve data on your application while keeping the cost in check.

# Running the Project Locally

Prerequisites:

- Docker must be installed.

Steps:

1. Clone this repository:

```
git clone https://github.com/AyushR1/GoLang-Server

```

1. Navigate to the cloned directory:

```
cd GoLang-Server

```

1. Run the following command in the terminal:

```
bash run.sh

```

The server will be running locally.

# Presentation

# Video

[https://youtu.be/0FrxTzyyNZg](https://youtu.be/0FrxTzyyNZg)

# Screenshots

## Home

![Untitled](Back-end%20Assignment%20-%20MoMoney%20521cb6050328472ea58aa18132d10223/Untitled.png)

## First Api call for posts

![Untitled](Back-end%20Assignment%20-%20MoMoney%20521cb6050328472ea58aa18132d10223/Untitled%201.png)

## Subsequent calls from cache

![Untitled](Back-end%20Assignment%20-%20MoMoney%20521cb6050328472ea58aa18132d10223/Untitled%202.png)