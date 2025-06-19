# --- Stage 1: Build React App ---
FROM node:24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Inject API URL from build arg into Vite environment
ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL

# Create .env.production for Vite to read
RUN echo "VITE_API_URL=$VITE_API_URL" > .env.production

# Install dependencies
RUN npm install

# Copy source code
COPY . .

# Build the production-ready static site
RUN npm run build

# --- Stage 2: Serve with NGINX ---
FROM nginx:alpine

# Remove default config if not needed
RUN rm -f /etc/nginx/conf.d/default.conf

# Optional: Add custom nginx.conf (if you have one)
# COPY nginx.conf /etc/nginx/nginx.conf

# Copy built app from previous stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port
EXPOSE 80

# Start NGINX
CMD ["nginx", "-g", "daemon off;"]
