# FatTree
[![Drone (cloud)](https://img.shields.io/drone/build/reinnet/topology.svg?style=flat-square)](https://cloud.drone.io/reinnet/topology)

## FatTree
K-ary fat tree:
three-layer topology (edge, aggregation and core):
  * Each pod consists of (k/2)^2 servers & 2 layers of k/2 k-port switches
  * Each edge switch connects to k/2 servers & k/2 aggr. switches
  * Each aggr. switch connects to k/2 edge & k/2 core switches
  * (k/2)^2 core switches that each one connects to k pods
## USNet
A topology with 24 nodes and 43 links.
