# Backend for Currency Converter


## How to run

- Pull the image

```docker 
docker pull docker.pkg.github.com/jasonstanleyyoman/currency_converter_be/currency_api:latest
```

- Run the image with SIMPLE_AUTH_TOKEN set as Environment Variable
```docker 
docker run --name image-name -e SIMPLE_AUTH_TOKEN=YOUR_TOKEN -d docker.pkg.github.com/jasonstanleyyoman/currency_converter_be/currency_api:latest
```
Change YOUR_TOKEN to your token choice

## Endpoint

- /v1/currency/

Get all currency rates with EUR as its base

Example response:

```json
{
  "status": 200,
  "timestamp:" 1624004704,
  "data": {
    "rates": [
      {
        "symbol": "XAF",
        "rate": 656.378491,
        "long": "Central African CFA franc"
      },
      {
        "symbol": "IDR",
        "rate": 17252.999581,
        "long": "Indonesian rupiah"
      },
      ...
    ]
  }

}
```

- /v1/currency/convert?from=FROM_SYMBOL&to=TO_SYMBOL&amount=AMOUNT

Example: /v1/currency/convert?from=IDR&to=EUR&amount=15000

Response:

```json
{
  "status": 200,
  "timestamp": 1624004953,
  "data": {
    "query": {
      "from": "IDR",
      "to": "EUR",
      "amount": 15000
    },
    "result": {
      "from": "IDR",
      "from_long": "Indonesian rupiah",
      "to": "EUR",
      "to_long": "European euro",
      "amount": 15000,
      "result": 0.8694140360681899
    }
  }
}
```

Example: /v1/currency/convert?from=IDR&to=EUR,IDR&amount=15000

Response:

```json
{
  "status": 200,
  "timestamp": 1624004981,
  "data": {
    "query": {
      "from": "IDR",
      "to": "EUR,IDR",
      "amount": 15000
    },
    "results": [
      {
        "from": "IDR",
        "from_long": "Indonesian rupiah",
        "to": "EUR",
        "to_long": "European euro",
        "amount": 15000,
        "result": 0.8694140360681899
      },
      {
        "from": "IDR",
        "from_long": "Indonesian rupiah",
        "to": "IDR",
        "to_long": "Indonesian rupiah",
        "amount": 15000,
        "result": 15000
      }
    ]
  }
}
```

## Others
Frontend code repository can be found on [here](https://github.com/jasonstanleyyoman/currency_converter_fe)