<style>
  #picker div:nth-child(2) input{width:100%!important;height:100%!important;}
</style>
<div class="create">
       {{if and (not .isWeixin) (not .user)}}
       <h4>微信登录发布</h4>
       <div class="row">
         <div class="col-sm-offset-1 wx">
          <p class="text-danger">(推荐，发布后可以非常方便地通过服务号或电脑对发布进行编辑、删除和置顶操作)</p>        
          
          <div class="create_login_qrcode">
            <img src="/static/img/loading.gif">
          </div>
           
         </div>
       </div>
       <hr>
       <h4>免注册发布</h4>
       {{end}}
        <form class="form-horizontal">
            <div class="form-group">            
              <div class="col-sm-offset-1 col-sm-4">
                <select id="cid" class="form-control">
                  <option value="">请选择分类</option>
                  {{range .cats}}
                    <option value="{{.Id}}" {{if  eq .Id $.cat.Id}}selected="selected"{{end}}>{{.Name}}</option>
                  {{end}}
                </select>
              </div>
            </div>          
            <div class="form-group">             
              <div class="col-sm-offset-1 col-sm-6">
                <label>图片和文字 <span class="text-danger label-tips">(至少一项)</span></label>
                <div class="img-up">
                    <div class="img-up-list clearfix">                        
                      
                    </div>
                    <div id="thelist" class="uploader-list"></div>
                    <div id="picker"><label class="user-img" for="imgs"></label></div>
                </div>
                <textarea class="form-control"  id="info_content" rows="5" maxlength="2500" placeholder="添加描述">
描述：
价格：
地址：
联系方式：
                </textarea>                
              </div>
            </div>

            {{if  eq .cat.Type 1}}
            <div class="form-group" >            
              <div class=" col-sm-offset-1 col-sm-6">
                <button type="button" class="btn btn-success" id="withdraw"> 添加红包 </button>
                <span class="withdraw_icon"></span>
              </div>
             </div>
            {{end}}

            <div class="form-group" style="{{if .user}}display: none;{{end}}">            
              <div class=" col-sm-offset-1 col-sm-3">
                <label>自动删除发布 <span class="text-danger label-tips">(建议填写)</span></label>
                <div class="row">
                  <div  class="col-md-8 col-xs-8">
                  <input type="text" class="form-control"  id="valid_day" maxlength="10" placeholder=""> 
                  </div>
                  <div class="col-md-4 col-xs-3 valid_date">
                    <span class="">天后</span>
                  </div>
                </div>
              </div>
            </div>
            
             <div class="form-group" style="{{if .user}}display: none;{{end}}">            
              <div class=" col-sm-offset-1 col-sm-4">
                <label>我们会发送编辑链接至邮箱 <span class="text-danger label-tips">(建议填写)</span></label>
                <input type="text" class="form-control" id="email" maxlength="55" placeholder="邮箱地址">
              </div>
            </div> 

            <div class="form-group">
              <div class="col-sm-offset-1 col-sm-10">
                <button type="button" id="create_info_btn" class="btn btn-primary btn-lg">发布</button>
              </div>
            </div>
          </form>
</div>

{{if  eq .cat.Type 1}}

{{template "pay/qrcode.tpl" .}} 

<!-- Modal -->
<div class="modal fade" id="withdrawModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" id="withdrawModalClose" class="close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="myModalLabel">添加红包</h4>
      </div>
      <div class="modal-body">
        
        <form class="form-horizontal">
          <div class="form-group">
            <label for="inputEmail3" class="col-sm-3 control-label">红包类型：</label>
            <div class="col-sm-9">
                        <label class="radio-inline">
                          <input type="radio" name="reward_type" value="1"> 阅读红包
                        </label>
                        <label class="radio-inline">
                          <input type="radio" name="reward_type"  value="2"> 留言红包
                        </label>
            </div>
          </div>
          <div class="form-group">
            <label for="inputPassword3" class="col-sm-3 control-label">平均金额：</label>
            <div class="col-sm-9">
                        <label class="radio-inline">
                          <input type="radio" name="reward_amount" value="0.1"> 0.1元
                        </label>
                        <label class="radio-inline">
                          <input type="radio" name="reward_amount" value="0.2"> 0.2元
                        </label>
                        <label class="radio-inline">
                         <input type="radio" name="reward_amount" value="0.5"> 0.5元
                        </label>
                        <label class="radio-inline">
                          <input type="radio" name="reward_amount" value="1"> 1元
                        </label>
            </div>
          </div>  
           <div class="form-group">
            <label for="inputPassword3" class="col-sm-3 control-label">红包个数：</label>
            <div class="col-sm-9">
                        <label class="radio-inline">
                          <input type="radio" name="reward_num"  value="10"> 10
                        </label>
                        <label class="radio-inline">
                          <input type="radio" name="reward_num"  value="20"> 20
                        </label>
                        <label class="radio-inline">
                         <input type="radio" name="reward_num"  value="30"> 30
                        </label>
                        <label class="radio-inline">
                          <input type="radio" name="reward_num"  value="50"> 50
                        </label>
            </div>
          </div>   
          <div class="form-group">
             <label for="inputPassword3" class="col-sm-3 control-label">总金额：</label> 
             <div class="col-sm-9">
               <span class="total_reward_amount"></span>
             </div>           
          </div>     
          
        </form>

      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-success width-draw-confirm">确定</button>
        <!--<button type="button" class="btn btn-primary width-draw-cancel">取消添加红包</button>-->
      </div>
    </div>
  </div>
</div>

<script type="text/javascript">

$(function(){
    // 添加红包
    $("#withdraw").click(function(){
          var reward_type;   
          reward_type = $('input[name="reward_type"]:checked').val();

          if (typeof(reward_type) == "undefined" ) {
              $("#withdrawModal").modal();              
           } else {
               $(this).text("添加红包");
               $('#withdrawModal input').removeAttr('checked'); 
                $(".total_reward_amount").html("");

                $(".withdraw_icon").html('');
                $("#withdrawModal").modal('hide');
           }
        
    });

    // 取消模态框
    $('#withdrawModalClose').click(function(){
                $('#withdrawModal input').removeAttr('checked'); 
                $(".total_reward_amount").html("");
                $("#withdrawModal").modal('hide');
    });

    // 红包处理
    $('input[name="reward_num"]').change(function(e){
         var reward_type,reward_amount,reward_num;   
          reward_type = $('input[name="reward_type"]:checked').val();
          reward_amount = $('input[name="reward_amount"]:checked').val();
          reward_num = $('input[name="reward_num"]:checked').val();

          if (typeof(reward_num) != "undefined" && typeof(reward_amount)!="undefined") {
              $(".total_reward_amount").html(reward_amount*reward_num+"元");
           }

          if (typeof(reward_type) == "undefined") { 
              prompt("请选择红包类型");
              return false;
          } 
         
          if (typeof(reward_amount) == "undefined") { 
              prompt("请选择红包平均金额");
              return false;
          }           
          
    });

    $('input[name="reward_amount"]').change(function(e){
        var reward_type,reward_amount,reward_num;   
          reward_type = $('input[name="reward_type"]:checked').val();
          reward_amount = $('input[name="reward_amount"]:checked').val();
          reward_num = $('input[name="reward_num"]:checked').val();

          if (typeof(reward_type) == "undefined") { 
              prompt("请选择红包类型");
          } 

           if (typeof(reward_num) != "undefined") { 
             $(".total_reward_amount").html(reward_amount*reward_num+"元");
          } 

    });

    $(".width-draw-confirm").click(function(){
        var reward_type,reward_amount,reward_num;   
        reward_type = $('input[name="reward_type"]:checked').val();
        reward_amount = $('input[name="reward_amount"]:checked').val();
        reward_num = $('input[name="reward_num"]:checked').val();

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

        $("#withdraw").text("取消红包");
        $(".withdraw_icon").html('<img class="img_tip1" src="/static/img/reward_type'+reward_type+'.png"/>');

        $("#withdrawModal").modal('hide');
    });

    // $(".width-draw-cancel").click(function(){
    //     $('#withdrawModal input').removeAttr('checked'); 
    //     $(".total_reward_amount").html("");

    //     $(".withdraw_icon").html('');
    //     $("#withdrawModal").modal('hide');
    // });
});

    

</script>

{{end}}
