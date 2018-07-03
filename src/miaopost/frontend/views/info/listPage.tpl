{{range .infos}}
<div class="info">
	<div class="row">
		<div class="col-md-8 col-xs-3 cat"><span class="label label-warning">{{.Cat.Name}}</span></div>
		<div class="col-md-4 col-xs- 9 meta text-right">
			<span>{{showtime .Info.Update_time}} </span>
			{{if $.isMy}}
			<span><a href="#" onclick="topInfo({{.Info.Id}});return false;">置顶</a></span>
			<span><a href="/info/edit?id={{.Info.Id}}">编辑</a></span>
			<span><a href="#" onclick="delInfo({{.Info.Id}}, this);return false;">删除</a></span>
			{{else}}
			<span>阅读({{.Info.Views}})</span>
			<span><a href="#" onclick="suggestDel({{.Info.Id}});return false;">建议删除</a></span>
			{{end}}
		</div>
	</div>
	<div class="row">
		<div class="info-content col-md-12">			
			<a href="/info/view?id={{.Info.Id}}" {{if not $.isMobile}}target="_blank"{{end}} class="list_content">{{str2html (showListInfo (substr .Info.Content 0 150))}} {{if (.Photos|len)}}<img class="img_tip" src="/static/img/image_s.png"/>{{end}}{{if eq (gt .Info.Reward_type 0) (eq .Info.Reward_pay 1)}}<img class="img_tip" src="/static/img/reward_type{{.Info.Reward_type}}.png"/>{{end}} ...</a>
			<div class="line"></div>
		</div>
	</div>
</div>
{{else}}
	<div class="alert alert-warning col-md-10" role="alert">亲，还没有数据哦！</div>
{{end}}