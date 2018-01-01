<div class="user_pg">
	<div class="title col-md-offset-2">
		<h4>我的账户</h4>
	</div>
	<div class="row">
		<div class="col-md-offset-2 col-md-1">
			<label class="text-muted">昵称</label> 
		</div>
		<div class="col-md-6">
			<table class="table">
				<tr>
					<td>
						<span>{{.user.Nickname}}</span>
						<a class="pull-right" href="#">编辑</a>
					</td>
				</tr>
			</table>
		</div>
	</div>
	<div class="row">
		<div class="col-md-offset-2 col-md-1">
			<label class="text-muted">解锁学校</label>
		</div>
		<div class="col-md-6">
			<table class="table">
				{{range .ugs}}
				<tr>
					<td>
						<span>{{.Name}}</span>
					</td>					
				</tr>
				{{end}}
			</table>
		</div>
	</div>
	<div class="row">
		<div class="col-md-offset-2 col-md-1">
			<label class="text-muted">我的发布</label>
		</div>
		<div class="col-md-6">
			<table class="table">
				{{range .pgs}}
				<tr>
					<td>
						<span>{{.Name}}</span>
						<a class="pull-right" href="#">编辑</a>
					</td>					
				</tr>
				{{end}}
			</table>
		</div>
	</div>
</div>