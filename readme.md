# pulsar-demo
    
    pulsar go demo
    
# pulsar docker run

    docker run -itd \
      -p 6650:6650 \
      -p 8080:8080 \
      --mount source=pulsardata,target=/pulsar/data \
      --mount source=pulsarconf,target=/pulsar/conf \
      apachepulsar/pulsar:2.6.1 \
      bin/pulsar standalone

# docs
    
    http://pulsar.apache.org/docs/zh-CN/standalone/
    