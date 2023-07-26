![markserv](https://github.com/donuts-are-good/git-gone/assets/96031819/09c359b5-ae83-420b-b47e-828923c91bed)
![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

# MarkServ

MarkServ is a lightweight, Markdown-focused server for the Markus document network. It's designed to be simple, performant, and compatible with any text files that use the Markdown syntax.

## Installation

To install MarkServ, clone the repository and build the project:

```shell
git clone https://github.com/donuts-are-good/markserv.git
cd markserv
go build
```
## Usage

MarkServ is configured through a `config.json` file, which should be located in the same directory as the executable. The configuration file should specify the server port, the location of the Markdown files, and the allowed file types. Here is an example configuration:

```json
{
  "port": "88",
  "web": "/path/to/your/markdown/files",
  "allowedFileTypes": [".md", ".txt", ".markdown"]
}
```
Once the server is configured, it can be started with the following command:

```shell
./markserv
```
The server will then be accessible at http://localhost:88 (or whichever port was specified in the configuration).
## Routing

The server works by responding to `HTTP GET` requests and returning the content of the corresponding Markdown file. The Markdown files should be organized in a directory structure that matches the URL structure. For example, a request to http://localhost:88/myfile.md would correspond to a file located at `/path/to/your/markdown/files/myfile.md`.

When a client requests a directory (like http://localhost:88/mydir/), the server will look for a file named index.md in that directory.

The server will respond with a 404 status code if the requested file does not exist, and a 400 status code if the request path contains an invalid file type or an attempt at a directory traversal attack.

## Security

MarkServ is designed with basic security measures to prevent unauthorized access to the server's files. It checks for and denies requests containing `".."`, which could potentially be used for a directory traversal attack. It also restricts the types of files that can be served to those specified in the configuration.

However, as with any server software, it is advisable to run MarkServ in a secured environment and to keep it up-to-date with the latest security patches.
## Contribution

If you find a bug, have a feature request, or want to contribute to the code, please feel free to open an issue or a pull request.

## license

MIT License 2023 donuts-are-good, for more info see license.md
