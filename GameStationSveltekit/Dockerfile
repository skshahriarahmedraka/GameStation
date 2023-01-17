FROM node:19 as build

ENV NODE_ENV=production 


WORKDIR /app
COPY . .
COPY package.json package-lock.json svelte.config.js ./
RUN npm run build 


FROM node:19-alpine3.16

WORKDIR /app
COPY --from=build /app .


EXPOSE 3000
CMD ["node", "./build/index.js"]
