# Use the official Python 3.11.2 base image
FROM python:3.11.2-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the test script and test receipts into the container
COPY test_server.py .
COPY test-receipts ./test-receipts

# Install dependencies
RUN pip install --no-cache-dir requests

# Command to run the Python test script
CMD ["python3", "test_server.py"]