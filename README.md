# transfer

transfer是一个使用二维码在局域网内传输文件的工具，支持上传和下载文件。它使用[cobra][1]构建，并在打包时，使用[go-bindata][2]将表单模板转换为go文件，然后生成可执行文件。

## Usage

1. `git clone git@github.com:One1ight/filetransfer.git`

2. `cd filetransfer && go build`

3. `./transfer [flag]`

## Flags

- `upload`  文件上传

- `download xxx.xx` 指定文件名下载

## Process

- 上传文件：开启http服务，输出二维码。手机扫码获取链接，浏览器打开链接，选择文件上传。服务端文件接受完毕后服务自动关闭，程序结束

- 下载文件：开启http服务，输出二维码。手机扫码获取链接，浏览器打开链接，选择文件夹下载文件。完成后手动ctrl+c关闭服务，程序结束

## Build

1. `go-bindata -pkg handler -o handler/tmpl.go web/template`

2. `go build`

## 参考

[qrcp][3]

<http://mrwaggel.be/post/golang-transmit-files-over-a-nethttp-server-to-clients/>

<https://github.com/mdp/qrterminal>

<https://github.com/plutov/go-bindata-tpl>

[1]: https://github.com/spf13/cobra
[2]: https://github.com/jteeuwen/go-bindata
[3]: https://github.com/claudiodangelis/qrcp
