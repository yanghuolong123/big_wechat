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
 					<a class="btn btn-info btn-sm" href="/pg/create" role="button">发布群</a>
 				</div>		  	
			</div>
	 		<div class="content">
	 			{{.LayoutContent}}
	 		</div>			
		</div>
		<div id="modalPage"></div>	 	
             	<script src="/static/plugin/jquery/jquery-2.2.4.js"></script>
		<script src="/static/plugin/bootstrap/js/bootstrap.min.js"></script>
            	<script src="/static/js/addwechat.js"></script> 	
	 </body>
</html>