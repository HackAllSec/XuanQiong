# 积分计算规则

总积分=漏洞评分 + Poc评分 + Exp评分 + 影响范围评分 + 其它评分

漏洞评分= CVSS 评分 × 10

Poc评分：

- Xray、Nuclei、Goby等完整Poc，误报低：20分
- Xray、Nuclei、Goby等完整Poc，误报较高：10分
- 仅包含Payload或无法工具化的Poc： 5 分
    
Exp评分：

- Xray、Nuclei、Goby等完整Exp，误报低：30分
- Xray、Nuclei、Goby等完整Exp，误报较高：15分
- 仅包含Payload或无法工具化的Exp：5分
   
影响范围评分：

- 互联网资产数大于 5000：30 分
- 互联网资产数介于 1000 到 5000：20 分
- 互联网资产数小于 1000：10 分

其它评分（自定义）：

其它评分 = 积分规则1分数 x 积分系数1 + 积分规则2分数 x 积分系数2 + ... + 积分规则n分数 x 积分系数n

系数为浮点数，评分结果将进行四舍五入取整
