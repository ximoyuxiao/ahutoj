FROM nginx

COPY --link ./docs/nginx.conf /etc/nginx/conf.d/default.conf
COPY --link ./tmp/nginx/html/dist /usr/share/nginx/html