FROM php:8.0-fpm-buster

LABEL "uk.hudson-viles.registry"="PHP PDO"
LABEL uk.hudson-viles.registry-image-base="php:8.0.25-fpm"
LABEL version="0.0.1"
LABEL description="This PHP Image is just a basic one, created off the back of the standard PHP Image. \
                   The only difference with it is it has the PDO extensions enabled to enable MySQL communication via a PHP project.\
                   This was designed to be used in Kubernetes where the docker-php-ext-install pdo pdo_mysql pdo_pgsql would not be able to be run without a intervention somewhere.\
                   This image is that intervention."

RUN apt-get update
RUN apt-get install -y libpq-dev
RUN docker-php-ext-configure pgsql -with-pgsql=/usr/local/pgsql

RUN docker-php-ext-install pdo pdo_mysql pgsql pdo_pgsql