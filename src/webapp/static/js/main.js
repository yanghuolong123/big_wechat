
var ws = {};
var $ul;

$(function() {
	$ul = $('#msg-list');
	if(sessionId>0) {
		listen()
	}

	$('#sendBtn').click(function(){
	  if(sessionId==0) {
		return
	  }
	  var content = $('#name').val();
	  var msg = {};

	  msg.uid= sessionId;
	  msg.gid = gids; 
	  msg.type = "message";
	  msg.content= content;
	  ws.send(JSON.stringify(msg));

	  $('#name').val("");
        });

	$('#login').click(function(){
		$.post("/login", {email:"yhl27ml@163.com", password:"654321"}, function(e){
			sessionId = e.data.user.Id;
			gids = e.data.gids
			listen()
		});
	});

	$("#logout").click(function(){
		$.post("/logout", function(e){
			alert(e.msg);
		});
	});

});

function listen() {
	ws = new WebSocket("ws://"+ window.location.host +"/chat");

        ws.onopen = function(e) {
          var msg = {}
          msg.uid= sessionId;
          msg.gid = gids;
          msg.type = "login";
          ws.send(JSON.stringify(msg));

          console.log("open");
        }

        ws.onmessage = function(e) {
          var msg = JSON.parse(e.data);
          var content = "时间:"+ msg.createTime + " 内容:"+ msg.content;
          $('<li>').text(content).appendTo($ul);
        };

        ws.onclose = function(e) {
                console.log("close");
        }

        ws.onerror = function(e) {
                console.log("error");
        }
}

