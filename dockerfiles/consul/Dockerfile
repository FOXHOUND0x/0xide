# Use the official Consul image from Docker Hub
FROM hashicorp/consul:1.20.0-rc1

# Expose necessary ports
# 8500: HTTP API
# 8600: DNS
EXPOSE 8500 8600

# Command to run Consul in server mode
CMD ["agent", "-server", "-bootstrap-expect=1", "-ui", "-client=0.0.0.0"]
