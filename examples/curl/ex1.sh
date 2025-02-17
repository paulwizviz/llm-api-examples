#!/bin/bash

curl -X POST \
     -H "Content-Type: application/json" \
     -d '{
           "model": "llama3.2:1b", 
           "prompt": "What is life?",
           "stream": false
         }' \
     http://localhost:11434/api/v1/generate