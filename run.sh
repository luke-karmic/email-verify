#!/bin/bash

# Set the path to your leads.csv file
LEADS_FILE="/path/to/your/leads.csv"

# Build the Docker image (optional if you have already built it)
docker build -t email-validator .

# Run the Docker container with your CSV file mounted
echo "Running the email validation container..."
CONTAINER_ID=$(docker run -d --rm -v "$LEADS_FILE:/app/leads.csv" email-validator)

# Wait a moment for the container to start processing
echo "Container started with ID $CONTAINER_ID. Viewing logs..."

# View the logs of the running container
docker logs -f "$CONTAINER_ID"

# The container will be removed automatically due to --rm, so no need to clean up
echo "Email validation completed. Container logs shown above."