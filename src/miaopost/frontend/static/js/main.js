
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

// 置顶信息
var  topInfo = function(id){
	$.post("/info/top",{id:id}, function(e){
		if(e.code<0) {
                                	prompt(e.msg);
                                	return false;
                        	}
                        	//prompt({msg:"置顶成功，您的发布已成为最新发布，您可以多次使用置顶，提升显示效果。",displayTime:3000});
                        	prompt({msg:e.msg,displayTime:3000});
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

		if(cid == "") {
			prompt("请先选择分类");
                        	return false;
		}
		var content="";
		cArr = info_content.split("\n");
		for(var i=0; i<cArr.length;i++) {
			line = cArr[i];
			c = line.split("：");		
			if(c[0]!="描述" && c[0]!="价格" && c[0]!="地址" && c[0]!="联系方式") {
				content += line;
				continue;
			}
			if($.trim(c[1])!=="") {
				content += line;
			}
		}
		info_content = $.trim(content);
		if(info_content == ""  && photo=="") {
			prompt("请添加文字描述或图片");
                        	return false;
		}
		if(valid_day!="") {
			if(!/^[0-9]+$/.test(valid_day)){
			        	prompt("自动删除发布请填写数字");
                        		return false;
			}
		}
		if(email!="") {
			if(!/^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/.test(email)){
			       	 prompt("邮箱格式不正确");
                        		return false;
			}
		}


		var reward_type,reward_amount,reward_num;		
		reward_type = $('input[name="reward_type"]:checked').val();
		reward_amount = $('input[name="reward_amount"]:checked').val();
		reward_num = $('input[name="reward_num"]:checked').val();
		if(reward_type>0 || reward_amount>0 || reward_num>0) {	
			if (typeof(reward_type) == "undefined") { 
			    prompt("请选择红包类型");
			    return false;
			} 

			if (typeof(reward_amount) == "undefined") { 
			    prompt("请选择红包平均金额");
			    return false;
			} 			
			
			if (typeof(reward_num) == "undefined") { 
			    prompt("请选择红包个数");
			    return false;
			} 
		}

	                $this.attr("disabled","disabled");
	                var url = "/info/create";
	                if (id>0) {
	                	url = "/info/edit";
	                }

	                $.post(url, {id:id,cid:cid, content:info_content, valid_day:valid_day, email:email,photo:photo,reward_type:reward_type,reward_amount:reward_amount,reward_num:reward_num}, function(e){	                	
	                        	if(e.code<0) {
	                                		prompt(e.msg);
	                                		$this.removeAttr("disabled");
	                                		return false;
	                        	}

	                        	if(e.data.Reward_type>0) {
	                        		amount = e.data.Reward_amount*e.data.Reward_num;

	                        		var balance=0;
					$.ajax({
				                url:"/pay/balance",
				                async:false,
				                type: "POST",
				                data: {amount:amount, type:2,product_id:e.data.Id},
				                success: function(e){
				                        if(e.code<0) {
							return false;		
						}

						if(e.code==0) {
							balance = -1;
							prompt("发布红包信息成功！");		
							return false;
						}

						balance = e.data.Amount;
				                }
				       	});

					if(balance<0) {
						window.location = "/info/view?id="+e.data.Id+"&chance=no";
						return false;
					}

					amount -= balance;
					amount = amount.toFixed(2);
	                        		prompt({msg:"发布成功！支付完成后即可成功添加红包!",displayTime:2500});	                        		
	                        		if(isWeiXin()){
						//window.location.href = "/pay/confirm?product_id="+e.data.Id+"&amount="+amount+"&info_id="+e.data.Id+"&type=2&msg=亲, 信息发布成功，红包需要支付";						
		                        		setTimeout(function(){
			                        		window.location.href = "/pay/confirm?product_id="+e.data.Id+"&amount="+amount+"&info_id="+e.data.Id+"&type=2&msg=亲, 信息发布成功，红包需要支付";
			                        	}, 2500);
					} else {
						//prompt("亲, 信息发布成功，红包需要支付");
						setTimeout(function(){
							$("#pay_qr_img").removeClass("qrimg").attr("src", "/static/img/loading.gif");
							$(".pay_amount").html("￥"+amount+"元");
							$('#qrPayModal').modal({backdrop: 'static', keyboard: false});

							$.post("/pay/wxscan", {product_id:e.data.Id, amount:amount, type:2}, function(e){
								if(e.code<0) {
									prompt(e.msg);
									return false;
								}

								$("#pay_qr_img").attr("src", e.data.qrurl).addClass("qrimg");
								orderNo = e.data.order_no;

								var timer = setInterval(function(){
								    $.post('/pay/check', {order_no:orderNo}, function(e){
								            if(e.code < 0) {
								                return false;
								            }
								            
								            clearInterval(timer);
								            $('#qrPayModal').modal("hide");
								            prompt({msg:"红包支付成功！感谢您的支持！",displayTime:3000});
								            setTimeout(function(){
						                        		window.location = "/info/view?id="+e.data.Product_id+"&chance=no";
						                        	}, 2500);
								            
								        });
								}, 1000);

							});
						}, 2500);

						
					}

	                        	} else {
	                        		window.location = "/info/view?id="+e.data.Id;
	                        	}

	                        	// window.location = "/info/view?id="+e.data.Id;
	                });

	});

	//  分类刷新
	$(".create #cid").change(function(){
		location.href = "/info/create?cid="+$(this).val();
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

	// $('#imgs').on('change', function() {
	// 	$('.img-up-list').after('<div class="progress">'+
	// 	  '<div class="progress-bar progress-bar-striped active" role="progressbar" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100" style="width: 1%">'+
	// 	    '<span class="sr-only">100% Complete</span>'+
	// 	  '</div>'+
	// 	'</div>');
	// 	var n = 10;
	// 	var t = setInterval(function(){
	// 		$(".progress-bar").css("width", n+"%");
	// 		if (n<95) {
	// 			n += 5;
	// 		}			
	// 	}, 50);
	// 	var formData = new FormData();
	// 	formData.append('file', $('#imgs')[0].files[0]);				
	// 	$.ajax({
	// 		url: '/uploadfile',
	// 		type: 'post',
	// 		cache: false,
	// 		data: formData,
	// 		processData: false,
	// 		contentType: false,
	// 		success:function(rs,textStatus,jqXHR){
	// 			clearInterval(t);
	// 			$(".progress-bar").css("width", "100%");				
	// 			setTimeout(function(){
	// 				$(".progress").remove();
	// 			},1000);
	// 			if( rs.code <0) {
	// 				prompt(rs.msg);
	// 				return false;
	// 			}

	// 			var upImg = rs.data;
	// 			if( rs.code == 0) {
	// 				$('.img-up-list').append('<div class="img-li img-li-new" data-url="' + upImg+ '"  data-big="' + upImg + '" style="background-image:url(' + upImg+ '!200!200)"><i></i></div>');
	// 			}	
				
	// 			// $('.img-li i').on('click', function() {
	// 			// 	$(this).parent('.img-li').remove();
	// 			// 	$('.user-img').show();
	// 			// 	return false;					
	// 			// });
							
	// 		}
	// 	});
	// });


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


	$("#loginBtn").click(function(){
		$("#login_qrcode_img").html('<h4 style="color: red;">专属通道（二维码）生成中…<img src="/static/img/loading.gif" /></h4>');
		$('#loginModal').modal({backdrop: 'static', keyboard: false});
		$.get("/login", function(e){
			$("#login_qrcode_img").html(e.data);

			var timer = setInterval(function(){
			    $.post('/login', {"sceneId":$("#sceneId").val()}, function(e){			    	
			            if(e.code < 0) {
			                	return false;
			            }
			            
			            clearInterval(timer);
			            if(e.code==0) {
			            	$('#loginModal').modal("hide");
			            	prompt({msg:"已登录！您可以通过用户中心进行编辑、删除和免费置顶操作",displayTime:3000});
		                        	setTimeout(function(){
		                        		window.location = location.href;
		                        	}, 2500);

			            	return;
			            }

			            window.location = "/";
			            
			        });
			}, 1000);
		});
	});


	// 发布页扫码登陆
	var clq = $(".create_login_qrcode");
	if(clq.length>0) {
		$.get("/login",function(e){
			clq.html(e.data);
			var timer = setInterval(function(){
			    $.post('/login', {"sceneId":$("#sceneId").val()}, function(e){
			            if(e.code < 0) {
			               	return false;
			            }
			            
			            clearInterval(timer);
			            if(e.code==0) {
			            	prompt({msg:"已登录！此时通过电脑或公众号发布后都可以进行编辑、删除和免费置顶操作",displayTime:3000});
		                        	setTimeout(function(){
		                        		//window.location = "/user";
		                        		window.location = location.href;
		                        	}, 2500);
			            	
			            	return;
			            }

			            window.location = "/";
			            
			        });
			}, 1000);
		});
	}

	// 弹出留言框
	$(".msg-btn").click(function(){
		var uid = $("#login_uid").val();
		if(uid==0) {
			//prompt("请先登陆！")
			$("#loginBtn").trigger("click");
			return false;
		}

		$('#msgModal').modal({backdrop: 'static', keyboard: false});
	});

	// 留言/回复
	$(".info-msg-btn").click(function(){
		var info_id = $("#info_id").val();
		var pid = $("#msg_pid").val();
		var content = $.trim($("#info_msg").val());
		if ( content=="") {
			//prompt("内容不能为空");
			return false;
		}

		$.post("/msg/create", {info_id:info_id,content:content,pid:pid}, function(e){
			$("#msg_pid").val(0);
			$('#msgModal').modal('hide');
			if(e.code<0) {
				prompt({msg:e.msg});
				return false;
			}

			//$("#info_msg").val("");
			editor.html("");
			var p = "";
			if (e.data.Im.Pid>0) {
				p= "@"+e.data.Parent.User.Nickname;
			}

			var comments = '<div class="msg">';
			comments += '<div class="row">';
			comments += ' <div class="col-md-8 col-xs-3">';
			comments += ' <span><a href="javascript:;">'+e.data.User.Nickname+'</a> '+p+'</span>';
			comments += '</div> ';
			comments += '<div class="col-md-4 col-xs- 9 text-right"> ';
			comments += ' <span><a href="#" onclick="replyMsg('+e.data.Im.Id+');return false;">回复</a></span>';
			comments += ' <span><a href="#" onclick="msgDelSuggest('+e.data.Im.Id+', this);return false;">建删</a></span>';
			if (e.data.Cat.Type==1) {
			comments += ' <span><a href="#" onclick="admire('+e.data.Im.Id+','+e.data.Im.Uid+'); return false;">赞赏</a></span>';
			comments += ' <span><a href="#" onclick="supportInfoMsg('+e.data.Im.Id+', this);return false;"><span class="glyphicon glyphicon-heart" aria-hidden="true"></span><span class="support_num">'+e.data.Im.Support+'</span></a></span>';			
			}
			comments += ' </div>';
			comments += ' </div>';
			comments += '<div class="row"> ';
			comments += '<div class="col-md-12">'+e.data.Im.Content+'</div> ';
			comments += ' </div>';
			comments += '</div>';

			$("#commentlist").prepend(comments);

			if(e.data.Ireward != null) {
				greeting({title:"提示",msg:"恭喜您! 您获得了 "+e.data.Ireward.Amount+"元 留言红包。"})	
			}
		});
	});


});

// 回复留言
var replyMsg = function(pid) {
	var uid = $("#login_uid").val();
	if(uid==0) {
		//prompt("请先登陆！")
		$("#loginBtn").trigger("click");
		return false;
	}
	
	$("#msg_pid").val(pid);
	$('#msgModal').modal({backdrop: 'static', keyboard: false});

}

// 点赞支持 留言
var supportInfoMsg= function(id, obj) {
	$this = $(obj);
	$.post("/msg/support",{id:id},function(e){
		if(e.code<0) {
			prompt({msg:e.msg});
			return false;
		}

		var support = $this.find(".support_num");
		support.text(parseInt(support.text())+1);

		//greeting({msg:"点赞成功，谢谢你的贡献。"})
	});
}

// 建议删除 留言
var msgDelSuggest = function(id) {
	$.post("/msg/suggestDel",{id:id},function(e){
		if(e.code<0) {
			prompt({msg:e.msg});
			return false;
		}

		greeting({msg:"提交成功，我们会在审核后删除影响平台秩序的留言。谢谢你的贡献。"})
	});
}

// 支付赞赏
var admire = function(id,toUid) {
	var uid = $("#login_uid").val();
	if(uid==0) {
		prompt("请先登陆！")
		return false;
	}

	$("#admire_msg_id").val(id);
	$("#toUid").val(toUid);
	$('#admireModal').modal({backdrop: 'static', keyboard: false});
}

var admirePay = function(amount) {
	$('#admireModal').modal('hide');	

	var mid = $("#admire_msg_id").val();
	var toUid = $("#toUid").val();
	var orderNo;

	var balance=0;
	$.ajax({
                url:"/pay/balance",
                async:false,
                type: "POST",
                data: {amount:amount, type:1, toUid:toUid,product_id:mid},
                success: function(e){
                        if(e.code==-1) {
			return false;		
		}

		if(e.code==0) {
			balance = -1;
			prompt("赞赏支付成功！感谢您的支持！");		
			return false;
		}

		balance = e.data.Amount;
                }
       	});

	if (balance<0) {
		return false;
	}

	if (amount>balance) {
		amount -=  balance ;
		amount = amount.toFixed(2);
	}

	if(isWeiXin()){
		window.location.href = "/pay/confirm?product_id="+mid+"&amount="+amount+"&info_id="+$("#info_id").val()+"&type=1&msg=亲, 感谢您对此留言信息赞赏，需要支付";
	} else {
		$("#pay_qr_img").removeClass("qrimg").attr("src", "/static/img/loading.gif");
		$(".pay_amount").html("￥"+amount+"元");
		$('#qrPayModal').modal({backdrop: 'static', keyboard: false});

		$.post("/pay/wxscan", {product_id:mid, amount:amount, type:1}, function(e){
			if(e.code<0) {
				prompt(e.msg);
				return false;
			}

			$("#pay_qr_img").attr("src", e.data.qrurl).addClass("qrimg");
			orderNo = e.data.order_no;

			var timer = setInterval(function(){
			    $.post('/pay/check', {order_no:orderNo}, function(e){
			            if(e.code < 0) {
			                return false;
			            }
			            
			            clearInterval(timer);
			            $('#qrPayModal').modal("hide");
			            prompt("赞赏支付成功！感谢您的支持！");
			        });
			}, 1000);

		});
	}
	
}

// 用户提现
var withDraw=function() {
	$('#withDrawModal').modal({backdrop: 'static', keyboard: false});
}

$(function(){
	$(".with-draw-btn").click(function(e){
		$('#withDrawModal').modal('hide');

		var amount = $("#withdraw_amount").val();

		$.post("/pay/withdraw",{amount:amount},function(e){
			if(e.code<0) {
				prompt(e.msg);
				return false;
			}
			
		});
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
