# pulsar-demo
    
    pulsar go demo
    
# pulsar docker run

    docker run -dit \
        --name pulsar-sever \
        -p 6650:6650 \
        -p 8080:8080 \
        --mount source=pulsardata,target=/pulsar/data \
        --mount source=pulsarconf,target=/pulsar/conf \
        apachepulsar/pulsar:2.7.0 \
        bin/pulsar standalone

# topic info

    http://192.168.0.11:8080/admin/v2/persistent/public/default/my-topic/stats

# go client sdk

    https://pulsar.apache.org/docs/zh-CN/client-libraries-go/

# nodejs client

    https://pulsar.apache.org/docs/zh-CN/client-libraries-node/

# docs
    
    http://pulsar.apache.org/docs/zh-CN/standalone/
