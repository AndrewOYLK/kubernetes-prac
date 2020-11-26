export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=192.168.189.128:23791
export MICRO_API_NAMESPACE=api.jtthink.com # 命名空间
micro web # 默认8082


# docker run \
#   -p 2379:2379 \
#   -p 2380:2380 \
#   --name etcd -d quay.io/coreos/etcd:latest \
#   /usr/local/bin/etcd \
#   --data-dir=/etcd-data --name node1 \
#   --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://${NODE1}:2380 \
#   --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://${NODE1}:2379 \
#   --initial-cluster node1=http://${NODE1}:2380