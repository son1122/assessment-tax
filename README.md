Here's an enhanced and more detailed version of your `README.md` content:

---

### How to Run the Program

#### Setup Instructions

1. **Start the Database**:
   Ensure Docker is installed on your machine. Begin by starting the database using Docker Compose:
   ```
   docker-compose up
   ```

2. **Build the Docker Image**:
   Build a Docker image for the application:
   ```
   docker build -t wongsatorn-tax .
   ```

#### Running the Application

You have two options to run the service: using Docker or directly on your PC:

- **Using Docker**:
  Use the following command to run the service in Docker, ensuring all environment variables are set correctly:
  ```
  docker run -e DATABASE_URL='host=host.docker.internal port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable' -e PORT=8080 -e ADMIN_USERNAME=adminTax -e ADMIN_PASSWORD=admin! wongsatorn-tax
  ```

- **Directly on PC**:
  To run directly on your PC, you can use Go to execute the main application file. This method requires setting
  environment variables either in your system or within your IDE:
  ```
  go run main.go
  ```

#### Ready to Use

After following these steps, your application should be up and running and ready to use. Ensure that you have set all
the required environment variables before starting the service, especially when running directly on your PC.

---