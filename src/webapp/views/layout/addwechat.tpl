<!DOCTYPE HTML>
<html>
	<head>
	    <title>Add Wechat-北美留学生微信群</title>
	    <meta charset="UTF-8">
	    <meta http-equiv="X-UA-Compatible" content="IE=edge">
	    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
	    <link rel="stylesheet" type="text/css" href="/static/plugin/bootstrap/css/bootstrap.min.css">
	    <link rel="stylesheet" type="text/css" href="/static/css/addwechat.css">  
	 </head>
	 <body>
	 	<div class="container">	 		 
 			<div class="page-header row">
 				<div class="logo col-md-10">
 					<a href="/"><img src="/static/images/logo10.png" /></a>
 				</div>
 				<div class="head-left col-md-2">
 					<!--<a class="btn btn-info btn-sm" href="/pg/create" role="button">发布群</a>-->
 					{{if .user}}
		                                            <input type="hidden" name="uid" id="uid" value="{{.user.Id}}">
		                                            <a  href="/pg/user" class="btn btn-info">
		                                              <span class="glyphicon glyphicon-user" aria-hidden="true"></span>  {{.user.Nickname}}
		                                            </a>                                            
		                                     {{else}}
		                                            <input type="hidden" name="uid" id="uid" value="0">
		                                            <button type="button" class="btn btn-info" id="loginBtn">
		                                              <span class="glyphicon glyphicon-log-in" aria-hidden="true"></span> 登录
		                                            </button>
		                                    {{end}}
 				</div>		  	
			</div>
	 		<div class="content">
	 			{{.LayoutContent}}
	 		</div>			
		</div>

		{{template "../login/login.tpl" .}} 

		<div id="modalPage">
			
		</div>	

             	<script src="/static/plugin/jquery/jquery-2.2.4.js"></script>
		<script src="/static/plugin/bootstrap/js/bootstrap.min.js"></script>
		<script src="/static/plugin/bootstrap/js/bootstrap3-typeahead.min.js"></script>
            	<script src="/static/js/addwechat.js"></script> 	
	 </body>
</html>