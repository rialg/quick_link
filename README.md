# Quick Link

## Overview

Forwarding proxy based in a short hostname. If the root is `myquicklink`, then every HTTP access can be
done in the following way:

```
http://myquicklin/some/uri
```

## Prerequisite

* Docker
* Tested in Ubuntu

## Install

Please, replace `your-quick-link-hostname` with the choosen hostname.

```
$ export QUICKLINK_ROOT="<your-quick-link-hostname>"
$ find . -iname *.go -or -iname *.bash | \
  sed -i "s/<QUICKLINK_ROOT>/${QUICKLINK_ROOT}/g"
$ sudo bash install.bash
```

## Remove

```
$ docker container stop quick_link
```
