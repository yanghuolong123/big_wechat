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
                <p>
                <a class="btn btn-success" role="button" data-toggle="collapse" href="#collapseReward" aria-expanded="false" aria-controls="collapseExample">
                  添加红包
                </a>
                </p>
                <div class="collapse" id="collapseReward">
                  <div class="well">

                    <div class="row">
                      <div class="col-sm-4"><label class="radio-inline reward_title">红包类型：</label></div>
                      <div class="col-sm-8">
                        <label class="radio-inline">
                          <input type="radio" name="reward_type" value="1"> 阅读红包
                        </label>
                        <label class="radio-inline">
                          <input type="radio" name="reward_type"  value="2"> 留言红包
                        </label>
                      </div>
                    </div>

                    <div class="row">
                      <div class="col-sm-4"><label class="radio-inline reward_title">平均金额：</label></div>
                      <div class="col-sm-8">
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

                    <div class="row">
                      <div class="col-sm-4"><label class="radio-inline reward_title">红包个数：</label></div>
                      <div class="col-sm-8">
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

                    <div class="row">
                      <div class="col-sm-12"><label class="radio-inline reward_title">总金额：<span class="total_reward_amount"></span></label></div>
                    </div>

                    <div class="row">
                      <div class="col-sm-offset-1 col-sm-11">
                        <label class="checkbox-inline reward_title">
                          <input type="checkbox" name="reward_confirm" id="reward_confirm" value="1"> 确认/取消
                        </label>
                      </div>
                    </div>

                  </div>
                </div>
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
{{end}}
