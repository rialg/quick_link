FROM ubuntu:focal

ENV TZ=Europe/Copenhagen
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt update -y
RUN apt install -y wget

RUN wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz
RUN tar -C /usr/local/ -zxvf go1.19.2.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin 
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/src  

ADD ./ /home/ubuntu/
ENV PROXYPORT="80"
RUN cd /home/ubuntu/src/ && \ 
    sed -i "s/<PROXYPORT>/${PROXYPORT}/g" /home/ubuntu/src/main.go && \
    go build -o proxy.bin

WORKDIR "/home/ubuntu/src"

CMD ["/home/ubuntu/src/proxy.bin"]
