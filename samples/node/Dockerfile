# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set the working directory
WORKDIR /usr/src/app

# @description: Environment (development/production)
# @default: development
ARG NODE_ENV=development

# Install dependencies first for caching
COPY package*.json ./
RUN npm install

# Copy source code
COPY . .

# Expose the application port
# @description: The default port for the Node.js application
EXPOSE 3000

# Start the application
CMD [ "npm", "start" ]
