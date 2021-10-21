FROM golang:alpine AS build

# 设置环境变量
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

# 项目代码存放的目录
WORKDIR /home/www/lin-cms-go

ADD go.mod .
ADD go.sum .

COPY . .

# 编译成二进制可执行文件名 app
RUN go build -o app .


FROM alpine AS prod

COPY --from=build /home/www/lin-cms-go/ .
COPY --from=build /home/www/lin-cms-go/configs/config.yaml ./configs/

# 端口
EXPOSE 3000

CMD ["./app"]


# docker build -t lin-cms:v1 -f Dockerfile .
# docker run -it -p 8080:3000 lin-cms:v1