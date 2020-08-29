FROM node:lts-alpine

RUN mkdir /app
WORKDIR /app

COPY ./app /app

RUN yarn install
RUN yarn build

CMD [ "yarn", "start" ]