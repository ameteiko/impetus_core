# A Docker image for a diagrams processing using MermaidJS-CLI.
#
# Build a local image:
#    docker build -f mermaidJS.Dockerfile -t mermaidcli:latest .
# Usage:
#    docker container run --rm  -it --entrypoint /bin/sh mermaidcli                         # For debugging purposes.
#    docker container run --rm  --mount type=bind,src=$( pwd ),dst=/data mermaidcli         # To convert all .mmd files
#                                                                                             in mounted directory.
#    docker container run --rm  --mount type=bind,src=$( pwd ),dst=/data mermaidcli path    # To convert all .mmd files
#                                                                                             in path.

FROM node:14.4.0-buster

ENV DATADIR=/data
ENV PATH="${PATH}:/node_modules/.bin"

# Install the MermaidJS.
RUN yarn add mermaid.cli@0.5.1

# install all svg-rendering dependencies.
RUN apt-get update && \
    apt-get install -y gconf-service libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 libcairo2 libcups2 libdbus-1-3 \
                       libexpat1 libfontconfig1 libgcc1 libgconf-2-4 libgdk-pixbuf2.0-0 libglib2.0-0 libgtk-3-0 \
                       libnspr4 libpango-1.0-0 libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 \
                       libxcomposite1 libxcursor1 libxdamage1 libxss1 libxtst6 libappindicator1 libnss3 libasound2 \
                       libatk1.0-0 libc6 ca-certificates fonts-liberation lsb-release xdg-utils wget

# Copy the conversion entrypoint.
COPY convert.sh /convert.sh

# Create a conversion config file.
RUN echo { \"args\": [\"--no-sandbox\"] } > puppeteer-config.json

WORKDIR $DATADIR

ENTRYPOINT ["/convert.sh"]