# FatTree
[![Drone (cloud)](https://img.shields.io/drone/build/reinnet/FatTree.svg?style=flat-square)](https://cloud.drone.io/reinnet/FatTree)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/roadtomsc/FatTree)

## Definition
K-ary fat tree:
three-layer topology (edge, aggregation and core):
  * Each pod consists of (k/2)^2 servers & 2 layers of k/2 k-port switches
  * Each edge switch connects to k/2 servers & k/2 aggr. switches
  * Each aggr. switch connects to k/2 edge & k/2 core switches
  * (k/2)^2 core switches that each one connects to k pods
