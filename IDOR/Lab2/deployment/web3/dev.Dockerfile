# Base image
FROM node:16-alpine

# Set the working directory
WORKDIR /usr/src/app

# Install additional dependencies
RUN apk add --no-cache bash

# Copy package.json and package-lock.json to install dependencies
COPY package*.json ./

# Install all dependencies including devDependencies
RUN npm install --include=dev

# Copy the rest of the application files
COPY . .

# Expose the application port
EXPOSE 8082

# Use nodemon for development
CMD ["npm", "run", "dev"]
