package routers

func init() {
	//前台界面
	FrontendRouter()

	//后台界面
	BackendRouter()

	//api接口
	ApiRouter()

	//杂项路由
	MiscRouter()
}
