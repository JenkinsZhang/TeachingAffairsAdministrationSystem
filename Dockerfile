FROM node:10.4.1 as builder
ADD . /
RUN yarn config set registry https://registry.npm.taobao.org
RUN yarn
RUN yarn run build

FROM keymetrics/pm2:10-alpine
COPY --from=builder / .
ENTRYPOINT ["pm2-runtime","start","npm","--name","TeachingAffairsAdministrationSystem","--","run","start"]
