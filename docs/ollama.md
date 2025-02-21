# Ollama

Ollama is a tool that simplifies running and managing large language models locally on your computer.

## Contents

* [Installation and Setup](#installing-and-setup)
* [Managing Models](#managing-models)
* [Interacting with Models](#interacting-with-the-model)
* [Programming Techniques](#programming-techniques)
* [Working Examples](#working-examples)

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

## Configuring models (`modelfiles`)

A `modelfile` is a configuration file that tells Ollama how to use a specific language model. It is not the same as a model.

| **Feature** | **Llama Model (or other LLM)** | `modelfile` |
|---|---|---|
| **What it is?** | The trained neural network | A configuration file |
| **Content** |	Weights, biases, architecture of the model | Instructions, settings, system prompt |
| **Purpose** | To understand and generate text | To tell Ollama how to use the model |
| **Size** | Large (gigabytes) | Small (kilobytes) |
| **Example** | llama2-7b.ggmlv3.q4_0.bin (a model file) | my-custom-model.modelfile |

A `modelfile`  is used to configure the following:

1. **Persona and Context (via the System Prompt)**: This is the most significant form of customization.  The system prompt within the modelfile (enclosed in <<SYS>> and <</SYS>> tags) defines the model's persona, role, and context.  This is where you tell the model:

    * "You are a helpful and friendly travel agent."
    * "You are a sarcastic movie critic."
    * "You are a technical expert on quantum physics."

The system prompt shapes how the model responds to user input, its tone of voice, its knowledge domain, and its overall behavior.

2. **Generation Parameters**:  The modelfile allows you to adjust various parameters that control the text generation process.  These parameters influence things like:

    * `temperature`: Controls the "randomness" or "creativity" of the model's output. A higher temperature (e.g., 0.8) makes the output more diverse and unpredictable, while a lower temperature (e.g., 0.2) makes it more deterministic and focused.
    * `top_p (nucleus sampling)`: A way to control the diversity of the output. It sets a threshold for the cumulative probability of the most likely tokens.
    * `top_k`: Another way to control diversity. It limits the model to considering only the top K most likely tokens at each step.
    * `repeat_penalty`: Helps prevent the model from repeating the same words or phrases too often.

3. **Fine-tuning (Advanced)**: While not directly within the modelfile itself, the modelfile often plays a role in fine-tuning.  Fine-tuning involves training the model on a specific dataset to adapt it to a particular task or domain. The base or from parameter in the modelfile indicates the starting model for fine-tuning.  After fine-tuning, you would then use the modelfile to configure and run the fine-tuned model.

4. **Specific Instructions and Examples (via Prompts)**: While the system prompt sets the overall persona, individual prompts given to the model also contribute to customization.  You can provide specific instructions, examples, or constraints in your prompts to guide the model's behavior in each interaction.

Here is the [official reference](https://github.com/ollama/ollama/blob/main/docs/modelfile.md)

## Programming Techniques

* [API Documentation](https://github.com/ollama/ollama/blob/main/docs/api.md)

## Working examples

### Go and Llama

* Example 1
    * Text completion
    * Based on the official [Ollama API](https://github.com/ollama/ollama) to interact with the server.
    * Response with no streaming.
    * [Source](../examples/gollama/ex1/main.go).
* Example 2
    * Same as Example 1 with no streaming response.
    * [Source](../examples/gollama/ex2/main.go).
* Example 3
    * Using custom client [here](../internal/gollama/).
    * Request the server to analyse a picture of a [kitten](../testdata/cat.jpeg).
    * [Source](../examples/gollama/ex3/main.go).

### Using curl

The following examples uses `curl`.

* [Example 1](../examples/curl/ex1.sh) - As per Go example 1.
* [Example 2](../examples/curl/ex2.sh) - As per Go example 2.