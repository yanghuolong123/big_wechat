<div class="edit_user">
	<div class="user_menu">
		<a href="/info/my">我的发布</a>
		<a class="label label-primary" href="/user/edit">标示修改</a>
		{{if not .isWeixin}}		
		<a href="/logout" class=""><span class="glyphicon glyphicon-log-out"></span> 退出</a>
		{{end}}
	</div>

	<form class="form-horizontal" method="post">
	  <div class="form-group">
	    <label for="nickname" class="col-sm-2 control-label">昵称</label>
	    <div class="col-sm-6">
	      <input type="text" class="form-control" name="nickname" id="nickname" maxlength="32" placeholder="昵称" value="{{.user.Nickname}}">
	    </div>
	  </div>
	  
	  <div class="form-group">
	    <div class="col-sm-offset-2 col-sm-6">
	      <button type="submit" class="btn btn-primary">保存</button>
	    </div>
	  </div>
	</form>	
</div>