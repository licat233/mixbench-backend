#!/bin/sh
rm -f /etc/prometheus/prometheus.default.yml
mv /etc/prometheus/prometheus.yml /etc/prometheus/prometheus.default.yml
cp /prometheus/config/template/prometheus.yml /etc/prometheus/prometheus.yml
for i in $(env); do
    if [ $(echo $i | egrep "^PROMETHEUS_") ]; then
        KEY=$(echo "$i" | sed "s/^PROMETHEUS_\([A-Z]*\)\=\(.*$\)/\1/")
        VALUE=$(echo "$i" | sed "s/^PROMETHEUS_\([A-Z]*\)\=\(.*$\)/\2/")
        sed -i "s/%%${KEY}%%/${VALUE}/g" /etc/prometheus/prometheus.yml
    fi
done

