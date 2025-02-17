# Ollama

Ollama is a tool that simplifies running and managing large language models locally on your computer.

## Installing and Setup

### MacOS 

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

## Interacting with the Model

There are two ways of interacting with Llama models:

* Test completion mode - `ollama run <name of model> <prompt>`.
* Conversational mode - `ollama chat <name of model>`.

| Feature | Chat | Text Completion |
|---|---|---|
| **Purpose** |	Conversational interactions | General text generation and manipulation |
| **Focus** | Dialogue, context, persona | Prompt-based generation, flexibility |
| **Control** |	Less direct control over formatting |	More direct control over input and output |
| **Output** | Natural, conversational text | Can be adapted to various formats |

### Chat personas

Llama chat is able to adopt different personas. You can use personas in multiple ways. One typical use case is a chatbot. Here is an example of the how it can be used.

Personas:

* **User (Human)**: This is the person interacting with your chatbot. They type messages, ask questions, and generally drive the conversation from their side.

* **Assistant (Your Chatbot's "Persona")**:  This is the role that your chatbot plays in the conversation.  It's defined by the system prompt you provide to the Llama 2 model.  It's not a separate entity or intermediary.  Think of it as the voice and personality you give to your chatbot.  It's how your chatbot presents itself to the user.

* **Llama 2 Model (The Engine)**: This is the large language model that powers your chatbot.  It's the "brain" behind the operation.  It receives the conversation history (including the system prompt and the user's messages) and generates the assistant's responses.

How it works together:

1. **User Input**: The human user types a message into your chatbot interface.

2. **Your Application**: Your chatbot application takes the user's message and formats it into the JSON structure we discussed earlier.  It adds the user's message as a new object in the messages array, with the role set to "user."  It also includes the original system prompt that defines the assistant's persona.

3. **API Request**: Your application sends this JSON payload to the Llama 2 Chat API.

4. **Llama 2 Processing**: The Llama 2 model receives the JSON. It uses the system prompt to understand the desired persona and the conversation history to generate a relevant and appropriate response.  The model acts as the assistant.

5. **API Response**: The API returns a JSON payload containing the model's generated response, with the role set to "assistant."

6. **Your Application**: Your chatbot application receives the JSON response. It extracts the assistant's message and displays it to the user in the chatbot interface.

7. **Repeat**: The process repeats as the user continues the conversation.

**Key takeaway**: The assistant isn't an intermediary; it's the role the Llama 2 model plays when generating responses.  Your chatbot application is the intermediary between the user and the Llama 2 model, formatting messages and managing the conversation flow.  The Llama 2 model becomes the assistant within the context of each API call.

## Programming Techniques

* [API Documentation](https://github.com/ollama/ollama/blob/main/docs/api.md)

## Working examples

### Go and Llama
    
* [Example 1](../examples/gollama/ex1/main.go) - This example demonstrates a text completion interaction with Llama via RESTFul API.

### Using curl

* [Example 1](../examples/curl/ex1.sh) - This example demonstrates a text completion scenario using `curl` to interact with Llama via RESTFul API