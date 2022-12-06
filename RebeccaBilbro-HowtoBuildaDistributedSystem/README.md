# How to Build a Distributed System

## Abstract

This talk tells the story of how we used Go to build our very own, eventually consistent, distributed system, currently deployed in production clusters across the US, Germany, and Singapore. As our system is described, key topics will be introduced that you’ll need, to understand distributed systems in practice (e.g. replication, consistency, and consensus). We will walk through how our team leveraged tools like gRPC, Kubernetes, LevelDB, and Prometheus to implement two new open source projects that serve as the heart of our system. Confessions of all the ways we messed up will also ensue — from struggling to debug protocol buffer errors, to tangling up send and receive goroutines, to reasoning about the phases of replication. Finally, Rebecca will explain why rolling out their own system made sense for their use case, and why it might also make sense for you. It won’t be the prettiest story, but we hope you’ll benefit from the lessons we learned, including the most important one — that you can build your own distributed system.


## Slides 

A PDF copy of the slides could be found [here](./slides.pdf)
