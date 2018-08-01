<input type="hidden" id="info_id" value="{{.info.Id}}">
<input type="hidden" id="chance" value="{{.chance}}">
<div class="view">
	<div class="row meta">
		<div class="col-md-9 col-xs-9">
			<span class="badge">{{.cat.Name}}</span>
			<span>阅读({{.info.Views}}) </span>
			<span>{{showtime .info.Update_time}}</span>
		</div>
		<div class="col-md-2 col-xs-3">
			<a href="#" onclick="suggestDel({{.info.Id}});return false;">建议删除</a>
		</div>
	</div>
	<div class="row">
		<div class="col-md-8">
			{{str2html .info.Content}}  
			<br>
			<span class="text-info more">联系我时请注明来自秒Po</span>
		</div>
		<div class="col-md-8 photos">
			{{range .photos}}
			<img src="{{.P.Url}}{{cutImgSize .Width .Height}}">
			{{end}}
		</div>
	</div>
	
	{{if .info.Uid}}
	{{template "info/infoMsg.tpl" .}}
	{{end}}
	
	<div class="adv row">
		{{if .adv}}
		<ul>
		{{range .adv}}
		<li><a href="{{if .Link}}{{.Link}}{{else}}/article/view?id={{.Id}}{{end}}"><img src="{{.Logo}}!800!800" /></a></li>
		{{end}}
		</ul>
		{{end}}
	</div>
</div>

{{if  eq .info.Reward_pay 1}}
<script type="text/javascript">

var getWithDraw = function() {
	var chance = $("#chance").val();
	if (chance=="no") {
		return false;
	}
	$.post("/reward/chance",{info_id:$("#info_id").val()}, function(e){
		if(e.code<=0) {
			return false;
		}

		if(e.code==1) {
			greeting({title:"提示",msg:"恭喜您! 您获得了 "+e.data.Amount+"元 阅读红包。"})	
		} else if(e.code==2) {
			greeting({title:"提示",msg:"恭喜您获得了一个 "+e.data.Amount+"元 的留言红包!，请在45分钟内完成留言。如45分钟后未完成留言，红包将释放给其他用户"})
		}
	})
};

$(function(){
	getWithDraw();
});
</script>
{{end}}