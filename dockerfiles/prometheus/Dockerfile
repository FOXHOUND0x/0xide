# Use the official Prometheus image as the base image
FROM prom/prometheus:main

# Copy the prometheus.yml file from the host to the container
COPY prometheus.yml /etc/prometheus/prometheus.yml

# Expose port 9090
EXPOSE 9090

# Set the command to run Prometheus with the specified config file
CMD ["--config.file=/etc/prometheus/prometheus.yml"]
