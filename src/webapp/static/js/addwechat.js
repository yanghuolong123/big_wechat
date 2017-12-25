// addwechat.js 

$(function(){
	$("#loginBtn").click(function(){
		$.get("/login", function(e){
			$("#modalPage").html(e.data).find('#loginModal').modal({backdrop: 'static', keyboard: false});

			var timer = setInterval(function(){
			    $.post('/login', {"sceneId":$("#sceneId").val()}, function(e){
			            if(e.code < 0) {
			                return false;
			            }
			            
			            clearInterval(timer);
			            window.location = "/";
			            
			        });
			}, 1000);
		});
	});
});