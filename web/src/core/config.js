/**
 * 网站配置文件
 */

import { ref } from 'vue'

const config = {
  appName: '酷儿乐',
  appLogo: '/src/assets/logo.jpg',
  showViteLogo: true,
  comboTypeOptions: [
    { id: 1, label: 'Vip次卡' },
    { id: 2, label: 'Vip周期卡' },
    { id: 3, label: 'Vip充值卡' },
  ],
  memberStateOptions: [
    { id: 1, label: 'Vip会员' },
    { id: 2, label: '会员已过期' },
    { id: 3, label: '已禁用' },
    { id: 4, label: '非会员' },
    { id: 5, label: '已退款' },
    { id: 6, label: '已删除' },
    { id: 7, label: '待审核' },
  ],
  stateOptions: {
    0: "正常",
    1: "禁用"
  }
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用Gin-Vue-Admin，开源地址：https://github.com/flipped-aurora/gin-vue-admin`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本:v2.5.6`
      )
    )
    console.log(
      chalk.green(
        `> 插件市场:https://plugin.gin-vue-admin.com`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
      )
    )
    console.log('\n')
  }
}

export default config
