FROM mongo:4.0.6

COPY import-data.sh /import-data.sh
RUN chmod +x /import-data.sh

COPY titanic.csv /mongo-seed/titanic.csv

ENTRYPOINT ["/import-data.sh"]
