### STAGE 1: Build ###
FROM node:14.5.0-alpine as build
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ENV PATH /usr/src/app/node_modules/.bin:$PATH
COPY package.json /usr/src/app/package.json
RUN yarn install --silent
COPY . /usr/src/app
RUN yarn run generate

### STAGE 2: NGINX ###
FROM nginx:stable-alpine
COPY --from=build /usr/src/app/dist /usr/share/nginx/html
ARG BASE_URL
ENV BASE_URL $BASE_URL
ARG API_URL
ENV API_URL $API_URL
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
