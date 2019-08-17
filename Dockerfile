FROM centos:7
COPY server /root/server
EXPOSE 50001
CMD /root/server
