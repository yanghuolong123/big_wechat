<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8"/>
    <title>Sample of websocket with golang</title>
    <script src="/static/plugin/jquery/jquery-3.2.1.min.js"></script>

    <script>
	function Message(author, body) {
		this.author = author;
		this.body = body;
		//this.createtime = new Date();
		this.createtime;
	}
        var ws;
      $(function() {
        var $ul = $('#msg-list');
        ws = new WebSocket("ws://"+ window.location.host +"/chat");
        ws.onmessage = function(e) {
          var msg = JSON.parse(e.data);
	  var content = "时间:"+ msg.createtime + " 内容:"+ msg.body;
          $('<li>').text(content).appendTo($ul);
        };
        $('#sendBtn').click(function(){
          var body = $('#name').val();
	 var msg = new Message("yhl", body); 
          ws.send(JSON.stringify(msg));
                $('#name').val("");
        });
      });
    </script>
</head>
<body>
<h2>{{.welcome}}</h2>
<input id="name" type="text"/>
<input type="button" id="sendBtn" value="send"/>
<ul id="msg-list"></ul>
</body>
</html>
