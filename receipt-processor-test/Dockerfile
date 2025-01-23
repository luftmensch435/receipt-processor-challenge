FROM python:3.11.2-slim

WORKDIR /test

COPY test_server.py .
COPY test-receipts ./test-receipts

RUN pip install requests
CMD ["python3", "test_server.py"]