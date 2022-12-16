#!/usr/bin/env bash

cd magento2
service apache2 status
service apache2 start
service apache2 status

php bin/magento setup:install --base-url=${MAGENTO_HOST} --db-host=${MAGENTO_DATABASE_HOST} --db-name=${MAGENTO_DATABASE_NAME} --db-user=${MAGENTO_DATABASE_USER} --db-password=${MAGENTO_DATABASE_PASSWORD} --admin-firstname=Admin --admin-lastname=Admin --admin-email=admin@admin.com --admin-user=admin --admin-password=${MAGENTO_ADMIN_PASSWORD} --language=en_US --currency=USD --timezone=Asia/Kolkata --backend-frontname=admin --search-engine=elasticsearch7 --elasticsearch-host=${ELASTICSEARCH_HOST} --elasticsearch-port=${ELASTICSEARCH_PORT_NUMBER}

service apache2 start
service apache2 status

php bin/magento indexer:reindex && php bin/magento se:up && php bin/magento se:s:d -f && php bin/magento c:f && php bin/magento module:disable Magento_TwoFactorAuth

php bin/magento setup:upgrade && php bin/magento setup:static-content:deploy -f && php bin/magento cache:flush && chmod -R 0777 var/ pub/ generated/

service apache2 start
service apache2 status

tail -f /dev/null