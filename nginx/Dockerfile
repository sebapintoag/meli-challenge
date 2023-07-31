# Official Nginx image (Yes, in this article I always use the `latest`. Kill me!)
FROM nginx:latest

# Copy Nginx configuration file to the container
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Expose port 40
EXPOSE 40