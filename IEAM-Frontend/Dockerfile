### STAGE 1: Build ###
FROM node:14.5.0-alpine 
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# ENV PATH /usr/src/app/node_modules/.bin:$PATH
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

### STAGE 2: NGINX ###
# FROM nginx:stable-alpine
# COPY nginx.conf /etc/nginx/nginx.conf
# COPY --from=build /usr/src/app/dist /usr/share/nginx/html
EXPOSE 3000
ENV NUXT_HOST=0.0.0.0

ENV NUXT_PORT=3000
CMD [ "npm", "prod" ]
