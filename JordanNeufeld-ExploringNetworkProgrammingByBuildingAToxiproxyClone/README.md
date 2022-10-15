# Exploring Network Programming By Building A Toxiproxy Clone

[Jordan Neufeld](https://www.linkedin.com/in/jordan-neufeld/)

## Abstract

This tutorial is for Gophers who have written a command line or an API application, but have little to no experience in lower-level concepts like reading and writing to sockets, working with channels, and managing multiple goroutines. We’ll dig into these network programming concepts by exploring the architecture of the popular open source chaos testing tool Toxiproxy. Toxiproxy is a tool written in golang that Shopify uses to intercept, manipulate, and disrupt TCP streams in a controlled manner in order to test the resiliency of applications and analyze their failure modes. We’ll talk about why Go is a great choice for networking tools, and discuss what aspects of the language make it especially easy to be productive writing tools such as this. Finally, Jordan will demonstrate how simple it is to get started in network programming with a live demo where he’ll build a bare-minimum clone of Toxiproxy that can intercept and add latency to TCP streams between a client and server.

## Speaker

### Jordan Neufeld, Senior Site Reliability Engineer - Shopify

Jordan Neufeld is a Senior Site Reliability Engineer with over 6 years' experience supporting a wide range of technologies. He has led many incident response efforts through all aspects of triage, troubleshooting, remediation, and retrospectives. When Jordan is not debugging production, collaborating with peers, or creating a morale-boosting meme, you might find him hacking on Go code or Kubernetes controllers. In his spare time, Jordan enjoys making music, BBQing, and hanging out with his wife and children.