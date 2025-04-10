# Используем официальный образ Node.js
FROM node:18 AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем package.json и package-lock.json
COPY ./package*.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем исходный код
COPY . .

# Собираем приложение
RUN npm run build

# Stage 2: Serve the application with Nginx
FROM nginx:alpine

# Build argument for environment
ARG NGINX_CONF=nginx.local.conf

COPY ./ssl /etc/nginx/ssl/

# Copy the appropriate Nginx config file
COPY --from=build /app/build /usr/share/nginx/html
COPY ${NGINX_CONF} /etc/nginx/conf.d/default.conf
#COPY --from=build /usr/src/app/девфлов.рф /etc/nginx/sites-available/девфлов.рф

# Expose port 80
EXPOSE 80 443

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
