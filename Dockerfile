FROM golang:1.6
# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/sgoertzen/pvi
RUN go get github.com/sgoertzen/pvi
RUN go install github.com/sgoertzen/pvi/cmd/pvweb
#ENTRYPOINT /go/bin/pvi
CMD ["/go/bin/pvweb", "/go"]
EXPOSE 80 443