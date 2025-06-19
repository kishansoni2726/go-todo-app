FROM node:24-alpine

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json for better caching
COPY package*.json ./

# Install dependencies (including dev dependencies for build)
RUN npm install

# Copy source code
COPY . .

# Set build-time environment variables
ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL

# Build the production-ready static site
RUN npm run build

# Install a simple static file server
RUN npm install -g serve

# Create non-root user
RUN addgroup -g 1001 -S nodejs && \
    adduser -S nextjs -u 1001

# Change ownership of the app directory
RUN chown -R nextjs:nodejs /app

# Switch to non-root user
USER nextjs

# Expose port
EXPOSE 3000

# Serve the built application
CMD ["serve", "-s", "dist", "-l", "3000"]