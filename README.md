# Quick Link

## Overview

Forwarding proxy based in a short hostname. If the root is `myquicklink`, then every HTTP access can be
done in the following way:

```
http://myquicklink/some/uri
```

## Prerequisite

* Docker
* Unix based machine (i.e. Linux)

## Install

Please, replace `your-quick-link-hostname` with the choosen hostname.

```
$ export QUICKLINK_ROOT="myquicklink" # <- your quick_link hostname
$ for file in $(find . -iname "*.go" -or -iname "*.bash"); do \
  sed -i "s/<QUICKLINK_ROOT>/${QUICKLINK_ROOT}/g" ${file}; \
  done
$ sudo bash install.bash
```

## Remove

```
$ docker container stop quick_link
```
