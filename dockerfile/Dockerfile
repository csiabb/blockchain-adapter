FROM alpine:3.11

ENV BLOCKCHAIN_ADAPTER_CFG_PATH=/etc/csiabb BLOCKCHAIN_ADAPTER_LOG_PATH=/var/log/csiabb BLOCKCHAIN_ADAPTER_DATA=/opt/csiabb/data
RUN mkdir -p ${BLOCKCHAIN_ADAPTER_CFG_PATH} ${BLOCKCHAIN_ADAPTER_LOG_PATH} ${BLOCKCHAIN_ADAPTER_DATA}

COPY build/bin/blockchain-adapter /usr/local/bin
COPY sampleconfig/blockchain-adapter.yaml ${BLOCKCHAIN_ADAPTER_CFG_PATH}/

EXPOSE 8989

CMD ["blockchain-adapter"]
