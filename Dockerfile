FROM debian:10-slim

COPY slut /

ENTRYPOINT [ "/slut" ]
