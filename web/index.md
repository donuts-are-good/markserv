# Markus Document Network

The Markus document network is a minimalistic system designed for sharing and viewing documents written in Markdown and plaintext. It aims to make the process of creating and maintaining a web of documents as easy and straightforward as possible.

## Purpose

The Markus document network provides a basic approach to sharing content. By using markdown as the basis for content creation, non-expert users can create rich documents without needing to worry about complex HTML or styling. The goal is to encourage users to focus on the content, not the tool. It's an ideal platform for knowledge sharing, blogging, documentation, and much more.

## How it Works

The Markus document network relies on two main components: a server and a client. The server, known as MarkServ, serves Markdown files over HTTP, while the client, known as MarkClient, retrieves and displays these documents.

## MarkServ

MarkServ is a simple HTTP server that reads and serves Markdown files from a specified directory. It is designed to run on port 88 and serves files based on their paths. If a path ends with a "/", it will look for an "index.md" file. Otherwise, it will attempt to serve the file specified by the path, assuming it has a valid file extension (defined in the configuration).

MarkServ also handles basic security measures, such as checking for path traversal attacks and serving only allowed file types.

## MarkClient

MarkClient is a terminal-based client designed to fetch and display Markdown documents served by MarkServ or any compatible server. The client displays Markdown with preserved text formatting using ANSI escape codes. It also includes an offline mode, saving viewed documents for offline reading.

When the client retrieves a document with links, these links are numbered and listed at the bottom of the document. Users can follow these links by inputting the corresponding number.
Compatibility and Extensibility

The Markus document network is compatible with any server that serves Markdown documents over HTTP. Moreover, while MarkClient is a dedicated client for the network, Markdown files can also be viewed in any web browser with a compatible Markdown viewer plugin, like [Markdown Viewer Webext](https://addons.mozilla.org/en-US/firefox/addon/markdown-viewer-webext/).

## Conclusion

In essence, the Markus document network is an hommage to simplicity and user-focused design. It aims to facilitate lo-fi sharing and viewing of documents, enabling users to concentrate on the content that matters.

For further details or to contribute, visit the GitHub repositories:

- [MarkServ](https://github.com/donuts-are-good/markserv)
- [MarkClient](https://github.com/donuts-are-good/markclient)

Enjoy the journey through the Markus document network!
