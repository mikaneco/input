# input
## Description
This is a simple program that posts a note to a
personal api endpoint. The note is stored in a
database and can be retrieved by the user.

## Installation
1. Clone the repository
You can clone the repository using the following command.
```bash
git clone git@github.com:mikaneco/input.git
```
2. Install the required packages
You can install the required packages using the following command.
```bash
go mod tidy
```

3. Run the program
You can run the program using the following command.
```bash
go run main.go
```

4. Build the program
You can also build the program using the following command and run the executable.
```bash
go build
```


## Usage
1. Run the program
2. Enter a note
3. The note will be stored in the database

## ApiEndpoint Requirements
The api endpoint must be a POST request that accepts a JSON object with the following format.
```json
{
    "input": "This is a note."
}
```

## License
This project is licensed under the MIT License.

## Support
If you have any questions, please contact me via github issues.

