FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ Asia/Kolkata

WORKDIR /var/www/html

ADD ./install-magento.sh .

RUN apt-get update && apt-get -y upgrade \
	&& apt-get install -y --no-install-recommends ca-certificates openssl curl unzip apache2 php7.4 libapache2-mod-php php-mysql php7.4-mbstring php7.4-bcmath php7.4-intl php7.4-soap php7.4-zip php7.4-gd php7.4-json php7.4-curl php7.4-cli php7.4-xml php7.4-xmlrpc php7.4-gmp php7.4-common

RUN phpenmod mbstring \
    && a2enmod rewrite \
    && sed -i '/index.php/c\    DirectoryIndex index.php index.html index.cgi index.pl index.xhtml index.htm' /etc/apache2/mods-enabled/dir.conf \
    && sed -i '/max_execution_time/c\max_execution_time = 18000' /etc/php/7.4/cli/php.ini \
    && sed -i '/max_input_time =/c\max_input_time = 1800' /etc/php/7.4/cli/php.ini \
    && sed -i '/memory_limit/c\memory_limit = 4G' /etc/php/7.4/cli/php.ini \
    && service apache2 restart \
    && curl -sS https://getcomposer.org/installer -o composer-setup.php \
    && php composer-setup.php --install-dir=/usr/bin --filename=composer \
    && composer create-project --repository=https://4d489df14897491f8c5c4e53c9e76b83:b5398b66651f8fde0dca63f656ad57a8@repo.magento.com/ magento/project-community-edition=2.4.4 magento2 \
    && cd magento2 \
    && find var generated vendor pub/static pub/media app/etc -type f -exec chmod g+w {} + \
    && find var generated vendor pub/static pub/media app/etc -type d -exec chmod g+ws {} + \
    && chown -R root:www-data . \
    && chmod u+x bin/magento \
	&& rm -rf /var/lib/apt/lists/*

ADD ./apache2.conf /etc/apache2/sites-available/000-default.conf

ENTRYPOINT [ "/var/www/html/install-magento.sh" ]