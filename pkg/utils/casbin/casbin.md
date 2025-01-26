## 用户

* root : 超级管理员  18888888888
* admin : 管理员     18811111111
* service : 服务商   18382131111
* agency : 代理商    18382132222 
* store : 商家       18382133333
* operation : 运营   18833333333



# 角色

* root : 超级管理员
* admin : 管理员
* service : 服务商
* agency : 代理商
* store : 商家
* operation : 运营


## 资源

## 资源

* G : 管理员
  * 概览 Dashboard
  	* 工作台  Workspace
  * 商家管理  merchantManage
    * 商家列表  merchantManageUser
    * 外卖专家  takeawayExpert
    * 外卖专家  takeawayExpertAgent
  * 代理管理  agentManage
    * 服务商  agentManageOne
    * 下级代理  agentManageTwo
  * 私域管理  fansManage
    * 粉丝看板  fansKanban
    * 粉丝群发  fansGroupPosting
    * 粉丝群发任务发布  fansGroupPostingCreate
    * 粉丝群发任务编辑  fansGroupPostingUpdate
    * 企微主体  enterpriseMicroEntities
    * 粉丝列表  enterpriseMicroEntitiesFansList
  * 数据管理  dataManage
    * 佣金结算  commissionSettlement
  * 系统管理  systemManage
    * 支付设置  systemManagePay
  * 日志管理  logManage
    * 操作日志  logManageList

----
* F : 服务商
  * 概览 Dashboard
  	* 工作台  Workspace
  * 商家管理  merchantManage
    * 商家列表  merchantManageUser
    * 外卖专家  takeawayExpertAgentList
  * 代理管理  agentManage
    * 下级代理  agentManageTwo
  * 私域管理  fansManage
    * 企微主体  enterpriseMicroEntities
    * 粉丝列表  enterpriseMicroEntitiesFansList
  * 数据管理  dataManage
    * 佣金结算  commissionSettlement
  * 系统管理  systemManage
    * 主体信息  systemManageUser
  * 日志管理  logManage
    * 操作日志  logManageList


----
* D : 代理商
  * 概览 Dashboard
  	* 工作台  Workspace
  * 商家管理  merchantManage
    * 商家列表  merchantManageUser
    * 外卖专家  takeawayExpertAgentList
  * 私域管理  fansManage
    * 企微主体  enterpriseMicroEntities
    * 粉丝列表  enterpriseMicroEntitiesFansList
  * 数据管理  dataManage
    * 佣金结算  commissionSettlement
  * 系统管理  systemManage
    * 主体信息  systemManageUser
  * 日志管理  logManage
    * 操作日志  logManageList

----
* S : 商家
  * 概览 Dashboard
  	* 工作台  Workspace
  * 私域管理  fansManage
    * 粉丝列表  fansList
  * 数据管理  dataManage
    * 佣金结算  commissionSettlement
  * 系统管理  systemManage
    * 主体信息  systemManageUser
  * 日志管理  logManage
    * 操作日志  logManageList

## 权限
 
* read : 只读权限

## 租户

* ldx : 路大兮


## 示例

p, admin, ldx, data1, read
g, 18888888888, admin, ldx



```
# Policy for root (超级管理员)
p, root, ldx, Dashboard, read
p, root, ldx, Workspace, read
p, root, ldx, merchantManage, read
p, root, ldx, merchantManageUser, read
p, root, ldx, takeawayExpert, read
p, root, ldx, takeawayExpertAgent, read
p, root, ldx, agentManage, read
p, root, ldx, agentManageOne, read
p, root, ldx, agentManageTwo, read
p, root, ldx, fansManage, read
p, root, ldx, fansKanban, read
p, root, ldx, fansGroupPosting, read
p, root, ldx, fansGroupPostingCreate, read
p, root, ldx, fansGroupPostingUpdate, read
p, root, ldx, enterpriseMicroEntities, read
p, root, ldx, enterpriseMicroEntitiesFansList, read
p, root, ldx, dataManage, read
p, root, ldx, commissionSettlement, read
p, root, ldx, systemManage, read
p, root, ldx, systemManagePay, read
p, root, ldx, logManage, read
p, root, ldx, logManageList, read

# Policy for admin (管理员)
p, admin, ldx, Dashboard, read
p, admin, ldx, Workspace, read
p, admin, ldx, merchantManage, read
p, admin, ldx, merchantManageUser, read
p, admin, ldx, takeawayExpert, read
p, admin, ldx, takeawayExpertAgent, read
p, admin, ldx, agentManage, read
p, admin, ldx, agentManageOne, read
p, admin, ldx, agentManageTwo, read
p, admin, ldx, fansManage, read
p, admin, ldx, fansKanban, read
p, admin, ldx, fansGroupPosting, read
p, admin, ldx, fansGroupPostingCreate, read
p, admin, ldx, fansGroupPostingUpdate, read
p, admin, ldx, enterpriseMicroEntities, read
p, admin, ldx, enterpriseMicroEntitiesFansList, read
p, admin, ldx, dataManage, read
p, admin, ldx, commissionSettlement, read
p, admin, ldx, systemManage, read
p, admin, ldx, systemManagePay, read
p, admin, ldx, logManage, read
p, admin, ldx, logManageList, read

# Policy for service (服务商)
p, service, ldx, Dashboard, read
p, service, ldx, Workspace, read
p, service, ldx, merchantManage, read
p, service, ldx, merchantManageUser, read
p, service, ldx, takeawayExpertAgentList, read
p, service, ldx, agentManage, read
p, service, ldx, agentManageTwo, read
p, service, ldx, fansManage, read
p, service, ldx, enterpriseMicroEntities, read
p, service, ldx, enterpriseMicroEntitiesFansList, read
p, service, ldx, dataManage, read
p, service, ldx, commissionSettlement, read
p, service, ldx, systemManage, read
p, service, ldx, systemManageUser, read
p, service, ldx, logManage, read
p, service, ldx, logManageList, read

# Policy for agency (代理商)
p, agency, ldx, Dashboard, read
p, agency, ldx, Workspace, read
p, agency, ldx, merchantManage, read
p, agency, ldx, merchantManageUser, read
p, agency, ldx, takeawayExpertAgentList, read
p, agency, ldx, fansManage, read
p, agency, ldx, enterpriseMicroEntities, read
p, agency, ldx, enterpriseMicroEntitiesFansList, read
p, agency, ldx, dataManage, read
p, agency, ldx, commissionSettlement, read
p, agency, ldx, systemManage, read
p, agency, ldx, systemManageUser, read
p, agency, ldx, logManage, read
p, agency, ldx, logManageList, read

# Policy for store (商家)
p, store, ldx, Dashboard, read
p, store, ldx, Workspace, read
p, store, ldx, fansManage, read
p, store, ldx, fansList, read
p, store, ldx, dataManage, read
p, store, ldx, commissionSettlement, read
p, store, ldx, systemManage, read
p, store, ldx, systemManageUser, read
p, store, ldx, logManage, read
p, store, ldx, logManageList, read

# Policy for operation (运营)
p, operation, ldx, Dashboard, read
p, operation, ldx, Workspace, read
p, operation, ldx, merchantManage, read
p, operation, ldx, merchantManageUser, read
p, operation, ldx, takeawayExpert, read
p, operation, ldx, takeawayExpertAgent, read
p, operation, ldx, agentManage, read
p, operation, ldx, agentManageOne, read
p, operation, ldx, agentManageTwo, read
p, operation, ldx, fansManage, read
p, operation, ldx, fansKanban, read
p, operation, ldx, fansGroupPosting, read
p, operation, ldx, fansGroupPostingCreate, read
p, operation, ldx, fansGroupPostingUpdate, read
p, operation, ldx, enterpriseMicroEntities, read
p, operation, ldx, enterpriseMicroEntitiesFansList, read
p, operation, ldx, dataManage, read
p, operation, ldx, commissionSettlement, read
p, operation, ldx, systemManage, read
p, operation, ldx, systemManagePay, read
p, operation, ldx, logManage, read
p, operation, ldx, logManageList, read

# Grouping policy
g, 18888888888, root, ldx
g, 18811111111, admin, ldx
g, 18382131111, service, ldx
g, 18382132222, agency, ldx
g, 18382133333, store, ldx
g, 18833333333, operation, ldx
```



CREATE TABLE `menu_permissions` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键，菜单唯一标识',
  `parent_id` int DEFAULT NULL COMMENT '父级菜单 ID，目录为 NULL',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由路径',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '菜单类型，0: 目录，1: 菜单',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '菜单图标',
  `order` int DEFAULT '0' COMMENT '排序，越小越靠前',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应组件路径',
  `keep_alive` tinyint NOT NULL DEFAULT '1' COMMENT '是否缓存，1: 缓存，0: 不缓存',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '菜单状态，1: 启用，0: 禁用',
  `active_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '配置当前激活的菜单',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '菜单标题',
  `hide_in_menu` tinyint NOT NULL DEFAULT '0' COMMENT '是否在菜单中隐藏，0: 不隐藏，1: 隐藏',
  `hide_in_tab` tinyint NOT NULL DEFAULT '0' COMMENT '是否在标签页中隐藏，0: 不隐藏，1: 隐藏',
  `hide_in_breadcrumb` tinyint NOT NULL DEFAULT '0' COMMENT '是否在面包屑中隐藏，0: 不隐藏，1: 隐藏',
  `affix_tab` tinyint NOT NULL DEFAULT '0' COMMENT '是否固定标签页，0: 不固定，1: 固定',
`created_at` bigint NOT NULL DEFAULT '1735196391839' COMMENT '创建时间',
  `updated_at` bigint NOT NULL DEFAULT '1735196391839' COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=130 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='系统菜单表，存储菜单信息及其属性';
