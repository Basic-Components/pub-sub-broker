FROM hsz1273327/pub-sub-broker:Base-v0
ADD . /app/src/github.com/Basic-Components/pub-sub-broker
ENV GOPATH="/app"
WORKDIR /app/src/github.com/Basic-Components/pub-sub-broker
RUN go build