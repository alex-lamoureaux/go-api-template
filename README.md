# go-api-template
basic api template project

## How to build

### Local
Building this project locally requires you to have go installed
In a terminal, navigate to the project directory and run the following command
`go build -o go-api-template`
for windows, make sure the output artifact ends with the `.exe` extension
`go build -o go-api-template.exe`

### Docker
To build with docker, you must have docker installed, and the docker engine must be running
In a terminal, navigate to the project directory and run the following command
`docker build . -t go-api-template`

## How to run
Assuming you have built the project with one of the above methods, you will either have a binary executable or a docker image that can be run
### Local (binary)
In a terminal, navigate to the directory with your stored artifact, and run the following command
`./go-api-template`
or for windows
`./go-api-template.exe`

### Docker
In a terminal, run the following command
`docker run -p 8080:8080 go-api-template`
