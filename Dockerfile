# SPDX-FileCopyrightText: 2021-2022 Darren <1912544842@qq.com>
# SPDX-License-Identifier: Apache-2.0


FROM ubuntu:22.04

MAINTAINER dreamcmi

COPY build/emqx-user-manager /root/emqx-user-manager

COPY config.toml /root/config.toml

WORKDIR /root

EXPOSE 5555

CMD ./emqx-user-manager
