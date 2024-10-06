# Use the official Traefik image from Docker Hub
FROM traefik:v2.10

# Copy Traefik configuration file
COPY traefik.yml /etc/traefik/traefik.yml

# Expose ports
# 80: HTTP
# 443: HTTPS
# 8080: Traefik Dashboard
EXPOSE 80 443 8080

# Command to run Traefik
CMD ["traefik"]
