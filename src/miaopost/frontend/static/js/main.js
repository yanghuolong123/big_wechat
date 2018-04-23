
// 建议删除
function suggestDel(infoId) {
	$.post("/info/suggestDel",{infoId:infoId}, function(e){
		if(e.code<0) {
			prompt(e.msg);
		}

		//prompt("提交成功，我们会在审核后删除影响平台秩序的发布。谢谢你的贡献。");
		greeting({msg:"提交成功，我们会在审核后删除影响平台秩序的发布。谢谢你的贡献。"})
	})
}

// 删除信息
var delInfo = function(id, obj){
	$this = $(obj);
	actionConfirm({msg:"你确认要删除?",confirm:function(){
                        	$.post("/info/delete",{id:id}, function(e){
			if(e.code<0) {
	                                	prompt(e.msg);
	                                	return false;
	                        	}
	                        	$this.parents(".info").remove();
	                        	prompt({msg:"信息删除成功!",displayTime:3000});
		});
                }});	
}

// 删除信息
var  topInfo = function(id){
	$.post("/info/top",{id:id}, function(e){
		if(e.code<0) {
                                	prompt(e.msg);
                                	return false;
                        	}
                        	prompt({msg:"信息置顶成功!",displayTime:3000});
                        	setTimeout(function(){
                        		window.location = location.href;
                        	}, 2500);
	});
}

$(function(){

	// 搜索
	$(".search-btn").click(function(){
		var sval = $("#search").val();
		if(sval=="") {
			return false;
		}

		window.location = "/info/list?search="+sval;
	});

	// 发布信息
	$("#create_info_btn, #edit_info_btn").click(function(){
		$this = $(this);
		var id = $("#info_id").val();
		var cid = $("#cid").val();
		var info_content = $("#info_content").val();
		var valid_day = $("#valid_day").val();
		var email = $("#email").val();
		var photo="";
		for( var s=0; s<$('.img-li-new').length; s++) {
			photo += $('.img-li-new').eq(s).attr('data-url') + ',';
		}
		photo=photo.substring(0,photo.length-1);

		var flag = true;
		$(".error_tips").html("");
		if(cid == "") {
			$(".error_tips").append("<p>请先选择分类</p>");
                        		flag = false;
		}
		if(info_content == ""  && photo=="") {
			$(".error_tips").append("<p>请添加文字描述或图片</p>");
                        		flag = false;
		}
		if(valid_day!="") {
			if(!/^[0-9]+$/.test(valid_day)){
			        $(".error_tips").append("<p>自动删除发布请填写数字</p>");
                        		        flag = false;
			}
		}
		if(email!="") {
			if(!/^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/.test(email)){
			        $(".error_tips").append("<p>邮箱格式不正确</p>");
                        		        flag = false;
			}
		}

		 if (!flag) {
	                        return false;
	                }

	                $this.attr("disabled","disabled");
	                var url = "/info/create";
	                if (id>0) {
	                	url = "/info/edit";
	                }

	                $.post(url, {id:id,cid:cid, content:info_content, valid_day:valid_day, email:email,photo:photo}, function(e){
	                	$this.removeAttr("disabled");
	                        	if(e.code<0) {
	                                	$(".error_tips").append(e.msg);
	                                	return false;
	                        	}

	                        	window.location = "/info/view?id="+e.data.Id;
	                });

	});

	// 删除信息
	$("#del_info_btn").click(function(){
		actionConfirm({msg:"您确定要删除您发布的此信息吗？",confirm:function(){
			var id = $("#info_id").val();
			$.post("/info/delete",{id:id}, function(e){
				if(e.code<0) {
		                                	prompt(e.msg);
		                                	return false;
		                        	}

		                        	prompt({msg:"信息删除成功!",displayTime:3000});
		                        	setTimeout(function(){
		                        		window.location = "/";
		                        	}, 2500);
			});
		}});		
	});


	// 图片上传
	$(".img-up-list").on("click", ".img-li i", function(){
		$(this).parent('.img-li').remove();
		$('.user-img').show();
		return false;
	});

	$('#imgs').on('change', function() {
		var formData = new FormData();
		formData.append('file', $('#imgs')[0].files[0]);				
		$.ajax({
			url: '/uploadfile',
			type: 'post',
			cache: false,
			data: formData,
			processData: false,
			contentType: false,
			success:function(rs,textStatus,jqXHR){
				if( rs.code <0) {
					prompt(rs.msg);
					return false;
				}

				var upImg = rs.data;
				if( rs.code == 0) {
					$('.img-up-list').append('<div class="img-li img-li-new" data-url="' + upImg+ '"  data-big="' + upImg + '" style="background-image:url(' + upImg+ '!200!200)"><i></i></div>');
				}	
				
				// $('.img-li i').on('click', function() {
				// 	$(this).parent('.img-li').remove();
				// 	$('.user-img').show();
				// 	return false;					
				// });
							
			}
		});
	});


	// 查看更多
	var hasMore = $("#hasMore").val();
	if(hasMore == 0 || hasMore=="false") {
		$(".load-more").hide();
	}

	$(".load-more").click(function(){
		$(".loading").append("<img src=\"/static/img/loading.gif\"/>");
		var page = parseInt($("#page").val())+1;
		var cid = $("#cid").val();
		var uid = $("#uid").val();
		$.post("/info/listPage", {cid:cid, page:page, uid:uid}, function(e){
			isloading = false;
			if( e.code<0) {
				prompt(e.msg);
			}
			$(".loading").empty();

			//$(".info-list").append(e.data.listData);
			$(".info-list .info:last").after(e.data.listData);
			$("#page").val(e.data.page);
						
			$("#hasMore").val(e.data.hasMore);
			hasMore = e.data.hasMore;
			if(hasMore == 0 || hasMore==false) {
				$(".load-more").hide();
			}
		});
		$(this).blur();
	});


	

});

// 分页列表
// var isloading = false;
// var hasMore  = 0;
// $(window).on('scroll', function() {	
// 	var windowHeight = $(window).height();
// 	var bodyHeight = $(document).height();
// 	var scrollHeight = $(document).scrollTop();
// 	hasMore = $("#hasMore").val();
// 	if (!isloading && hasMore == 1 && (scrollHeight >= ( bodyHeight - windowHeight )/2 ) && ( bodyHeight >= windowHeight ) ) {
// 		isloading = true;
// 		var page = parseInt($("#page").val())+1;
// 		var cid = $("#cid").val();
// 		$.post("/info/listPage", {cid:cid, page:page}, function(e){
// 			isloading = false;
// 			if( e.code<0) {
// 				prompt(e.msg);
// 			}

// 			$(".info-list").append(e.data.listData);
// 			$("#page").val(e.data.page);
						
// 			$("#hasMore").val(e.data.hasMore);
// 			hasMore = e.data.hasMore;
// 		});
// 	}		
// });