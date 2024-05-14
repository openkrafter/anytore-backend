# Anytore (Backend)

## Introduction to Anytore

Anytore is a simple web application for recording your daily training. The backend is in this repository, and the frontend can be found at the following repository: [https://github.com/openkrafter/anytore](https://github.com/openkrafter/anytore)

## Usage

You can deploy it to a local environment using Docker or set it up on AWS.

### Local Deployment Steps

> **Note:** The following steps require Docker, Docker Compose, and AWS CLI to be installed.

1. **Backend Application Local Setup**

   - Clone the repository.

     ```sh
     git clone https://github.com/openkrafter/anytore-backend.git
     ```

   - Configure AWS CLI settings.

     ```sh
     aws configure set aws_access_key_id dummy
     aws configure set aws_secret_access_key dummy
     aws configure set region ap-northeast-1
     ```

     > **Note:** The values can be arbitrary. If AWS credentials are already set up, reconfiguration is not necessary.

   - Run the local environment setup script.

     ```sh
     cd anytore-backend
     platform/local/local-setup.sh
     ```

2. **Frontend Application Local Setup**
   - Refer to the [frontend application repository](https://github.com/openkrafter/anytore).

## Testing

Refer to the [frontend application repository](https://github.com/openkrafter/anytore).

## Cleanup

To delete Anytore (Backend), run the following script:

```sh
platform/local/local-teardown.sh
```

## Tech Stack

- macOS / Linux
- Docker
- Golang
- Gin
- MySQL
- Amazon DynamoDB
