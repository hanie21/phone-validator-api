# Phone Number Validator API

This project is a phone number validation API built using **Go** with the **Gin framework** and **libphonenumber** for parsing and validating phone numbers.

---

## Installation and Running the Application

### Prerequisites:
- **Go** must be installed on your system (version 1.20 or later).

### Steps to Run the Application Locally:

1. **Download or Copy the Project Files**:
   - Ensure all project files are downloaded or copied onto your local machine. The project should include:
     - `main.go`
     - `/handlers` directory (for API handlers)
     - `/services` directory (for business logic and validation)
     - `go.mod` and `go.sum` files (for dependency management)

2. **Set Up Dependencies**:
   - Open a terminal and navigate to the project directory where `go.mod` is located.
   - Run the following command to download and install dependencies:
     ```bash
     go mod tidy
     ```

3. **Run the Application**:
   - Once the dependencies are installed, start the server by running:
     ```bash
     go run main.go
     ```

   - The API will now be running on `http://localhost:8080`.

### API Usage:
You can test the API by sending a request using `curl` or an API client like Postman or Insomnia.

#### Example Request:
```bash
curl "http://localhost:8080/v1/phone-numbers?phoneNumber=915%20872200&countryCode=ES"
```

---

## Explanation of Choices

### Programming Language: **Go (Golang)**
- **Why Go?** Go is ideal for building high-performance web services due to its concurrency model, simplicity, and speed. It's widely used for API development and backend services.

### Framework: **Gin**
- **Why Gin?** The Gin framework is a lightweight, high-performance HTTP router in Go that simplifies the development of APIs and web services. It allows for fast request routing and is easy to use.

### Library: **libphonenumbers**
- **Why libphonenumbers?** I chose this library for its proven capabilities in parsing and validating phone numbers in international formats. It handles a wide variety of phone number formats with ease, ensuring global compatibility.

---

## Deployment to Production

Although this project is designed to be run locally, here's how you could deploy it to production:

1. **Containerization (Docker)**:
   - Create a `Dockerfile` to containerize the application. This ensures consistency across environments.
   - Example `Dockerfile`:
     ```dockerfile
     FROM golang:1.20
     WORKDIR /app
     COPY . .
     RUN go mod tidy
     RUN go build -o phone-validator .
     CMD ["./phone-validator"]
     ```

2. **Cloud Platform**:
   - Deploy the containerized application to a cloud service such as **AWS ECS**, **Google Cloud Run**, or **Azure Container Instances**.
   - These platforms can handle scaling, load balancing, and high availability automatically.

3. **CI/CD Pipeline**:
   - Use CI/CD tools like **GitHub Actions**, **CircleCI**, or **Jenkins** to automate the build, test, and deployment process.
   
4. **Environment Variables**:
   - In production, environment variables can be used to configure aspects like the server port, logging level, or API keys.


## Assumptions Made

1. **Phone Number Formatting**: The phone number can be provided with or without the `+` symbol, and spaces are only allowed between the country code, area code, and local number.
2. **Country Code**: If a phone number doesn’t include the country code, it can be provided separately using the `countryCode` parameter in **ISO 3166-1 alpha-2** format (e.g., `US` for the United States).
3. **Spaces**: The API only allows spaces in specific parts of the phone number (e.g., between country code and area code).
4. **Port**: The application runs on port `8080` by default, and the user can make requests to `http://localhost:8080`.

---

## Improvements to Be Made

1. **Better Error Handling**: 
   - Enhance error messages to provide more detail (e.g., what part of the phone number is invalid). This would make debugging for API consumers easier.

   - Edge cases were not covered in testing.
   
2. **Logging**:
   - Implement structured logging (e.g., using **logrus**) for better insight into the server's operation and easier troubleshooting.

3. **Rate Limiting**:
   - Implement rate limiting to prevent abuse or overuse of the API, particularly for public-facing deployments.
   
4. **Configuration Management**:
   - Use environment variables to manage configuration (like port number, logging levels, and any external service API keys) for better flexibility in different environments.

5. **Security Enhancements**:
   - Implement secure headers, SSL/TLS encryption, and input validation to ensure the API is secure and resistant to attacks like injection and XSS.