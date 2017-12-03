var ws = {};

$(function() {
	var $ul = $('#msg-list');
	ws = new WebSocket("ws://"+ window.location.host +"/chat");

	ws.onopen = function(e) {
		console.log("open");
	}

	ws.onmessage = function(e) {
	  var msg = JSON.parse(e.data);
	  var content = "时间:"+ msg.createtime + " 内容:"+ msg.body;
	  $('<li>').text(content).appendTo($ul);
	};

	$('#sendBtn').click(function(){
	  var content = $('#name').val();
	  var msg = {};

	  msg.author = 'yhl';
	  msg.body = content;
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

