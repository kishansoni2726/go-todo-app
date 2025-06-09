FROM nginx:alpine

# Remove default config (optional)
RUN rm /etc/nginx/conf.d/default.conf

# Copy your custom config
COPY default.conf /etc/nginx/conf.d/default.conf

# Copy build files
COPY dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
