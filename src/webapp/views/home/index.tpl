<!DOCTYPE HTML>
<html>
<head>
    <title>BigWechat-最简洁的中国留学生交流群。选校、接机、买卖二手、房屋出租,轻松搞定</title>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/static/plugin/bootstrap/css/bootstrap.min.css">    
    <link rel="stylesheet" type="text/css" href="/static/css/style.css">    
    <link rel="stylesheet" type="text/css" href="/static/css/jquery.mobile.flatui.css" />
</head>
<body>
<div data-role="page">
    <div data-role="header" class="header linear-g navbar-fixed-top">
        <a href="#panel-left" data-iconpos="notext" class="glyphicon glyphicon-user col-lg-2 col-xs-2 text-right"> </a>
	<div class="col-lg-6 input-group text-center col-xs-8">
	      <input type="text" class="form-control" placeholder="Search for...">
	      <span class="input-group-btn">
		<button class="btn btn-default" type="button">Go!</button>
	      </span>
    	</div>
    </div>
    <div data-role="panel" data-position="left" data-display="push" class="user_box text-center dn linear-g" id="panel-left">
        <div class="u_info">
            <img class="avatar" src="/static/images/avatar.png" alt="头像">
            <span class="username">{{.user.Nickname}}</span>
        </div>
        <ul class="user_menu">
          <li class="menu"><a href="#"><span class="glyphicon glyphicon-cog"> </span> &nbsp;基本设置</a></li>
          <li class="menu"><a href="#"><span class="glyphicon glyphicon-lock"> </span> &nbsp;修改密码</a></li>
          <li class="menu"><a href="#"><span class="glyphicon glyphicon-picture"> </span> &nbsp;上传头像</a></li>
          <li class="menu"><a href="#"><span class="glyphicon glyphicon-off"> </span> &nbsp;安全退出</a></li>
        </ul>
    </div>
    <div data-role="content" class="container" role="main">
        <ul id="msg-list" class="content-reply-box mg10 container resizable">          
           
        </ul>        
	<footer class="footer navbar-fixed-bottom linear-g ">
	  <div class="container">
		<div class="row sendsmg">
			<textarea id="msgContent" class="col-md-9"></textarea>			
			<input id="sendBtn" class="btn btn-success col-md-1" type="button" value="发送">
		</div>
	  </div>
	</footer>
    </div>
</div>

<div id="modalPage"></div>

{{template "../inc/script.tpl" .}}

</body>
</html>
