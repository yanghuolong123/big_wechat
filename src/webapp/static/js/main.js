
var ws = {};
var $ul;

$(function() {	
	$ul = $('#msg-list');
	if(sessionId>0) {
		listen()
	} else {
		$('#loginModal').modal({backdrop: 'static', keyboard: false});
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
		var username = $("#username").val();
		var password = $("#password").val();
		$.post("/login", {email:username, password:password}, function(e){
			if(e.code<0) {
				alert(e.msg);
				return;
			}
			$('#loginModal').modal('hide');

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
          var content = '<a class="user" href="#"><img class="img-responsive avatar_" src="/static/images/avatar-1.png" alt=""><span class="user-name">'+msg.nickName+'</span></a>'+
	'<div class="reply-content-box">'+
	'<span class="reply-time">'+msg.createTime+'</span>'+
	'<div class="reply-content pr">'+
	'<span class="arrow">&nbsp;</span>' +
	msg.content +
	'</div>'+
	'</div>';
	  
	  var show = "odd";
	  if(sessionId==msg.uid) {
		show = "even";
	  }	
          $('<li class="'+show+'">').html(content).appendTo($ul);
        };

        ws.onclose = function(e) {
                console.log("close");
        }

        ws.onerror = function(e) {
                console.log("error");
        }
}

