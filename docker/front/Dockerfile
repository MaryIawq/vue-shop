FROM node:14

WORKDIR /app

EXPOSE 5173 5173

RUN npm install && \
    npm run build && \
    npm cache clean --force

CMD ["npm", "start"]
