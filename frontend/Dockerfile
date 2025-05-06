FROM node:23.6.1-alpine AS builder

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn install

COPY . ./

ARG VITE_BACKEND_URI
ENV VITE_BACKEND_URI="${VITE_BACKEND_URI}"

RUN yarn build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
