package controllers

// Operations about Cate
type CateController struct {
	BaseController
}

// @Title CreateTheme
// @Description create Theme
// @Param	body		body 	models.Cate	true		"body for Cate content"
// @Success 200 {int} model.Cate.Id
// @Failure 403 body is empty
// @router / [post]
func (this *CateController) Post() {
}
