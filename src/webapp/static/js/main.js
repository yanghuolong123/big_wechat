
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
	  var content = $('#msgContent').val();
	  var msg = {};

	  msg.uid = sessionId;
	  msg.gid = gid;
	  msg.type = "message";
	  msg.nickname = nickname;
	  msg.groupname = groupname;
	  msg.content= content;
	  ws.send(JSON.stringify(msg));

	  $('#msgContent').val("");
        });

	$('#login').click(function(){
		$.post("/login", {email:"yhl27ml@163.com", password:"123456"}, function(e){
			if(e.code<0) {
				alert(e.msg);
				return;
			}

			sessionId = e.data.user.Id;
			gid = e.data.user.Gid;
			nickname = e.data.user.Nickname;
			follow = e.data.follow;
			groupname = e.data.group.Name;
			listen();
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
          msg.uid = sessionId;
	  msg.gid = gid;
          msg.follow = follow;
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

