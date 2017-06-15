#Tools
1. Chrome :F12打开Debug 模式，点击Network就可以看到页面的交互流程以及每一步交互的数据（包括异步获取的，ajax, js）. 同时通过Chrome的Debug 模式Elements选项可以查看动态生成页面的最终DOM文件。
2. Phantomjs 很有用
3. Casperjs 很有用
4. golang html/template的分割符“{{}}”可以通过 template.Template的func (t *Template) Delims(left, right string) *Template来重定义。通过这种方法可以解决golang template与一些前端框架（比如说vue）的分割符冲突问题。
5. vue.js的template 分割符“{{}}”可以通过vue.Look 或者vue.config.delimiters=XXX来重新定义。这可以用来解决分割符冲突问题。
6.  一个典型的golang web server: (js/css/image/file通过FileServer来实现). 
	```http.HandleFunc("/", MainPage)
	http.Handle("/asset/web/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)```

7. 注意golang的http.Handle中路由表项末尾带"/"和不带"/"其结果是截然不同的.
8. 要在给客户端的回复中带上Cookie可以 通过http.SetCookie来设置)
#9  Web APP的开发要尽量做到前后端分离，也就是说server端不应该对前端的表示做任何假设，同样前端也不应该对server端的实现做任何假设。前端通过访问服务端定义的一组良好的接口来获取和提交数据，（这也是前端和后端应该仅有的交互）
10. chrome的console中可以直接修改页面目前的参数,以及脚本中的变量.
11. vue // 模板定界符 修改
	Vue.config.delimiters = ['{{', '}}']

	// html模板定界符
	Vue.config.unsafeDelimiters = ["{{{", "}}}"]

12. 要将前端的表单数据转换成json(JSON.stringify())以后再传到后端，这样就会使后端的表单解析得到极大的简化。
