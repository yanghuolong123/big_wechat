
var ws = {};
var $ul;

$(function() {	
	$ul = $('#msg-list');
	if(sessionId>0) {
		listen()
	} else {
		$.get("/login",function(e){
			$("#modalPage").html(e.data).find('#loginModal').modal({backdrop: 'static', keyboard: false});
		});
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

	$('#modalPage').on("click","#loginLink",function(){
		$.get("/login", function(e){
			$("#modalPage").html(e.data).find('#loginModal').modal({backdrop: 'static', keyboard: false});
		});
	});

	$('#modalPage').on("click","#loginBtn",function(){
		var username = $("#username").val();
		var password = $("#password").val();
		$.post("/login", {username:username, password:password}, function(e){
			if(e.code<0) {
				alert(e.msg);
				return;
			}
			$('#loginModal').modal('hide');
			$(".modal-backdrop").remove();

			sessionId = e.data.user.Id;
			gid = e.data.user.Gid;
			nickname = e.data.user.Nickname;
			follow = e.data.follow;
			groupname = e.data.group.Name;
			listen();
		});
	});

	$('#modalPage').on("click","#registerLink",function(){
		$.get("/register", function(e){
			$("#modalPage").html(e.data).find('#registerModal').modal({backdrop: 'static', keyboard: false});
		});
	});

	$("#modalPage").on("click","#registerBtn",function(){
		var group = $("#group").val();
		if(group == "0") {
			alert("请选择学校");
			return false;
		}
		var username = $("#username").val();
		if(username == "") {
			alert("请填写账号");
			return false;
		}
		var nickname = $("#nickname").val();
		if(nickname == "") {
			alert("请填写昵称");
			return false;
		}
		var password = $("#password").val();
		if(password == "") {
			alert("请填写密码");
			return false;
		}
		var repassword = $("#repassword").val();
		if(repassword == "") {
			alert("请填写重复密码");
			return false;
		}

		if(password != repassword) {
			alert("密码输入不一致");
			return false;
		}

		$.post("/register",{group:group,username:username,nickname:nickname,password:password,repassword:repassword},function(e){
			if(e.code<0) {
				alert(e.msg);	
				return false;
			}
			$('#registerModal').modal('hide');
			$(".modal-backdrop").remove();

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

