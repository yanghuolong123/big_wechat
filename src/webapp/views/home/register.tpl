<div id="registerModal" class="modal fade" tabindex="-1" role="dialog" aria-labelledby="gridSystemModalLabel">
  <div class="modal-dialog" role="document">
    <div class="modal-content">      
      <div class="modal-body">
        <form>
	  <div class="form-group">
	    <label for="group">学校</label>
	    <select class="form-control" id="group" >
		<option value="0">请选择</option>
		{{range .groupList}}
  		<option value="{{.Id}}">{{.Name}}</option>
		{{end}}
	    </select>
	  </div>
	  <div class="form-group">
	    <label for="username">账号</label>
	    <input type="text" class="form-control" id="username" placeholder="电话或邮箱">
	  </div>
	  <div class="form-group">
	    <label for="nickname">昵称</label>
	    <input type="text" class="form-control" id="nickname" placeholder="昵称">
	  </div>
	  <div class="form-group">
	    <label for="password">密码</label>
	    <input type="password" class="form-control" id="password" placeholder="密码">
	  </div>	  
	  <div class="form-group">
	    <label for="repassword">重复密码</label>
	    <input type="password" class="form-control" id="repassword" placeholder="重复密码">
	  </div>	  
	</form>
      </div>
      <div class="modal-footer">
	<!-- <button type="button" class="btn btn-default" data-dismiss="modal">取消</button> -->
        <button id="registerBtn" type="button" class="btn btn-primary">注册</button>
	&nbsp;&nbsp;&nbsp;&nbsp;
	<a id="loginLink" href="">登录</a>
	&nbsp;&nbsp;&nbsp;&nbsp;
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
