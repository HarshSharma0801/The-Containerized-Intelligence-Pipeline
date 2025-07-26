FROM node:18-alpine

RUN apk --no-cache add wget

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"] 