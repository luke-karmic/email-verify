#!/bin/bash

# Set the path to your leads.csv file
LEADS_DIR="./leads/"

# Check if the image exists; if not, build it
if ! docker image inspect email-validator > /dev/null 2>&1; then
    echo "Image does not exist. Building the Docker image..."
    docker build -t email-validator .
else
    echo "Image already exists. Skipping build..."
fi

# Run the Docker container with your CSV file mounted
echo "Running the email validation container..."
CONTAINER_ID=$(docker run -d --rm -v "$LEADS_DIR:/app/leads/" email-validator)

# Wait a moment for the container to start processing
echo "Container started with ID $CONTAINER_ID. Viewing logs..."

# View the logs of the running container
docker logs -f "$CONTAINER_ID" | ccze -A

# The container will be removed automatically due to --rm, so no need to clean up
echo "Email validation completed. Container logs shown above."
