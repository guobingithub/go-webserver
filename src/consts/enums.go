package schema

// ClusterStatus 集群状态
type ClusterStatus uint8

const (
	ClusterStatusNormal   ClusterStatus = 1 // 可用
	ClusterStatusAbnormal ClusterStatus = 2 // 异常
	ClusterStatusMaintain ClusterStatus = 3 // 维护
)

// 主机类型
const (
	NodeTypeMaster = 1
	NodeTypeWorker = 2
)

// 流水线步骤描述
const (
	PipelineStepJavaBuild          = "Java构建"
	PipelineStepJavaUnitTest       = "Java单元测试"
	PipelineStepJavaBuildAndUpload = "Java构建上传"
	PipelineStepGoBuild            = "Go构建"
	PipelineStepGoUnitTest         = "Go单元测试"
	PipelineStepGoBuildAndUpload   = "Go构建上传"
)

// 事件动作类型
const (
	EventActionTypeCreate = 1 // 创建
	EventActionTypeRead   = 2 // 查询
	EventActionTypeUpdate = 3 // 更新
	EventActionTypeDelete = 4 // 删除
)

// 事件级别
const (
	EventLevelInfo     = 0 // 提醒
	EventLevelWarning  = 1 // 次要
	EventLevelError    = 2 // 重要
	EventLevelCritical = 3 // 紧急
)

// 权限许可
const (
	PermissionAllow = "allow"
	PermissionDeny  = "deny"
)

// 任务状态
const (
	TaskStatusEnable  = 1 // 启用
	TaskStatusDisable = 2 // 停用
	TaskStatusDeleted = 3 // 删除
)

// 用户状态
const (
	UserStatusToReview   uint8 = 1 //待审核
	UserStatusReviewOk   uint8 = 2 //通过
	UserStatusReviewBack uint8 = 3 //驳回
)
