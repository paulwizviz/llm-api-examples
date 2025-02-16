# Ollama

Ollama is a tool that simplifies running and managing large language models locally on your computer.

## Installing and Setup

* MacOS 
    * Homebrew `brew install ollama`.
        * This does not install a macOS Application
        * You will need to start ollama's controller manually by running the command `ollama serve` before you can run a model.
    * Other approach refer to [Github](https://github.com/ollama/ollama).

## Managing models

On macOS and Linux, LLM models are stored in:

```
~/.ollama
```

* Pull models run `ollama pull <model name>`
* List models run `ollama list`
* Remove model run `ollama rm llama2:latest`

Here are a list of [models](https://ollama.com/library).

## Running Models

* Run the command `ollama run <name of model> <prompt>`
* Run the command `ollama chat <name of model>` to start a REPL