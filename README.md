# Currency Converter CLI(Go)

**cur-converter-cli-go** is a simple command-line interface for checking exchange rates. This project foucuses on working with API, input validation and error handling.

## Features

The CLI provides next functionality:

- 💱 **Display board of current rates**: App displays exchange rates relative to your base currency(USD by default).
- 🧮 **Calculate exchange amounts**: Convert specific amount from your base currency to another.
- 🔁 **Switch base currency**: You can switch the base currency to any supported currency.

## How to use

Since this is a learning project, the easiest way to run it is directly from the command line using:

```
go run .
```

Alternatively, you can build the binary and use it as a standalone CLI tool:

```
go build -o main
./main
```

## Tech Stack

- **Language:** Go
- **API:** [Francfurter API](https://frankfurter.dev/) for fetching the latest currncy rates
