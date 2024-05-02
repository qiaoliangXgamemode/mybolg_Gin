package FriendChainManagement

// 类似于快递信息识别的东西。
// 个人博客名称：昼湖先生
// 个人博客网站：https://gitpor.cn
// 个人博客头像：https://gitpor.cn/46279874.jpg
// 个人博客介绍：I've been on the low, I been taking my time.

// Json parser structure
// {
// "blog_name":"昼湖先生"
// "blog_website":"https://gitpor.cn",
// "blog_personalphoto":"https://gitpor.cn/46279874.jpg",
// "blog_Description":"I've been on the low, I been taking my time.",
// }

type FriendChainStructure struct {
	blog_name 					string `json:"blog_name"`
	blog_website 				string `json:"blog_website"`
	blog_personalphoto 			string `json:"blog_personalphoto"`
	blog_Description 			string `json:"blog_Description"`
}

// response HTML5 Code friendshipUrl.vue
func ResponsefriendChina(Chain *FriendChainStructure) {
	
}
// register FriendChina to (*func)ResponsefriendChina
func newFriendChina(blog_name 	string, blog_website 	string, blog_personalphoto 	string, blog_Description 	string) {

}
