FROM node:18

WORKDIR /client

COPY ./package.json /client

RUN npm install --legacy-peer-deps
RUN echo $(ls)

CMD ["npm", "run", "dev"]
