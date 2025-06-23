# 🧩 Full Stack TODO Application - Kubernetes Deployment

This repository contains the manifests and deployment instructions for a containerized full stack TODO application deployed on a Kubernetes cluster using **Amazon Elastic Container Registry (ECR)** for image storage.

---

## 📌 Tech Stack

- **Frontend**: React (Vite) UI
- **Backend**: Node.js with Fiber (or Express)
- **Container Registry**: Amazon ECR
- **Kubernetes**: Self-managed or EKS
- **Ingress Controller**: NGINX (optional)

---

# 🛠️ Kubernetes Deployment with Amazon ECR Integration

This guide outlines the steps to deploy your application in a Kubernetes cluster using container images stored in **Amazon Elastic Container Registry (ECR)**. It covers secret creation, deployment, and best practices.

---

## 📦 Prerequisites

- AWS CLI configured (`aws configure`)
- Kubernetes cluster access (`kubectl` context set)
- IAM permissions for `ecr:GetAuthorizationToken`
- Docker images pushed to ECR

---

## 🔑 Create Docker Registry Secret for ECR

ECR authentication tokens are valid for **12 hours**. To pull images from ECR, Kubernetes needs a Docker registry secret.

Run the following command to create (or update) the ECR image pull secret:

```bash
kubectl create secret docker-registry ecr-registry-secret \
  --docker-server=<aws_account_id>.dkr.ecr.<region>.amazonaws.com \
  --docker-username=AWS \
  --docker-password="$(aws ecr get-login-password --region <region>)" \
  --namespace=default \
  --dry-run=client -o yaml | kubectl apply -f -