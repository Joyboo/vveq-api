package controllers

// Operations about Tag
type TagController struct {
	BaseController
}

// @Title CreateTheme
// @Description create Theme
// @Param	body		body 	models.Tag	true		"body for Tag content"
// @Success 200 {int} model.Tag.Id
// @Failure 403 body is empty
// @router / [post]
func (this *TagController) Post() {
}
