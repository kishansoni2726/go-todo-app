# Stage 1: Build the app
FROM node:24-alpine AS builder

WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install

# Copy source code
COPY . .

# Set build-time environment variable
ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL

# Build static site
RUN npm run build

# Stage 2: Serve with NGINX
FROM nginx:alpine

# Copy custom nginx config
COPY default.conf /etc/nginx/conf.d/default.conf

# Copy build output to NGINX's public directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose HTTP port
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
