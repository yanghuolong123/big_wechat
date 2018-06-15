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
	<div class="comment ">
		<h4 class="text-muted">留言</h4>
		<div class="row">	
			<div class="col-md-7">	
				<input type="hidden" id="info_id" name="info_id" value="{{.info.Id}}">
				<textarea id="info_msg" class="form-control"></textarea>
			</div>
		</div>		 
		<div class="comment_btn col-sm-offset-6">
			<button class="btn btn-success info-msg-btn">提交</button>
		</div>
		<ul id="commentlist">
			{{range .imvos}}
			<li>				
				<h5>{{.User.Nickname}}</h5>
				<p>{{.Im.Content}}</p>
				<p>{{showtime .Im.Create_time}}</p>				
			</li>
			{{end}}
		</ul>
	</div>
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