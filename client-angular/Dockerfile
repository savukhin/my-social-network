FROM node:16.15.0-alpine AS build

WORKDIR /usr/src/app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

# RUN npm run build
EXPOSE 4200
# CMD ["npm", "start", "--host", "0.0.0.0"]
CMD ["npm", "run", "start-docker"]

# FROM nginx:1.17.1-alpine
# COPY nginx.conf /etc/nginx/nginx.conf
# COPY --from=build /usr/src/app/dist/client-angular /usr/share/nginx/html
