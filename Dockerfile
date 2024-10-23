FROM golang:1.23.2-alpine3.20 AS base
RUN apk add --no-cache tzdata

################ BUILD ################
FROM base AS build
WORKDIR /app

# install dependencies
COPY go.* ./
RUN go mod download

# copy source code
COPY . ./

# build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

################ PRODUCTION  ################
FROM base AS production
WORKDIR /app

# copy the binary from the build stage
COPY --from=build /app/main ./
COPY --from=build /app/env.yaml.example ./env.yaml
COPY --from=build /app/docker/entrypoint.sh ./
COPY --from=build /app/static ./static

RUN chmod +x ./entrypoint.sh

# set timezone
ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 5001
# run app
ENTRYPOINT [ "/app/entrypoint.sh" ]