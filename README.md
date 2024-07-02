# Voice notes

Full-stack web application for recording voice notes and transcribing them to text using OpenAI Whisper API. The backend is written using Golang and Gin framework. Frontend is written using Next.js and TailwindCSS.

## Start the application

Before starting the server make sure to set up the following environment variables:

```
PORT - Server port
OPENAI_AUTH_TOKEN - OpenAI API authentication token
```

Start the server:

```
go run .
```

Run the server tests:

```
go test -tags mock ./tests
```

Go to the `client` folder for the instructions to run the frontend part of the app
