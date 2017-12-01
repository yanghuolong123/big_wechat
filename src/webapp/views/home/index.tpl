<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8"/>
    <title>Sample of websocket with golang</title>
    <script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>

    <script>
	function Message(author, body) {
		this.author = author;
		this.body = body;
	}
        var ws;
      $(function() {
        var $ul = $('#msg-list');
        ws = new WebSocket("ws://localhost:8080/chat");
        ws.onmessage = function(e) {//alert(e);
		var msg = JSON.parse(e.data);
          $('<li>').text(msg.body).appendTo($ul);
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
