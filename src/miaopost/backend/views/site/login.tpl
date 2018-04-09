<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>北京绽德科技 | 秒Po后台管理</title>
  {{template "inc/base-style.tpl"}}
  <link rel="stylesheet" href="/static/plugin/adminlte/plugins/iCheck/square/blue.css">
  <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,600,700,300italic,400italic,600italic">
  <style type="text/css">
  .login-box, .register-box {
    width: 390px;
}
  </style>
</head>
<body class="hold-transition login-page">
<div class="login-box">
  <div class="login-logo">
    <a href="javascript:;"><b>北京绽德科技</b> - 秒Po后台</a>
  </div>
  <!-- /.login-logo -->
  <div class="login-box-body">

    <form id="login-form" action="/login" method="post">
      <div class="form-group has-feedback">
        <input type="email" class="form-control" id="username" placeholder="Email">
        <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
      </div>
      <div class="form-group has-feedback">
        <input type="password" class="form-control" id="password" placeholder="Password">
        <span class="glyphicon glyphicon-lock form-control-feedback"></span>
      </div>
      <div class="row">
        <div class="col-xs-8">
          <div class="checkbox icheck">
            <label>
              <input type="checkbox" checked="checked"> Remember Me
            </label>
          </div>
        </div>
        <!-- /.col -->
        <div class="col-xs-4">
          <button type="submit" class="btn btn-primary btn-block btn-flat">Sign In</button>
        </div>
        <!-- /.col -->
      </div>
    </form>

   
  </div>
  <!-- /.login-box-body -->
</div>
<!-- /.login-box -->

  {{template "inc/base-script.tpl"}}
<script src="/static/plugin/adminlte/plugins/iCheck/icheck.min.js"></script>
<script>
  $(function () {
    $('input').iCheck({
      checkboxClass: 'icheckbox_square-blue',
      radioClass: 'iradio_square-blue',
      increaseArea: '20%' /* optional */
    });

    $("form").submit(function(){
    	var username = $.trim($("#username").val());
    	var password = $.trim($("#password").val());
	 if (username=="" || password=="") {
	    	prompt("帐号密码不能为空");	    	
	} else {
	    	$.post("/login",{username:username,password:password},function(e){
	    		if (e.code<0) {
	    			prompt(e.msg);
	    			return false;
	    		}
	    		location.href = "/";
	    	});
	}
	return false; 
      });
  });

</script>
</body>
</html>