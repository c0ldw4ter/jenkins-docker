FROM nginx:mainline-alpine3.19


COPY nginx/www/trainee-22.zapto.org/ /var/www/trainee-22.zapto.org/
COPY nginx/www/trainee-22.zapto.org/conf/nginx.conf /etc/nginx/nginx.conf
COPY nginx/www/trainee-22.zapto.org/conf/trainee-22.conf /etc/nginx/conf.d/default.conf


RUN apk update && \
    adduser -D -g "www" www && \
    mkdir -p /var/cache/nginx && \
    chown -R www:www /var/cache/nginx && \
    chown -R www:www /var/log/ && \
		chown -R www:www /etc/nginx/ &&	\
    chown -R www:www /var/run/


USER www

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
