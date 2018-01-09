<div class="user_pg">
	<div class="row">
		<div class="title col-md-offset-2 col-md-2">
			<h4>我的账户</h4>
		</div>
		<div class="col-md-offset-5 col-md-1">			
			<a href="/logout" class="btn btn-warning btn-sm"><span class="glyphicon glyphicon-log-out"></span> 退出</a>
		</div>
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
						<a class="pull-right" href="/user/edit">编辑</a>
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
						<span>{{if .Name}}{{.Name}}{{else}}{{.En_name}}{{end}}</span>
					</td>					
				</tr>
				{{else}}
				<tr>
					<td>
						<span class="text-warning">您还没有解锁的学校</span>
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
						<a class="pull-right" href="/pg/edit?id={{.Id}}">编辑</a>
					</td>					
				</tr>
				{{else}}
				<tr>
					<td>
						<span class="text-warning">您还没有发布过群</span>
					</td>
				</tr>
				{{end}}
			</table>
		</div>
	</div>
</div>