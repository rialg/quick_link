# Quick Link

## Overview

Forwarding proxy based in a short hostname. If the root is `myquicklink`, then every HTTP access can be
done in the following way:

```
http://myquicklin/some/uri
```

## Install

Please, replace `your-quick-link-hostname` with the choosen hostname.

```
$ export QUICKLINK_ROOT="<your-quick-link-hostname>"
$ sed -i "s/<QUICKLINK_ROOT>/${QUICKLINK_ROOT}/g" install.bash
$ sudo bash install.bash
```

## Remove

```
$ docker container stop quick_link
```
