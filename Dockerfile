FROM golang:1.18.4

# Create app directory
WORKDIR /application

# Install extensions
RUN apt-get update -y
RUN apt-get -y install gcc libgl1-mesa-dev xorg-dev gcc-mingw-w64