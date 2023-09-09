# Country Dialing Code API Documentation

This simple Go (Golang) API provides information about countries, including their names, dialing codes, and country codes. It reads this information from a JSON file and exposes an endpoint to retrieve these data.

## Configuration

Before running the API, you need to set up an `.env` file in the project's root directory with the following variables:

- `PORT`: The port on which the server will start (optional, default: 8080).

You can create an `.env` file based on the provided example:

```plaintext
PORT=8080
```

## How to Run

Make sure you have Go installed on your system. You can run the API as follows:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/your-repository.git
   cd your-repository
   ```

2. Install dependencies using `go get`:

   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u github.com/joho/godotenv
   ```

3. Start the server:

   ```bash
   go run main.go
   ```

This will start the server on the specified port or on the default port 8080 if the `PORT` environment variable is not set.

## Endpoint

The API has a single endpoint for retrieving information about countries with dialing codes:

- **GET /countries-with-dialing-code**

  This endpoint returns a JSON containing a list of countries with their names, dialing codes, and country codes.

  Example response:

  ```json
  [
    {
      "name": "Brazil",
      "dial_code": "+55",
      "code": "BR"
    },
    {
      "name": "United States",
      "dial_code": "+1",
      "code": "US"
    },
    ...
  ]
  ```

## Contribution

If you wish to contribute to this project, feel free to fork the repository, make the necessary changes, and submit a pull request.

## License

This project is licensed under the MIT License. Please refer to the [LICENSE](LICENSE) file for more details.

We hope this documentation helps you set up and use the Country Dialing Code API. If you have any questions or need further information, please don't hesitate to get in touch.

## Created by:
**[Francisco Inoque](franciscoinoque.tech)**
