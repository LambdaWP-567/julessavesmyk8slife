# Deployment Documentation

This document describes how to deploy the **k8dclusterlife** platform.

## Local Development with Docker Compose

To start the entire stack locally (Backend, Frontend, PostgreSQL, Redis), run:

```bash
docker-compose up --build
```

The application will be available at `http://localhost:8080`.

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_URL` | PostgreSQL Connection String | `postgres://user:pass@db:5432/k8dclusterlife?sslmode=disable` |
| `REDIS_URL` | Redis Connection String | `redis://redis:6379` |

## Kubernetes Deployment with Helm

The Helm chart is located in `helm/k8dclusterlife`.

### Prerequisites

- A running Kubernetes cluster.
- Helm 3.x installed.

### Installation

1. Navigate to the helm directory:
   ```bash
   cd helm/k8dclusterlife
   ```

2. Install the chart:
   ```bash
   helm install k8dclusterlife .
   ```

### Configuration

You can customize the deployment by editing `values.yaml` or by using the `--set` flag.

```bash
helm install k8dclusterlife . --set service.type=LoadBalancer
```

## CI/CD and Releases

The project uses GitHub Actions for continuous integration and automated releases.

- **CI**: Runs on every PR and push to `main`. Includes Go tests, Vitest tests, and basic K8s integration via `kind`.
- **Release**: Triggered by pushes to `main` via `semantic-release`. Automates versioning and GitHub Releases.
