# ListingD

A simple API for managing company information.

## Requirements

- Go 1.17 or higher
- Docker and Docker Compose (optional)
- Kubernetes and kubectl (optional)

## Configuration

Create a `config.yaml` file in the project's root directory with the following content:

```yaml
server:
  host: "localhost"
  port: 8080
database:
  host: "localhost"
  port: 5432
  user: "db_user"
  password: "db_password"
  dbname: "listingd"
  sslmode: "disable"
jwt:
  secret: "your_jwt_secret"
  duration: 3600
```
An example `config.yaml` file is included in the repository once cloned

## Setup

1. Clone the repository:

```sh
git clone https://github.com/APoniatowski/listingd.git
cd listingd
```

2. Install dependencies:

```sh
go mod download
```

3. Build the application:

```sh
go build -o listingd
```

4. Run the application:

```sh
./listingd
```

The API will be available at `http://localhost:8080`.

## Using Docker and Docker Compose

1. Build the Docker image:
```sh
docker build -t listingd .
```

2. Create a docker-compose.yml file:

(Include the Docker Compose content you created earlier in this step.)

3. Start the services:

```sh
docker-compose up
```

The API will be available at `http://localhost:8080`.

## Deploying to Kubernetes

1. Build the Docker image and push it to a container registry:
```sh
docker build -t <your-registry>/listingd:<tag> .
docker push <your-registry>/listingd:<tag>
```

Replace <your-registry> with your container registry's address and <tag> with the desired image tag.

2. Update the `listingd-deployment.yaml` file:

Replace <your-registry> and <tag> in the image field with the appropriate values from step 1.

3. Apply the Kubernetes manifests:
```sh
kubectl apply -f k8s/listingd-deployment.yaml
kubectl apply -f k8s/listingd-service.yaml
```

4. Access the API:
Run the following command to get the IP address or domain of the Kubernetes LoadBalancer service:
```sh
kubectl get svc listingd-service
```

The API will be available at `http://<IP-or-domain>:8080`.