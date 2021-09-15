# Nodejs笔记

## 从 .env 文件中为NodeJS加载环境变量

在 Node.js 项目的根目录中创建一个 .env 文件。以 NAME = VALUE 的形式在添加特定的环境的变量。例：

```.env
ETHERSCAN_API_KEY=<Etherscan API KEY>
RINKEBY_URL=https://rinkeby.infura.io/v3/<projectID>
```

使用名为 dotenv 的 npm 模块将 .env 文件中的配置加载到代码中。安装：

```shell
npm install dotenv --save
```

安装好后，将以下代码添加到项目入口的最前面，以确保所有地方都可以使用到配置中的环境变量。

```js
require("dotenv").config();
```
