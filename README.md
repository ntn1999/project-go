# Customer Management System

Customer Management System is a simple Go application for managing customer information. This application allows you to perform basic operations such as adding, updating, deleting, and viewing customer details.

## Installation

1. Ensure you have Go installed on your computer. If not, you can download Go from the official Go website: https://golang.org/

2. Clone the project from the GitHub repository:
git clone https://github.com/ntn1999/project-go.git
## Running the Application

1. Start the application server with the following command:
go run main.go

2. The application will start on port 3000. You can access the application through a web browser at the following URL:
http://localhost:3000

## Usage

1. **View a List of Customers:** Access `/customers` to see a list of all customers.

2. **View Specific Customer Information:** Access `/customers/{id}` to view information about a specific customer by replacing `{id}` with the actual customer ID.

3. **Add a New Customer:** Use a POST request to `/customers` to add a new customer. Send customer data in JSON format.

4. **Update Customer Information:** Use a PUT request to `/customers/{id}` to update information for a specific customer. Send the updated customer data in JSON format.

5. **Delete a Customer:** Use a DELETE request to `/customers/{id}` to delete a specific customer by replacing `{id}` with the actual customer ID.


