# Go Fiber Template

This is a template backend repository which abstracts away a bunch of the complexities Fiber brings to the table with the help of utility functions.
This repository also structures code in an opinionated way of modules (controllers + services) in place of having controllers and responses

### Some of the key utilities are as follows

1. Standard Response Formats - This repo expects all API responses to be in a particular format and enforces that with the help of util response handlers.
2. Agnostic Services - Services do not take in any Http contexts and are treated as pure functions. They can either `Panic` or return an `interface{}` which will be passed on as the http response. The goal of these services is to be reused within the code as much as possible, and hence not tie them up to http methods.
3. Auto Request Validations - Validating request body _AUTOMATICALLY_ with the help of the `validate` package, given that the DTO's are tagged appropriately. This is again achieved with the help of util request handlers. And you don't have to write any code to validate the request bodies.
4. Error Handling - The server will never crash, regardless of any runtime errors which may occur (as long as you don't call `log.Fatal("")` in your code). A 500 response will be sent back to the client gracefully.
5. An intuitive Controller-Service-Model approach, where controllers handle routing, services handle the data and models (schemas) provide the data. This ensures that your code is maintainable especially when your project grows large.

** Things to do **

1. Abstract away the complexities of quering lists from MongoDB
2. Make swagger doc implementation as easy as possible
3. Bind models to services, such that services control the data-flow

---

As you can see, this architecture is extermely influcenced by NestJS. I intend to make Fiber as easy as possible to develop on with my handly util packages
