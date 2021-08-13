FROM debian:10-slim

COPY slu /

ENTRYPOINT [ "/slu" ]
