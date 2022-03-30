FROM registry.access.redhat.com/ubi8-minimal

EXPOSE 8442

RUN microdnf update -y && rm -rf /var/cache/yum && microdnf install git go make -y && microdnf clean all

COPY . /opt/jump-app-golang-events
WORKDIR /opt/jump-app-golang-events

RUN make build

CMD ["/opt/jump-app-golang-events/bin/back-golang-events"]