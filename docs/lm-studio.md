# LM Studio

LM Studio is a desktop application to enable users to run LLMs locally.

## Contents

## Installation and Setup

### macOS

* Homebrew `brew install -cask lm-studio`
* Download application from [LM Studio](https://lmstudio.ai/)

## LM Studio Operations

1. Start/Stop model server

![img](../assets/img/lm-studio-server.png)

## Interactring with the AI Model

This is the REST API to interact with AI models

* [Version 1](https://lmstudio.ai/docs/app/api/endpoints/openai)

## Working Examples

### Go and LM Studio interaction

* Example 1
    * Using [REST API client implementation](../internal/golm)
    * List models
    * [Source](../examples/lm-studio/ex1/main.go)
* Example 2
    * Using [REST API client implementation](../internal/golm)
    * Using the endpoing /v1/completion with no streaming
    * [Source](../examples/lm-studio/ex2/main.go)
* Example 3
    * Using [REST API client implementation](../internal/golm)
    * Using the endpoing /v1/completion with streaming
    * [Source](../examples/lm-studio/ex3/main.go)