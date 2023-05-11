package errors

import "github.com/limeschool/easy-admin/server/types"

var (
	//基础相关
	ServerError     = &types.Response{Code: 000001, Msg: "系统错误"}
	ServerBusyError = &types.Response{Code: 000002, Msg: "系统繁忙"}
	AssemblyError   = &types.Response{Code: 000003, Msg: "系统繁忙"}
	IpLimitError    = &types.Response{Code: 000004, Msg: "系统繁忙，请稍后再试"}

	ParamsError             = &types.Response{Code: 100002, Msg: "参数验证失败"}
	AssignError             = &types.Response{Code: 100003, Msg: "数据赋值失败"}
	DBError                 = &types.Response{Code: 100004, Msg: "数据库操作失败"}
	DBNotFoundError         = &types.Response{Code: 100005, Msg: "未查询到指定数据"}
	UserNotFoundError       = &types.Response{Code: 100006, Msg: "账号不存在"}
	UserDisableError        = &types.Response{Code: 100007, Msg: "账号已被禁用"}
	PasswordError           = &types.Response{Code: 100008, Msg: "账号密码错误"}
	RsaPasswordError        = &types.Response{Code: 100009, Msg: "非法账号密码"}
	SuperAdminEditError     = &types.Response{Code: 100011, Msg: "超级管理员不允许修改"}
	SuperAdminDelError      = &types.Response{Code: 100012, Msg: "超级管理员不允许删除"}
	RoleDisableError        = &types.Response{Code: 1000013, Msg: "账户角色已被禁用"}
	PasswordExpireError     = &types.Response{Code: 1000014, Msg: "登陆密码时效已过期"}
	CaptchaError            = &types.Response{Code: 1000015, Msg: "验证码错误"}
	CreateCaptchaError      = &types.Response{Code: 1000015, Msg: "生成验证码失败"}
	RefreshActiveTokenError = &types.Response{Code: 1000016, Msg: "禁止刷新可用的Token"}
	CaptchaSendError        = &types.Response{Code: 1000017, Msg: "邮箱验证码发送失败！"}
	DulSendCaptchaError     = &types.Response{Code: 1000015, Msg: "不能重复发送验证码，请稍后重试"}

	//auth相关
	NotResourcePowerError = &types.Response{Code: 4003, Msg: "暂无接口资源权限"}
	TokenExpiredError     = &types.Response{Code: 4001, Msg: "登陆信息已过期，请重新登陆"}
	RefTokenExpiredError  = &types.Response{Code: 4000, Msg: "太长时间未登陆，请重新登陆"}
	DulDeviceLoginError   = &types.Response{Code: 4000, Msg: "你已在其他设备登陆"}
	MetadataError         = &types.Response{Code: 4000, Msg: "获取用户元数据失败"}
	TokenDataError        = &types.Response{Code: 4000, Msg: "token数据异常失败"}
	TokenValidateError    = &types.Response{Code: 4000, Msg: "token验证失败"}
	TokenEmptyError       = &types.Response{Code: 4000, Msg: "token信息不存在"}

	//menu相关
	DulMenuNameError    = &types.Response{Code: 1000030, Msg: "菜单name值不能重复"}
	MenuParentIdError   = &types.Response{Code: 1000031, Msg: "父菜单id值异常"}
	DeleteRootMenuError = &types.Response{Code: 1000032, Msg: "不能删除根菜单"}

	//team相关
	NotAddTeamError      = &types.Response{Code: 1000040, Msg: "暂无此部门的下级部门创建权限"}
	NotEditTeamError     = &types.Response{Code: 1000041, Msg: "暂无此部门的修改权限"}
	NotDelTeamError      = &types.Response{Code: 1000042, Msg: "暂无此部门的删除权限"}
	NotAddTeamUserError  = &types.Response{Code: 1000043, Msg: "暂无此部门的人员创建权限"}
	NotEditUserRoleError = &types.Response{Code: 1000043, Msg: "暂无此人员的角色修改权限"}
	NotEditTeamUserError = &types.Response{Code: 1000044, Msg: "暂无此部门的人员修改权限"}
	NotDelTeamUserError  = &types.Response{Code: 1000045, Msg: "暂无此部门的人员删除权限"}
	TeamParentIdError    = &types.Response{Code: 1000046, Msg: "父部门不能为自己"}

	//role相关
	DulKeywordError     = &types.Response{Code: 1000050, Msg: "角色标志符已存在"}
	DisableCurRoleError = &types.Response{Code: 1000051, Msg: "不能禁用当前用户所在角色"}

	// upload相关
	InitUploadError           = &types.Response{Code: 1000060, Msg: "文件上传初始化失败"}
	UploadLimitMaxSizeError   = &types.Response{Code: 1000061, Msg: "文件超过规定大小限制"}
	OpenFileError             = &types.Response{Code: 1000062, Msg: "上传文件打开失败"}
	UploadTypeNotSupportError = &types.Response{Code: 1000063, Msg: "不支持上传此类型文件"}
	UploadTypeError           = &types.Response{Code: 1000064, Msg: "存在不允许上传的文件类型"}
	UploadDirError            = &types.Response{Code: 1000065, Msg: "非法的上传目录"}

	// registry
	InitDockerError         = &types.Response{Code: 1000070, Msg: "初始化docker-cli 失败"}
	LoginImageRegistryError = &types.Response{Code: 1000071, Msg: "登陆镜像仓库失败"}
	LoginCodeRegistryError  = &types.Response{Code: 1000071, Msg: "登陆代码仓库失败"}

	// cluster
	ParseClusterConfigError = &types.Response{Code: 1000080, Msg: "解析集群config失败"}
	ConnectClusterError     = &types.Response{Code: 1000081, Msg: "连接集群失败"}
	GetClusterError         = &types.Response{Code: 1000081, Msg: "获取集群数据失败"}
	UpdateClusterError      = &types.Response{Code: 1000081, Msg: "更新集群数据失败"}
)
