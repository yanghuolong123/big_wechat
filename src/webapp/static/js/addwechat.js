// addwechat.js 

function uploadFile(obj) {
	var $this = $(obj);
	  //取得前面的DIv的第一个元素，也就是file框。
	  var $pre = $this.prev();	  
	  $pre.click();

	  $pre.on('change',function(){
	  	 var $upload = $(this);
	  	 if ($upload.val() == null || $upload.val() =="") {
	  	 	alert("请选择文件！");
			 return false;
		 } 

		 var formData = new FormData();
		formData.append('file', $upload[0].files[0]);				
		$.ajax({
			url: '/uploadfile',
			type: 'post',
			cache: false,
			data: formData,
			dataType: "json",
			processData: false,
			contentType: false,
			success:function(e,textStatus,jqXHR){
				$this.parent().css( 'background-image', 'url(/' +  e.data + ')' );
				$this.parent().addClass('active');
				$upload .prev().val(e.data);
			}
		});		
	  });
}

$(function(){

	$("#loginBtn").click(function(){
		$('#loginModal').modal({backdrop: 'static', keyboard: false});
		$.get("/login", function(e){
			$("#login_qrcode_img").html(e.data);

			var timer = setInterval(function(){
			    $.post('/login', {"sceneId":$("#sceneId").val()}, function(e){
			            if(e.code < 0) {
			                return false;
			            }
			            
			            clearInterval(timer);
			            if(e.data>0) {
			            	window.location = "/user";
			            	return;
			            }

			            window.location = "/";
			            
			        });
			}, 1000);
		});
	});


	$("#create_pg_btn").click(function(){
		$this = $(this);
		var gid = $("#search_group").val();
		var name = $.trim($("#name").val());
		var introduction = $("#introduction").val();
		var qrcode = $("#qrcode").val();
		var ower_qrcode = $("#ower_qrcode").val();
		var wechat_id = $.trim($("#wechat_id").val());

		var flag = true;
		$(".error_tips").html("");
		if (gid=="") {
			$(".error_tips").append("<p>应用范围为必填项。</p>");
			flag = false;
		}
		if (name=="") {
			$(".error_tips").append("<p>群名称为必填项。</p>");
			flag = false;
		}
		if(qrcode == "" && ower_qrcode=="" && wechat_id == "") {
			$(".error_tips").append("<p>上传联系信息至少要填一项。</p>");
			flag = false;
		}

		if (!flag) {
			return false;
		}

		$this.attr("disabled","disabled");
		$.post("/pg/create", {gid:gid, name:name, introduction:introduction,qrcode:qrcode, ower_qrcode:ower_qrcode, wechat_id:wechat_id}, function(e){
			$this.removeAttr("disabled");
			if(e.code<0) {
				$(".error_tips").append(e.msg);
				return false;
			}

			window.location = "/pg/view?id="+e.data.Id;
		});
	});

	$("#edit_pg_btn").click(function(){
		$this = $(this);
		var id = $("#pg_id").val();
		var gid = $("#search_group").val();
		var name = $.trim($("#name").val());
		var introduction = $("#introduction").val();
		var qrcode = $("#qrcode").val();
		var ower_qrcode = $("#ower_qrcode").val();
		var wechat_id = $.trim($("#wechat_id").val());

		var flag = true;
		$(".error_tips").html("");
		if (gid=="") {
			$(".error_tips").append("<p>应用范围为必填项。</p>");
			flag = false;
		}
		if (name=="") {
			$(".error_tips").append("<p>群名称为必填项。</p>");
			flag = false;
		}
		if(qrcode == "" && ower_qrcode=="" && wechat_id == "") {
			$(".error_tips").append("<p>上传联系信息至少要填一项。</p>");
			flag = false;
		}

		if (!flag) {
			return false;
		}

		$this.attr("disabled","disabled");
		$.post("/pg/edit", {id:id, gid:gid, name:name, introduction:introduction,qrcode:qrcode, ower_qrcode:ower_qrcode, wechat_id:wechat_id}, function(e){
			$this.removeAttr("disabled");
			if(e.code<0) {
				$(".error_tips").append(e.msg);
				return false;
			}

			window.location = "/pg/view?id="+id;
		});
	});


	$(".pgmsg-btn").click(function(){
		var pg_id = $("#pg_id").val();
		var content = $.trim($("#pg_msg").val());
		if ( content=="") {
			return false;
		}

		$.post("/pg/createPgMsg", {pg_id:pg_id,content:content}, function(e){
			if(e.code<0) {
				return false;
			}

			$("#pg_msg").val("");

			var comments = '';
			comments += "<li>";
			comments += " 	<h5>"+e.data.User.Nickname+"</h5>";
			comments += "		<p>"+e.data.Pgm.Content+"</p>";
			comments += "		<p>1秒前</p>";
			comments += "</li>";

			$("#commentlist").prepend(comments);
		});
	});

	$(".report_pg").click(function(){
		$("#pgReportModal").modal({backdrop: 'static', keyboard: false});
	});

	$("#report_pg_btn").click(function(){
		var pg_id = $("#pg_id").val();
		var content = $.trim($("#pg_report_content").val());
		if ( content=="") {
			return false;
		}

		$.post("/pg/createReport", {pg_id:pg_id,content:content}, function(e){
			if(e.code<0) {
				return false;
			}

			$("#pg_report_content").val("");
			$("#pgReportModal").modal('hide');
		});
	});

	$("#unlock_group").click(function(){
		$this = $(this);
		var gid = $("#gid").val();

		$.post("/pg/unlock",{gid:gid}, function(e){
			if(e.code<0) {
				prompt(e.msg);
				return false;
			}

			location.reload() ;

			$this.text("已解锁");
			$this.attr("disabled","disabled");
		});	
	});

	$('[data-toggle="popover"]').popover();

	var objects = {};
	$("#search").typeahead({
		source: function(query, process) {
			$.post("/search/group",{name: query}, function(e){
				if(e.code<0) {
					return false;
				}

				var results = [];
				var data = e.data;
				for (var i = 0; i < data.length; i++) {
					// if(data[i].Name!="") {
					// 	objects[data[i].Name] = data[i].Id;
					// 	results.push(data[i].Name);
					// } else if(data[i].Short_name!="") {
					// 	objects[data[i].Short_name] = data[i].Id;
			  //               		results.push(data[i].Short_name);
					// } else {
						objects[data[i].En_name] = data[i].Id;
						results.push(JSON.stringify(data[i]));
					// }			                 	
			                }
				process(results);
			});
		},
	                highlighter: function (jsonStr) {
	                	var item = JSON.parse(jsonStr);
	                	var  str = "<div class='media search_group'>";
	                	str += "	<div class='media-left'>";
	                	str += "		<img class='media-object img-rounded search_logo' src='/static/images/default_group_logo.png' >";
	                	str += "	</div>";
	                	str += "	<div class='media-body'>";
	                	str += "		<p>"+item.En_name+"</p>";
	                	str += "		<p class='group_region'>"+item.Region+"</p>";
	                	str += "	</div>";
	                	str += "</div>";
	                	return str;
		    },
		updater: function (jsonStr) {
		        var item = JSON.parse(jsonStr);
		        return item.En_name;
		},
		afterSelect: function (item) { 
			$("#search_group").val(objects[item]);
			
			 if ( !$("#search").hasClass("school") ) {
			 	window.location.href = "/pg/list?gid="+objects[item];
			 }			
		},
	});

	$("#publish_pg, .pg_list .list a").click(function(){
		var uid = $("#uid").val();
		if(uid<=0) {
			$("#loginBtn").trigger("click");
			return false;
		} 
	});

});