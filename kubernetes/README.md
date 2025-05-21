# Kubernetes Deployment for Workout App

This directory contains Kubernetes configuration files for deploying the Workout App and its MySQL database.

## Components

- **Backend Application**
  - Deployment: `backend-deployment.yaml`
  - Service: `backend-service.yaml`
  - ConfigMap: `backend-config.yaml`
  - Secrets: `backend-secrets.yaml`

- **MySQL Database**
  - Deployment: `database-deployment.yaml`
  - Service: `database-service.yaml`
  - Secrets: `database-secrets.yaml`
  - Storage: `database-storage.yaml`

- **Ingress**
  - Configuration: `ingress.yaml`

## Deployment Instructions

### Prerequisites

- Kubernetes cluster (local or cloud-based)
- kubectl configured to communicate with your cluster
- Docker images for the application (built and pushed to a registry)

### Deployment Steps

1. **Update image references**

   Update the image reference in `backend-deployment.yaml` to point to your Docker registry:
   ```yaml
   image: your-registry/workout-app:latest
   ```

2. **Update domain name**

   Update the host in `ingress.yaml` to your actual domain:
   ```yaml
   host: your-actual-domain.com
   ```

3. **Apply configurations**

   Option 1: Apply individual files:
   ```bash
   kubectl apply -f database-storage.yaml
   kubectl apply -f database-secrets.yaml
   kubectl apply -f database-deployment.yaml
   kubectl apply -f database-service.yaml
   kubectl apply -f backend-secrets.yaml
   kubectl apply -f backend-config.yaml
   kubectl apply -f backend-deployment.yaml
   kubectl apply -f backend-service.yaml
   kubectl apply -f ingress.yaml
   ```

   Option 2: Apply all at once using kustomize:
   ```bash
   kubectl apply -k .
   ```

4. **Verify deployment**

   Check if all pods are running:
   ```bash
   kubectl get pods
   ```

   Check services:
   ```bash
   kubectl get svc
   ```

   Check ingress:
   ```bash
   kubectl get ingress
   ```

## Configuration

### Scaling

To scale the backend application:
```bash
kubectl scale deployment workout-app --replicas=5
```

### Updating Secrets

If you need to update secrets:
1. Generate new base64 encoded values:
   ```bash
   echo -n "new_password" | base64
   ```
2. Update the relevant secret YAML file
3. Apply the changes:
   ```bash
   kubectl apply -f backend-secrets.yaml
   ```
4. Restart the affected deployments:
   ```bash
   kubectl rollout restart deployment workout-app
   ```

## Troubleshooting

### Checking Logs

```bash
kubectl logs deployment/workout-app
kubectl logs deployment/workout-db
```

### Debugging Pods

```bash
kubectl describe pod [pod-name]
```

### Accessing the Database

```bash
kubectl port-forward svc/workout-db 3306:3306
```
Then connect to the database using a MySQL client on localhost:3306.