# Stage 1: Build the app
FROM node:24-alpine AS builder

WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install

# Copy source code
COPY . .

# Build static site
RUN npm run build

# Stage 2: Production-ready serve
FROM node:24-alpine

WORKDIR /app

# Install serve globally
RUN npm install -g serve

# Create non-root user
RUN addgroup -g 1001 -S nodejs && adduser -S appuser -u 1001 -G nodejs

# Copy build output
COPY --from=builder /app/dist ./dist

# Set ownership
RUN chown -R appuser:nodejs /app

# Switch to non-root user
USER appuser

# Expose port 3000
EXPOSE 3000

# Production serve with optimizations
CMD ["serve", "-s", "dist", "-p", "3000", "--single"]