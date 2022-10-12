# CGO-less Foreign Function Interface with WebAssembly

## Abstract

Foreign Function Interface(FFI) is a mechanism to invoke functions written in another language. In Go, there are only a few ways to achieve this, calling C functions via CGO, embedding Lua VM, etc. FFI is especially useful when we want to extend the existing Go projects by leveraging other programming languages’ ecosystems.

WebAssembly(Wasm) is a relatively-new instruction format for a virtual machine, and numerous programming languages can be compiled down to Wasm binary format. Therefore, if there were a Wasm virtual machine that is embeddable in Go projects, executing Wasm binary in Go could be a way to achieve FFI in a secure and flexible way since Wasm is designed to be secure, and any programming language can have support for it.

At Tetrate, we have open-sourced a WebAssembly runtime purely in Go called “wazero”. wazero is 100% compliant with the WebAssembly Core Specification 1.0 and 2.0. It has zero dependencies and doesn't rely on CGO. This means you can run applications in other languages and still keep cross-compilation, etc. The prime motivation behind wazero was to write a Wasm runtime purely in Go to avoid CGO usage since there was no pure Go Wasm runtime out there (except a few that are archived), therefore we were not able to achieve FFI via Wasm without CGO.

This tutorial will cover the basics of WebAssembly, the state of the Go ecosystem around it, how CGO-less FFI with Wasm looks like through a live demo, a deep dive into the wazero’s WebAssembly to machine code compiler and the virtual machine implementation, and how Gophers can start leveraging CGO-less FFI in their project.

## Speaker Details

- Name: Takeshi Yoneda
- Twitter: https://twitter.com/mathetake
- GitHub: https://github.com/mathetake
