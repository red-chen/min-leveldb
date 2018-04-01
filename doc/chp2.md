# 从零开始构建LevelDB


### 配置文件

```

language: go

go:
  - 1.5
  - 1.6
  - 1.7
  - 1.8
  - 1.9
  - tip

script:
  - go test -timeout 1h ./...

```

### 在Readme中配置Travis-CI的状态图标