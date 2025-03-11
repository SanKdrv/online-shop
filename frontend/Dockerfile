# Используем официальный образ Node.js
FROM node:18

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем package.json и package-lock.json
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем исходный код
COPY . .

# Собираем приложение
RUN npm run build

# Указываем порт, который будет использовать приложение
EXPOSE 3000
#
## Запускаем приложение
#CMD ["npm", "start"]

# Build the React application in production mode
RUN npm start

# Stage 2: Serve the application with Nginx
FROM nginx:alpine

# Build argument for environment
ARG NGINX_CONF=nginx.local.conf

COPY ./ssl /etc/nginx/ssl/

# Copy the appropriate Nginx config file
COPY ${NGINX_CONF} /etc/nginx/conf.d/default.conf

# Copy the built Angular app files to the Nginx HTML folder
COPY --from=build /usr/src/app/dist/frontend/browser /usr/share/nginx/html
COPY --from=build /usr/src/app/девфлов.рф /etc/nginx/sites-available/девфлов.рф

# Expose port 80
EXPOSE 80 443

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
