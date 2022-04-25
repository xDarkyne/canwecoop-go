# -------------------- Frontend | Dependency Stage --------------------
FROM node:17-alpine AS deps
RUN apk add --no-cache libc6-compat

# prepare the project
WORKDIR /app
COPY ./client/package.json ./client/yarn.lock ./

# install the dependencies
RUN yarn install --frozen-lockfile

# -------------------- Frontend | Build Stage --------------------
FROM node:17-alpine AS builder

WORKDIR /app

ENV NEXT_TELEMETRY_DISABLED 1

# collect compiled depencies and project-files
COPY --from=deps /app/node_modules ./node_modules
COPY client/. .

# build the project
RUN yarn export
# -------------------- Backend | Build Stage --------------------
FROM golang:1.18-alpine AS go

WORKDIR /app
COPY . .

RUN go build .
# -------------------- Deploy Stage --------------------
FROM golang:1.18-alpine

WORKDIR /app
COPY --from=go /app/steamgo ./steamgo
COPY --from=builder /app/out/ ./public

CMD ["./steamgo"]
