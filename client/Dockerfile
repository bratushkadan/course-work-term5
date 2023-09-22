FROM node:18-alpine3.17 AS build

RUN mkdir /app /build

WORKDIR /build

COPY . .

RUN npm ci
RUN npm run build

RUN cd ./serve && rm -rf node_modules && npm ci --production

RUN mv ./build ./serve ./package* /app

FROM node:18-alpine3.17

COPY --from=build /app /var/www/html/floral

WORKDIR /var/www/html/floral

ENV NODE_ENV=production

CMD ["node", "serve/index.js"]
