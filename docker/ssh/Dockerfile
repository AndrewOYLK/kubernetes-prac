FROM centos:latest

RUN yum -y update && \
  yum -y install net-tools passwd cracklib-dicts && \
  useradd mokalo && \
  echo "123qwe"|passwd --stdin mokalo && \
  yum provides sshd && \
  yum -y install openssh

CMD /usr/sbin/sshd -D &

