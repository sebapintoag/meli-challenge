# Use a Node.js image
FROM node:latest

# Declare env
ENV NODE_ENV development

# Set the working directory
WORKDIR /meli

# Copy the package.json and package-lock.json files
COPY package.json /meli/package.json

# Install dependencies
RUN npm install

# Copy everything to the container
COPY . /meli

# Expose port
EXPOSE 3000

CMD ["npm","start"]