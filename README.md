### casbin 最佳实践

### casbin 介绍
Casbin是一个强大的、高效的开源访问控制框架，其权限管理机制支持多种访问控制模型。
|支持的权限模型|
|-|
|ACL|
|RBAC|
|ABAC|
|......|
[官网](https://casbin.org/)

#### [RBAC 单角色](https://github.com/MasterJoyHunan/casbin-demo/tree/master)
rbac0 单角色模型，可以看做 acl 一样的模型。

只不过一个是对角色进行分配权限，一个对用户分配权限。

在 casbin 中对其进行了抽象 （sub 对象对 obj 资源进行 act 操作）

这里 sub 可以理解为角色，也可以同样理解为用户

#### [RBAC 多角色](https://github.com/MasterJoyHunan/casbin-demo/tree/rbac0)

rbac0 多角色模型：一个用户可以拥有多个角色，这是一种非常常见的需求

在 casbin 中，将单角色修改为多角色易如反掌，其核心就是在配置中加入

```config
[role_definition]
g = _, _

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && r.act == p.act
```
定义 `policy`
```cvs
p, admin, data1, read  // 可以理解为 (用户/角色) 对资源 data1 可以 read 操作
p, admin, data1, write  // 可以理解为 (用户/角色) 对资源 data1 可以 write 操作
p, editor, data2, read
p, editor, data2, write
g, alice, admin  // [重点] 可以理解为 admin 角色含有 alice 用户
g, alice, editor // [重点] 可以理解为 editor 角色含有 alice 用户
```
其他代码完全不需要改变，即可完成了多角色的定义

#### [RBAC 角色继承](https://github.com/MasterJoyHunan/casbin-demo/tree/rbac1)

角色继承在可以完全复用多角色的代码，仅仅修改 `policy` 即可实现
```cvs

```

#### LICENCE
[MIT](https://en.wikipedia.org/wiki/MIT_License)