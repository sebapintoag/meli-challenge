# Use a Node.js image
FROM node:current-alpine

# Declare env
ENV NODE_ENV development

# Set the working directory
WORKDIR /meli

# Copy the package.json and package-lock.json files
COPY package*.json /meli

# Install dependencies
RUN npm install

# Copy everything to the container
COPY . /meli

# Expose port (development)
#EXPOSE 3000

# Build app (production build)
CMD ["npm","run", "build"]

# Run app (development)
#CMD ["npm","start"]