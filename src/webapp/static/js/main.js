var ws = {};
$(function() {
	var $ul = $('#msg-list');
	ws = new WebSocket("ws://"+ window.location.host +"/chat");

	ws.onopen = function(e) {
	  var msg = {}
	  msg.uid= 1;
	  msg.gid = [1,2];
	  msg.type = "login";
	  ws.send(JSON.stringify(msg));

	  console.log("open");
	}

	ws.onmessage = function(e) {
	  var msg = JSON.parse(e.data);
	  var content = "时间:"+ msg.createTime + " 内容:"+ msg.content;
	  $('<li>').text(content).appendTo($ul);
	};

	$('#sendBtn').click(function(){
	  var content = $('#name').val();
	  var msg = {};

	  msg.uid= 1;
	  msg.gid = [1,2];
	  msg.type = "message";
	  msg.content= content;
	  ws.send(JSON.stringify(msg));

	  $('#name').val("");
	 });

	ws.onclose = function(e) {
		console.log("close");
	}

	ws.onerror = function(e) {
		console.log("error");
	}
});

