# svnall
svn all: 对所有svn仓库进行操作的小工具

## 安装
方式1、通过brew安装
```shell
brew tap ns-cn/ttools && brew install svnall
```
方式2、手动安装，下载地址如下
```
https://github.com/ns-cn/ttools/releases/
```

## 使用实例
#### SVNALL_REPOSITORIES
环境变量，全局配置仓库，可在执行命令时自动读取， 仓库的地址配置格式：```{仓库路径}[#{查询深度}]```，

其中查询深度可不指定，默认值为2，多个使用冒号[```:```]分割，

例如：
```
export SVNALL_REPOSITORIES=~/workspace:~/codes#3
```

#### SVNALL_DEPTH
环境变量，数字类型，全局配置查询深度，可再执行命令时自动读取；

仓库执行深度优先级：仓库自行指定>命令行depth参数>环境变量

例如：
```
# 全局配置遍历深度为3
export SVNALL_DEPTH=3
```
#### svnall使用实例

```shell
# 查看程序版本号
> svnall version

# 手动更新环境变量全局配置的svn仓库
> svnall update 
> svnall update -d 3

# 手动更新指定目录
> svnall update ~/workspace                       # 同时执行workspace目录的svn update以及环境变量SVNALL_REPOSITORIES指定值
> svnall update ~/workspace ~/workbeach           # 同时执行workspace目录的svn update以及环境变量SVNALL_REPOSITORIES指定值
> svnall update -d 3 ~/workspace ~/workbeach      # 向下遍历三层查找svn仓库
> svnall update -d 3 ~/workspace#2 ~/workbeach    # 通用遍历三层，但workspace只向下遍历2层
> svnall update -d 3 -e ~/workspace#2 ~/workbeach # 仅执行workspace和workbeach目录的svn update，忽略环境变量中的配置
```