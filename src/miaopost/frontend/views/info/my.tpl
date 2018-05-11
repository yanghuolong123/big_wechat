<div class="row my">
	<input type="hidden" id="hasMore"  value="{{.hasMore}}" />
	<input type="hidden" id="page"  value="{{.page}}" />
	<input type="hidden" id="uid"  value="{{.uid}}" />
	<div class="user_menu col-md-12">
		<a class="label label-primary" href="/info/my">我的发布</a>
		<a href="/user/edit">标示修改</a>
		{{if not .isWeixin}}		
		<a href="/logout" class=""><span class="glyphicon glyphicon-log-out"></span> 退出</a>
		{{end}}
	</div>
	<div class="col-md-12 info-list">
		{{template "info/listPage.tpl" .}}
		<a href="javascript:;" class="load-more">加载更多<span class="loading"></span></a>
	</div>
</div>