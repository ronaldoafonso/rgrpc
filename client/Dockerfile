FROM python:3.8.2-buster

RUN adduser productinfoclient

WORKDIR /home/productinfoclient

USER productinfoclient:productinfoclient

COPY --chown=productinfoclient:productinfoclient . /home/productinfoclient

RUN pip install grpcio-tools
