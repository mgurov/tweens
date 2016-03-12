FROM golang:latest

RUN apt-get update

RUN apt-get install -y libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libgl1-mesa-dev

RUN go get github.com/go-gl/gl/v2.1/gl github.com/go-gl/glfw/v3.1/glfw

ADD . $GOPATH/src/github.com/mgurov/tweens

WORKDIR $GOPATH/src/github.com/mgurov/tweens
