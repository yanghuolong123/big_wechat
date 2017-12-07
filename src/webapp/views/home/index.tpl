<!DOCTYPE HTML>
<html>
<head>
    <title>BigWechat-最简洁的中国留学生交流群。选校、接机、买卖二手、房屋出租,轻松搞定</title>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" type="text/css" href="/static/plugin/bootstrap/css/bootstrap.min.css">    
    <link rel="stylesheet" type="text/css" href="/static/css/style.css">    
    <link rel="stylesheet" type="text/css" href="/static/css/jquery.mobile.flatui.css" />
</head>
<body>
<div data-role="page">
    <div data-role="header" class="header linear-g">
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
            <span class="username">developer</span>
        </div>
        <ul class="user_menu">
          <li class="menu"><a href="#"><span class="glyphicon glyphicon-cog"> </span> &nbsp;基本设置</a></li>
          <li class="menu"><a href="#"><span class="glyphicon glyphicon-lock"> </span> &nbsp;修改密码</a></li>
          <li class="menu"><a href="#"><span class="glyphicon glyphicon-picture"> </span> &nbsp;上传头像</a></li>
          <li class="menu"><a href="#"><span class="glyphicon glyphicon-off"> </span> &nbsp;安全退出</a></li>
        </ul>
    </div>
    <div data-role="content" class="container" role="main">
        <ul class="content-reply-box mg10">
            <li class="odd">
                <a class="user" href="#"><img class="img-responsive avatar_" src="/static/images/avatar-1.png" alt=""><span class="user-name">奔波儿灞</span></a>
                <div class="reply-content-box">
                	<span class="reply-time">03-08 15：00</span>
                    <div class="reply-content pr">
                    	<span class="arrow">&nbsp;</span>
                    	为什么小鑫的名字里有三个金呢？
                    </div>
                </div>
            </li>
            <li class="even">
                <a class="user" href="#"><img class="img-responsive avatar_" src="/static/images/avatar-1.png" alt=""><span class="user-name">灞波儿奔</span></a>
                <div class="reply-content-box">
                	<span class="reply-time">03-08 15：10</span>
                    <div class="reply-content pr">
                    	<span class="arrow">&nbsp;</span>
                    	他命里缺金，所以取名叫鑫，有些人命里缺水，就取名叫淼，还有些人命里缺木就叫森。
                    </div>
                </div>
            </li>
            <li class="odd">
                <a class="user" href="#"><img class="img-responsive avatar_" src="/static/images/avatar-1.png" alt=""><span class="user-name">奔波儿灞</span></a>
                <div class="reply-content-box">
                	<span class="reply-time">03-08 15：20</span>
                    <div class="reply-content pr">
                    	<span class="arrow">&nbsp;</span>
                    	那郭晶晶命里缺什么？
                    </div>
                </div>
            </li>
            <li class="even">
                <a class="user" href="#"><img class="img-responsive avatar_" src="/static/images/avatar-1.png" alt=""><span class="user-name">灞波儿奔</span></a>
                <div class="reply-content-box">
                	<span class="reply-time">03-08 15：30</span>
                    <div class="reply-content pr">
                    	<span class="arrow">&nbsp;</span>
                    	此处省略一百字...
                    </div>
                </div>
            </li>
        </ul>        
	<div class="row ">
		<div class="form-group">			
			<textarea class="form-control"></textarea>			
		</div>
		<div class="form-group">
			<div class="col-sm-offset-11">
				<input class="btn btn-default" type="submit" value="发送">
			</div>
		</div>
	</div>
    </div>
</div>

<script src="/static/plugin/jquery/jquery-2.2.4.js"></script>
<script src="/static/plugin/jquery.mobile/jquery.mobile-1.4.5.min.js"></script>
<script src="/static/plugin/bootstrap/js/bootstrap.min.js"></script>
<script type="text/javascript">
	$(function(){
		/* 
		** 不同页面切换转场效果
		** $.mobile.changePage ('/test.html', 'slide/pop/fade/slideup/slidedown/flip/none', false, false);
		*/
		$('.list-group-item,.menu a').click(function(){
			$.mobile.changePage($(this).attr('href'), {
				transition : 'flip', //转场效果
				reverse : true       //默认为false,设置为true时将导致一个反方向的转场
			});	
		});
	});
</script>
</body>
</html>
