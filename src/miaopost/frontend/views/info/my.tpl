<div class="row">
	<input type="hidden" id="hasMore"  value="{{.hasMore}}" />
	<input type="hidden" id="page"  value="{{.page}}" />
	<input type="hidden" id="uid"  value="{{.uid}}" />
	<div class="col-md-12 info-list">
	{{template "info/listPage.tpl" .}}
	<a href="javascript:;" class="load-more">加载更多<span class="loading"></span></a>
	</div>
</div>