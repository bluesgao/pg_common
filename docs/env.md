## portainer

docker volume create portainer_data

docker run -d -p 8000:8000 -p 9443:9443 --name portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce:lts

## **rocketMq环境搭建**


docker network create mynet

-- rocketmq-nameserver
docker run -d \
--restart=always \
--name rmqnamesrv \
-e "MAX_HEAP_SIZE=256M" \
-e "HEAP_NEWSIZE=128M" \
-p 9876:9876 \
-t apache/rocketmq:5.3.2 sh mqnamesrv

-- rocketmq-broker
docker run -d \
--restart=always \
--name rmqbroker \
-p 10912:10912 -p 10911:10911 -p 10909:10909 \
-p 8080:8080 -p 8081:8081 \
-e "NAMESRV_ADDR=43.134.114.153:9876" \
-t apache/rocketmq:5.3.2 sh mqbroker --enable-proxy \
-c /home/rocketmq/rocketmq-5.3.2/conf/broker.conf


docker run -d \
--restart=always --name rmqbroker --privileged=true \
-p 10912:10912 -p 10911:10911 -p 10909:10909 \
-p 8080:8080 -p 8081:8081 \
-v /data/rocketmq/broker/logs:/root/logs \
-v /data/rocketmq/broker/store:/root/store \
-v /data/rocketmq/broker/conf/broker.conf:/home/rocketmq/rocketmq-5.3.2/conf/broker.conf \
-e "NAMESRV_ADDR=43.134.114.153:9876" \
-t apache/rocketmq:5.3.2 sh mqbroker --enable-proxy \
-c /home/rocketmq/rocketmq-5.3.2/conf/broker.conf




-- rocketmq-dashboard
docker run -d --name rmqdashboard \
--restart=always \
-e "JAVA_OPTS=-Xmx256M -Xms256M -Xmn128M -Drocketmq.namesrv.addr=43.134.114.153:9876" \
-p 18080:8080 \
-t apacherocketmq/rocketmq-dashboard:latest

docker run -d \
--name rmqbroker \
--network mynet \
-p 10912:10912 -p 10911:10911 -p 10909:10909 \
-p 8080:8080 -p 8081:8081 \
-e "NAMESRV_ADDR=rmqnamesrv:9876" \
-v /Users/bluesgao/env/rmq/conf/broker.conf:/home/rocketmq/rocketmq-5.3.2/conf/broker.conf \
-t apache/rocketmq:5.3.2 sh mqbroker --enable-proxy \
-c /home/rocketmq/rocketmq-5.3.2/conf/broker.conf


docker run -d \
--name rmqbroker \
-p 10912:10912 -p 10911:10911 -p 10909:10909 \
-p 8080:8080 -p 8081:8081 \
-e "NAMESRV_ADDR=rmqnamesrv:9876" \
-t apache/rocketmq:5.3.2 sh mqbroker --enable-proxy \
-c /home/rocketmq/rocketmq-5.3.2/conf/broker.conf


docker run -p 3306:3306 --name mysql --restart=always \
 -e MYSQL_ROOT_PASSWORD=Dev123456\
 -d mysql:8

docker run -p 6379:6379 --name redis --restart=always \
-d redis:8 \
redis-server --requirepass "Dev123456"
