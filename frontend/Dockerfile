# Используем актуальный LTS образ
FROM node:24-alpine


RUN apk update && apk upgrade

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 5173

# Флаг --host обязателен для работы Vite внутри Docker
CMD ["npm", "run", "dev", "--", "--host"]