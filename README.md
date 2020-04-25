filetransfer是一个是用局域网扫码传输文件的工具

命令：

> `filetransfer upload`  文件上传
>
> `filetransfer download xxx.xx` 指定文件名下载

流程:

> 上传文件时：开启http服务，输出二维码。手机扫码浏览器打开链接，选择文件上传。服务端文件接受完毕后服务自动关闭，程序结束
>
> 下载文件时：开启http服务，输出二维码。手机扫码浏览器打开链接，选择文件夹下载文件。完成后手动ctrl+c关闭服务，程序结束

灵感来自[qrcp](https://github.com/claudiodangelis/qrcp)

