# receipt-processor-challenge

This web service is entirely developed using Golang, utilizing the Gin web framework for routing. It is structured to scale the application effectively and is also production-ready.

## How to run locally

1. Install Golang in your system (I installed latest version 1.22.1)
2. Clone the repository using `git clone https://github.com/DivyaMaddipudi/receipt-processor-challenge.git`
3. Now, go to the root directory `cd receipt-processor-challenge` and run using `go run .` command
   - #### Note: If you encounter any io timeout issues when downloading modules, try to resolve those errors by setting `export GOPROXY=direct`
5. This will start the server locally on the port 8000

## Routes

1. GET `/receipts`: Get all the receipts data available
2. POST `/receipts/process` : This takes JSON receipt and returns JSON object with an unique ID.
3. GET `/receipts/:id/points `: This takes receipt Id and returns the JSON object with calculated points. If the points are calculated for specific receipt, It returns the pre-computed value instead of calculating again

## Features

1. Error handling has been implemented to ensure that any invalid input results in the appropriate error message being returned.
2. The project demonstrates a high degree of organization and modularity in its structure.
3. To optimize performance, redundant calculations of points are avoided by utilizing pre-computed values.
