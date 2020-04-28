### casbin 最佳实践

#### [RBAC 单角色](https://github.com/MasterJoyHunan/casbin-demo/tree/master)
rbac0 单角色模型，可以看做 acl 一样的模型。

只不过一个是对角色进行分配权限，一个对用户分配权限。

在 casbin 中对其进行了抽象 （sub 对象对 obj 资源进行 act 操作）

这里 sub 可以理解为角色，也可以同样理解为用户

#### [RBAC 多角色](https://github.com/MasterJoyHunan/casbin-demo/tree/rbac0)

#### [RBAC 角色继承](https://github.com/MasterJoyHunan/casbin-demo/tree/rbac1)

#### LICENCE
[MIT](https://en.wikipedia.org/wiki/MIT_License)