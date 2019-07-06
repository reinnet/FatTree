# FatTree
[![Drone (cloud)](https://img.shields.io/drone/build/roadtomsc/FatTree.svg?style=flat-square)](https://cloud.drone.io/roadtomsc/FatTree)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/roadtomsc/FatTree)

## Definition
K-ary fat tree: three-layer topology (edge, aggregation and core):
* each pod consists of (k/2)^2 servers & 2 layers of k/2 k-port switches
* each edge switch connects to k/2 servers & k/2 aggr. switches
* each aggr. switch connects to k/2 edge & k/2 core switches
* (k/2)^2 core switches: each connects to k pods
