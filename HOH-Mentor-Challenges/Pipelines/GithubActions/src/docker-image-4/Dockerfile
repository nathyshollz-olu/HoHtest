FROM alpine:latest AS builder
RUN apk --no-cache add build-base

FROM builder AS build1
COPY src/source1.cpp source.cpp
RUN g++ -o /binary source.cpp

FROM builder AS build2
COPY src/source2.cpp source.cpp
RUN g++ -o /binary source.cpp
