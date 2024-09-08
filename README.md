<div align="center">
  <img src="/docs/docs/assets/img/lilbox.svg" height="200"/>
</div>

<h1 align="center" style="margin-top: -10px"> HomeBox </h1>
<p align="center" style="width: 100;">
   <a href="https://hay-kot.github.io/homebox/">Docs</a>
   |
   <a href="https://homebox.fly.dev">Demo</a>
   |
   <a href="https://discord.gg/tuncmNrE4z">Discord</a>
   |
   <a href="https://github.com/hfhc1309/homebox-zh">汉化：hfhc1309</a>
</p>

## 汉化说明
  
  [HomeBox中文汉化版](https://github.com/hfhc1309/homebox-zh)
  版本：v0.10.3 ~ 构建： e8449b3a7363a6cfd5bc6151609e6d2d94b4f7d8
  汉化时间：2024/9/9

  [原版连接](https://github.com/hay-kot/homebox)

## Quick Start

[Configuration & Docker Compose](https://hay-kot.github.io/homebox/quick-start)

```bash
docker pull docker.io/hfhc1309/homebox-zh:latest
# If using the rootless image, ensure data 
# folder has correct permissions
# /path/to/data/folder/记得修改为对应路径
mkdir -p /path/to/data/folder
chown 65532:65532 -R /path/to/data/folder
docker run -d \
  --name homebox \
  --restart unless-stopped \
  --publish 3100:7745 \
  --env TZ=Europe/Bucharest \
  --volume /path/to/data/folder/:/data \ 
  docker.io/hfhc1309/homebox-zh:latest
# 汉化镜像 docker.io/hfhc1309/homebox-zh:latest
# 原版镜像 ghcr.io/hay-kot/homebox:latest-rootless

```

## Credits
- Source by [@hay-ko](https://github.com/hay-kot)
- Logo by [@lakotelman](https://github.com/lakotelman)
