<div class="user">
	<div class="row">
		<div class="title col-md-offset-1 col-md-2">
			<h4>我的账户</h4>
		</div>
		<div class="col-md-2">
			<h4><a href="/info/my">我的发布</a></h4>
		</div>
		<div class="col-md-offset-5 col-md-1">			
			<a href="/logout" class="btn btn-warning btn-sm"><span class="glyphicon glyphicon-log-out"></span> 退出</a>
		</div>
	</div>
	<div class="row">
		<div class="col-md-offset-1 col-md-1">
			<label class="text-muted">昵称</label> 
		</div>
		<div class="col-md-9">
			<table class="table">
				<tr>
					<td>
						<span>{{.user.Nickname}}</span>
						<a class="pull-right" href="/user/edit">编辑</a>
					</td>
				</tr>
			</table>
		</div>
	</div>	
</div>